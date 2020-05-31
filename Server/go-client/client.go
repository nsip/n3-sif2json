package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// DO : fn ["HELP", "SIF2JSON", "JSON2SIF"]
func DO(configfile, fn string, args Args) (string, error) {
	failOnErrWhen(!initEnvVarFromTOML("Cfg-Clt-S2J", configfile), "%v", eg.CFG_INIT_ERR)

	Cfg := env2Struct("Cfg-Clt-S2J", &Config{}).(*Config)
	server := Cfg.Server
	protocol, ip, port := server.Protocol, server.IP, server.Port
	timeout := Cfg.Access.Timeout
	setLog(Cfg.LogFile)

	mFnURL, fields := initMapFnURL(protocol, ip, port, Cfg.Route)
	url, ok := mFnURL[fn]
	if err := warnOnErrWhen(!ok, "%v: Need %v", eg.PARAM_NOT_SUPPORTED, fields); err != nil {
		return "", err
	}

	chStr, chErr := make(chan string), make(chan error)
	go func() {
		rest(fn, url, args, chStr, chErr)
	}()

	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		return "", warnOnErr("%v: Didn't get response in %d(s)", eg.NET_TIMEOUT, timeout)
	case str := <-chStr:
		err := <-chErr
		if err == eg.NO_ERROR {
			return str, nil
		}
		return str, err
	}
}

// rest :
func rest(fn, url string, args Args, chStr chan string, chErr chan error) {

	paramV, paramN := "", ""
	if args.Ver != "" {
		paramV = fSf("sv=%s", args.Ver)
	}
	if args.ToNATS {
		paramN = fSf("nats=true")
	}
	url = fSf("%s?%s&%s", url, paramV, paramN)
	url = sReplace(url, "?&", "?", 1)
	url = sTrimRight(url, "?&")

	logWhen(true, "accessing ... %s", url)

	var (
		Resp    *http.Response
		Err     error
		RetData []byte
	)

	switch fn {
	case "HELP":
		if Resp, Err = http.Get(url); Err != nil {
			goto ERR_RET
		}

	case "SIF2JSON", "JSON2SIF":
		str := string(args.Data)
		if fn == "SIF2JSON" && !isXML(str) {
			Err = eg.PARAM_INVALID_XML
			goto ERR_RET

		} else if fn == "JSON2SIF" && !isJSON(str) {
			Err = eg.PARAM_INVALID_JSON
			goto ERR_RET
		}
		if Resp, Err = http.Post(url, "application/json", bytes.NewBuffer(args.Data)); Err != nil {
			goto ERR_RET
		}
	}

	if Resp == nil {
		Err = eg.NET_NO_RESPONSE
		goto ERR_RET
	}
	defer Resp.Body.Close()

	if RetData, Err = ioutil.ReadAll(Resp.Body); Err != nil {
		goto ERR_RET
	}

ERR_RET:
	if Err != nil {
		chStr <- ""
		chErr <- Err
		return
	}

	chStr <- string(RetData)
	chErr <- eg.NO_ERROR
	return
}
