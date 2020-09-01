package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/cdutwhu/n3-util/n3err"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"
	cvt2json "github.com/nsip/n3-sif2json/2JSON"
	cvt2sif "github.com/nsip/n3-sif2json/2SIF"
)

func shutdownAsync(e *echo.Echo, sig <-chan os.Signal, done chan<- string) {
	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	failOnErr("%v", e.Shutdown(ctx))
	time.Sleep(20 * time.Millisecond)
	done <- "Shutdown Successfully"
}

// HostHTTPAsync : Host a HTTP Server for SIF or JSON
func HostHTTPAsync(sig <-chan os.Signal, done chan<- string) {
	defer func() { logGrp.Do("HostHTTPAsync Exit") }()

	e := echo.New()
	defer e.Close()

	// waiting for shutdown
	go shutdownAsync(e, sig, done)

	// Add Jaeger Tracer into Middleware
	c := jaegertracing.New(e, nil)
	defer c.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2G"))
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST},
		AllowCredentials: true,
	}))

	e.Logger.SetOutput(os.Stdout)
	e.Logger.Infof(" ------------------------ e.Logger.Infof ------------------------ ")

	var (
		Cfg    = n3cfg.FromEnvN3sif2jsonServer(envKey)
		port   = Cfg.WebService.Port
		fullIP = localIP() + fSf(":%d", port)
		route  = Cfg.Route
		mMtx   = initMutex(&Cfg.Route)
	)

	defer e.Start(fSf(":%d", port))
	logGrp.Do("Echo Service is Starting ...")

	// *************************************** List all API, FILE *************************************** //

	path := route.HELP
	e.GET(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		return c.String(http.StatusOK,
			// fSf("wget %-55s-> %s\n", fullIP+"/client-linux64", "Get Client(Linux64)")+
			// 	fSf("wget %-55s-> %s\n", fullIP+"/client-mac", "Get Client(Mac)")+
			// 	fSf("wget %-55s-> %s\n", fullIP+"/client-win64", "Get Client(Windows64)")+
			// 	fSf("wget -O config.toml %-40s-> %s\n", fullIP+"/client-config", "Get Client Config")+
			// 	fSf("\n")+
			fSf("POST %-40s-> %s\n"+
				"POST %-40s-> %s\n",
				fullIP+route.SIF2JSON, "Upload SIF(XML), return JSON. [sv]: SIF Spec. Version",
				fullIP+route.JSON2SIF, "Upload JSON, return SIF(XML). [sv]: SIF Spec. Version"))
	})

	// ------------------------------------------------------------------------------------ //

	// mRouteRes := map[string]string{
	// 	"/client-linux64": Cfg.File.ClientLinux64,
	// 	"/client-mac":     Cfg.File.ClientMac,
	// 	"/client-win64":   Cfg.File.ClientWin64,
	// 	"/client-config":  Cfg.File.ClientConfig,
	// }

	// routeFun := func(rt, res string) func(c echo.Context) error {
	// 	return func(c echo.Context) (err error) {
	// 		if _, err = os.Stat(res); err == nil {
	// 			fPln(rt, res)
	// 			return c.File(res)
	// 		}
	// 		return warnOnErr("%v: [%s]  get [%s]", n3err.FILE_NOT_FOUND, rt, res)
	// 	}
	// }

	// for rt, res := range mRouteRes {
	// 	e.GET(rt, routeFun(rt, res))
	// }

	// ------------------------------------------------------------------------------------------------------------- //
	// ------------------------------------------------------------------------------------------------------------- //

	path = route.SIF2JSON
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		var (
			status  = http.StatusOK
			ret     string
			results []reflect.Value
		)

		logGrp.Do("Parsing Params")
		pvalues, sv, msg := c.QueryParams(), "", false
		if ok, v := url1Value(pvalues, 0, "sv"); ok {
			sv = v
		}
		if ok, n := url1Value(pvalues, 0, "nats"); ok && n != "" && n != "false" {
			msg = true
		}

		logGrp.Do("Reading Request Body")
		bytes, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			status = http.StatusInternalServerError
			ret = err.Error() + " @ Read Request Body"
			goto RET
		}
		if len(bytes) == 0 {
			status = http.StatusBadRequest
			ret = n3err.HTTP_REQBODY_EMPTY.Error() + " @ Read Request Body"
			goto RET
		}
		if !isXML(string(bytes)) {
			status = http.StatusBadRequest
			ret = n3err.PARAM_INVALID_XML.Error() + " @ Read Request Body"
			goto RET
		}

		logGrp.Do("cvt2json.SIF2JSON")
		// ret, svUsed, err = cvt2json.SIF2JSON(Cfg.Cfg2JSON, string(bytes), sv, false)
		// Trace [cvt2json.SIF2JSON], uses (variadic parameter), must wrap it to [jaegertracing.TraceFunction]
		results = jaegertracing.TraceFunction(c, func() (string, string, error) {
			return cvt2json.SIF2JSON(string(bytes), sv, false)
		})
		ret = results[0].Interface().(string)
		if !results[2].IsNil() {
			status = http.StatusInternalServerError
			ret = results[2].Interface().(error).Error()
			goto RET
		}
		logGrp.Do(results[1].Interface().(string) + " applied")

		// Send a copy to NATS
		if msg {
			url, subj, timeout := Cfg.NATS.URL, Cfg.NATS.Subject, time.Duration(Cfg.NATS.Timeout)
			nc, err := nats.Connect(url)
			if err != nil {
				status = http.StatusInternalServerError
				ret = err.Error() + fSf(" @NATS Connect @Subject: [%s@%s]", url, subj)
				goto RET
			}
			msg, err := nc.Request(subj, []byte(ret), timeout*time.Millisecond)
			if err != nil {
				status = http.StatusInternalServerError
				ret = err.Error() + fSf(" @NATS Request @Subject: [%s@%s]", url, subj)
				goto RET
			}
			logGrp.Do(string(msg.Data))
		}

	RET:
		if status != http.StatusOK {
			warnGrp.Do(ret + " --> Failed")
		} else {
			logGrp.Do("--> Finish SIF2JSON")
		}
		return c.String(status, ret) // ret is already JSON String, so return String
	})

	// ------------------------------------------------------------------------------------------------------------- //
	// ------------------------------------------------------------------------------------------------------------- //

	path = route.JSON2SIF
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		var (
			status  = http.StatusOK
			ret     string
			results []reflect.Value
		)

		logGrp.Do("Parsing Params")
		pvalues, sv := c.QueryParams(), ""
		if ok, v := url1Value(pvalues, 0, "sv"); ok {
			sv = v
		}

		logGrp.Do("Reading Body")
		bytes, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			status = http.StatusInternalServerError
			ret = err.Error() + " @ Read Request Body"
			goto RET
		}
		if len(bytes) == 0 {
			status = http.StatusBadRequest
			ret = n3err.HTTP_REQBODY_EMPTY.Error() + " @ Read Request Body"
			goto RET
		}
		if !isJSON(string(bytes)) {
			status = http.StatusBadRequest
			ret = n3err.PARAM_INVALID_JSON.Error() + " @ Read Request Body"
			goto RET
		}

		logGrp.Do("cvt2json.JSON2SIF")
		// ret, svUsed, err := cvt2sif.JSON2SIF(Cfg.Cfg2SIF, string(bytes), sv)
		// Trace [cvt2sif.JSON2SIF]
		results = jaegertracing.TraceFunction(c, cvt2sif.JSON2SIF, string(bytes), sv)
		ret = results[0].Interface().(string)
		if !results[2].IsNil() {
			status = http.StatusInternalServerError
			ret = results[2].Interface().(error).Error()
			goto RET
		}
		logGrp.Do(results[1].Interface().(string) + " applied")

	RET:
		if status != http.StatusOK {
			warnGrp.Do(ret + " --> Failed")
		} else {
			logGrp.Do("--> Finish JSON2SIF")
		}
		return c.String(status, ret)
	})
}
