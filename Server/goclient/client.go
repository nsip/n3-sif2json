package goclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/cdutwhu/n3-util/n3err"
	"github.com/opentracing/opentracing-go"
	tags "github.com/opentracing/opentracing-go/ext"
)

// DOwithTrace :
func DOwithTrace(ctx context.Context, config, fn string, args *Args) (string, error) {

	Cfg := n3cfg.ToEnvN3sif2jsonGoclient(nil, envKey, config)
	failOnErrWhen(Cfg == nil, "%v", n3err.CFG_INIT_ERR)

	service := Cfg.Service
	if span := opentracing.SpanFromContext(ctx); span != nil {
		tracer := initTracer(service)
		span := tracer.StartSpan(fn, opentracing.ChildOf(span.Context()))
		tags.SpanKindRPCClient.Set(span)
		tags.PeerService.Set(span, service)
		if args != nil {
			span.SetTag(fn, *args)
		}
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	return DO(config, fn, args)
}

// DO : fn ["HELP", "SIF2JSON", "JSON2SIF"]
func DO(config, fn string, args *Args) (string, error) {

	Cfg := n3cfg.ToEnvN3sif2jsonGoclient(nil, envKey, config)
	failOnErrWhen(Cfg == nil, "%v", n3err.CFG_INIT_ERR)

	server := Cfg.Server
	protocol, ip, port := server.Protocol, server.IP, server.Port
	timeout := Cfg.Access.Timeout

	mFnURL, fields := initMapFnURL(protocol, ip, port, &Cfg.Route)
	url, ok := mFnURL[fn]
	if err := warnOnErrWhen(!ok, "%v: Need %v", n3err.PARAM_NOT_SUPPORTED, fields); err != nil {
		return "", err
	}

	chStr, chErr := make(chan string), make(chan error)
	go func() {
		rest(fn, url, args, chStr, chErr)
	}()

	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		return "", warnOnErr("%v: Didn't get response in %d(s)", n3err.NET_TIMEOUT, timeout)
	case str := <-chStr:
		err := <-chErr
		if err == n3err.NO_ERROR {
			return str, nil
		}
		return str, err
	}
}

// rest :
func rest(fn, url string, args *Args, chStr chan string, chErr chan error) {
	paramV, paramN := "", ""
	if args != nil && args.Ver != "" {
		paramV = fSf("sv=%s", args.Ver)
	}
	if args != nil && args.ToNATS {
		paramN = fSf("nats=true")
	}
	url = fSf("%s?%s&%s", url, paramV, paramN)
	url = sReplace(url, "?&", "?", 1)
	url = sTrimRight(url, "?&")

	logWhen(true, "accessing ... %s", url)

	var (
		Resp *http.Response
		Err  error
		Ret  []byte
	)

	switch fn {
	case "HELP":
		if Resp, Err = http.Get(url); Err != nil {
			goto ERR_RET
		}

	case "SIF2JSON", "JSON2SIF":
		if args == nil {
			Err = n3err.PARAM_INVALID
			goto ERR_RET
		}

		str := string(args.Data)
		if fn == "SIF2JSON" && !isXML(str) {
			Err = n3err.PARAM_INVALID_XML
			goto ERR_RET

		} else if fn == "JSON2SIF" && !isJSON(str) {
			Err = n3err.PARAM_INVALID_JSON
			goto ERR_RET
		}
		if Resp, Err = http.Post(url, "application/json", bytes.NewBuffer(args.Data)); Err != nil {
			goto ERR_RET
		}
	}

	if Resp == nil {
		Err = n3err.NET_NO_RESPONSE
		goto ERR_RET
	}
	defer Resp.Body.Close()

	if Ret, Err = ioutil.ReadAll(Resp.Body); Err != nil {
		goto ERR_RET
	}

ERR_RET:
	if Err != nil {
		chStr <- ""
		chErr <- Err
		return
	}

	chStr <- string(Ret)
	chErr <- n3err.NO_ERROR
	return
}
