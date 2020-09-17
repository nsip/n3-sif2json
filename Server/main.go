package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"time"

	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/cdutwhu/n3-util/n3cfg/attrim"
	"github.com/cdutwhu/n3-util/n3cfg/strugen"
	"github.com/cdutwhu/n3-util/n3err"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"
	cvt2json "github.com/nsip/n3-sif2json/2JSON"
	cvt2sif "github.com/nsip/n3-sif2json/2SIF"
	cfg "github.com/nsip/n3-sif2json/Config/cfg"
)

func mkCfg4Clt(cfg interface{}) {
	forel := "./config_rel.toml"
	n3cfg.Save(forel, cfg)
	outoml := "./client/config.toml"
	outsrc := "./client/config.go"
	os.Remove(outoml)
	os.Remove(outsrc)
	attrim.SelCfgAttrL1(forel, outoml, "Service", "Route", "Server", "Access")
	strugen.GenStruct(outoml, "Config", "client", outsrc)
	strugen.GenNewCfg(outsrc)
}

func main() {
	// Load global config.toml file from Config/
	n3cfg.SetDftCfgVal("n3-sif2json", "0.0.0")
	pCfg := cfg.NewCfg(
		"Config",
		map[string]string{
			"[s]":    "Service",
			"[v]":    "Version",
			"[port]": "WebService.Port",
		},
		"./Config/config.toml",
		"../Config/config.toml",
	)
	failOnErrWhen(pCfg == nil, "%v: Config Init Error", n3err.CFG_INIT_ERR)
	Cfg := pCfg.(*cfg.Config)

	// Trim a shorter config toml file for client package
	if len(os.Args) > 2 && os.Args[2] == "trial" {
		mkCfg4Clt(Cfg)
		return
	}

	ws := Cfg.WebService
	var IService interface{} = Cfg.Service // Cfg.Service can be "string", can be "interface{}"
	service := IService.(string)

	// Set Jaeger Env for tracing
	os.Setenv("JAEGER_SERVICE_NAME", service)
	os.Setenv("JAEGER_SAMPLER_TYPE", "const")
	os.Setenv("JAEGER_SAMPLER_PARAM", "1")

	// Set LOGGLY
	setLoggly(true, Cfg.Loggly.Token, service)

	// Set Log Options
	syncBindLog(true)
	enableWarnDetail(false)
	enableLog2F(true, Cfg.Log)

	logGrp.Do(fSf("local log file @ [%s]", Cfg.Log))
	logGrp.Do(fSf("[%s] Hosting on: [%v:%d], version [%v]", service, localIP(), ws.Port, Cfg.Version))

	// Start Service
	done := make(chan string)
	c := make(chan os.Signal)
	signal.Notify(c, os.Kill, os.Interrupt)
	go HostHTTPAsync(c, done)
	logGrp.Do(<-done)
}

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
		Cfg    = rflx.Env2Struct("Config", &cfg.Config{}).(*cfg.Config)
		port   = Cfg.WebService.Port
		fullIP = localIP() + fSf(":%d", port)
		route  = Cfg.Route
		mMtx   = initMutex(&Cfg.Route)
	)

	defer e.Start(fSf(":%d", port))
	logGrp.Do("Echo Service is Starting ...")

	// *************************************** List all API, FILE *************************************** //

	path := route.Help
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
				fullIP+route.ToJSON, "Upload SIF(XML), return JSON. [sv]: SIF Spec. Version",
				fullIP+route.ToSIF, "Upload JSON, return SIF(XML). [sv]: SIF Spec. Version"))
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

	path = route.ToJSON
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

	path = route.ToSIF
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
