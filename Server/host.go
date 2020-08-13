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
	defer func() { logBind(logger, loggly("info")).Do("HostHTTPAsync Exit") }()

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

	Cfg := n3cfg.FromEnvN3sif2jsonServer("Cfg")
	port := Cfg.WebService.Port
	fullIP := localIP() + fSf(":%d", port)
	route := Cfg.Route
	mMtx := initMutex(&Cfg.Route)

	defer e.Start(fSf(":%d", port))
	logBind(logger, loggly("info")).Do("Echo Service is Starting")

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

	// ------------------------------------------------------------------------------------ //

	path = route.SIF2JSON
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		var (
			status  = http.StatusOK
			errSvr  error
			infoSvr string
			jsonRet string

			results []reflect.Value
			svUsed  string
		)

		pvalues := c.QueryParams()
		sv, pub2nats := "", false
		if ok, v := url1Value(pvalues, 0, "sv"); ok {
			sv = v
		}
		if ok, n := url1Value(pvalues, 0, "nats"); ok && n != "" {
			pub2nats = true
		}

		infoSvr = "Read Request Body"
		bytes, errSvr := ioutil.ReadAll(c.Request().Body)
		if errSvr != nil {
			status = http.StatusInternalServerError
			goto ERR
		}
		if !isXML(string(bytes)) {
			errSvr = n3err.PARAM_INVALID_XML
			status = http.StatusBadRequest
			goto ERR
		}

		infoSvr = "[cvt2json.SIF2JSON]"
		// jsonRet, svUsed, errSvr = cvt2json.SIF2JSON(Cfg.Cfg2JSON, string(bytes), sv, false)

		// Trace [cvt2json.SIF2JSON]
		// [cvt2json.SIF2JSON] uses (variadic parameter), must wrap it to [jaegertracing.TraceFunction]
		results = jaegertracing.TraceFunction(c, func() (string, string, error) {
			return cvt2json.SIF2JSON(Cfg.Cfg2JSON, string(bytes), sv, false)
		})
		jsonRet = results[0].Interface().(string)
		svUsed = results[1].Interface().(string)
		if !results[2].IsNil() {
			errSvr = results[2].Interface().(error)
			status = http.StatusInternalServerError
			goto ERR
		}

		// send a copy to NATS
		if pub2nats {
			url, subj, timeout := Cfg.NATS.URL, Cfg.NATS.Subject, time.Duration(Cfg.NATS.Timeout)

			infoSvr += fSf(" | To NATS@Subject: [%s@%s]", url, subj)
			nc, errSvr := nats.Connect(url)
			if errSvr != nil {
				status = http.StatusInternalServerError
				goto ERR
			}

			msg, errSvr := nc.Request(subj, []byte(jsonRet), timeout*time.Millisecond)
			if msg != nil {
				infoSvr += fSf(" | NATS responded: [%s]", string(msg.Data))
			}
			if errSvr != nil {
				status = http.StatusInternalServerError
				goto ERR
			}
		}

	ERR:
		if errSvr != nil {
			return c.JSON(status, result{
				Data:  "",
				Info:  infoSvr,
				Error: errSvr.Error(),
			})
		}

		return c.JSON(status, result{
			Data:  jsonRet,
			Info:  infoSvr + fSf(" | SIF Ver: [%s]", svUsed),
			Error: "",
		})

		// return c.String(http.StatusOK, jsonRet) // jsonRet is already JSON String, so return String
	})

	// ------------------------------------------------------------------------------------ //

	path = route.JSON2SIF
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		if bytes, err := ioutil.ReadAll(c.Request().Body); err == nil {
			jsonstr := string(bytes)
			// log("\n%s\n", jsonstr)

			if len(bytes) == 0 {
				warnOnErr("%v: \n%s", n3err.HTTP_REQBODY_EMPTY, jsonstr)
			}
			if !isJSON(jsonstr) {
				warnOnErr("%v: \n%s", n3err.JSON_INVALID, jsonstr)
				goto ERR
			}

			sv := ""
			if ok, ver := url1Value(c.QueryParams(), 0, "sv"); ok {
				sv = ver
			}

			// sif, svUsed, err := cvt2sif.JSON2SIF(Cfg.Cfg2SIF, jsonstr, sv)

			// Trace [cvt2sif.JSON2SIF]
			results := jaegertracing.TraceFunction(c, cvt2sif.JSON2SIF, Cfg.Cfg2SIF, jsonstr, sv)
			sif := results[0].Interface().(string)
			svUsed := results[1].Interface().(string)
			if !results[2].IsNil() {
				err = results[2].Interface().(error)
			} else {
				err = nil
			}

			if err != nil {
				return c.JSON(http.StatusInternalServerError, result{
					Data:  "",
					Info:  "",
					Error: err.Error(),
				})
			}
			return c.JSON(http.StatusOK, result{
				Data:  sif,
				Info:  svUsed,
				Error: "",
			})
		}
	ERR:
		return c.JSON(http.StatusBadRequest, result{
			Data:  "",
			Info:  "",
			Error: "JSON Data must be provided via Request BODY as Valid JSON",
		})
	})
}
