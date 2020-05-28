package webapi

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	eg "github.com/cdutwhu/n3-util/n3errs"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/middleware"
	"github.com/nats-io/nats.go"
	cvt2json "github.com/nsip/n3-sif2json/2JSON"
	cvt2sif "github.com/nsip/n3-sif2json/2SIF"
	cfg "github.com/nsip/n3-sif2json/Server/config"
)

// HostHTTPAsync : Host a HTTP Server for SIF or JSON
func HostHTTPAsync() {
	e := echo.New()
	defer e.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2G"))

	// Add Jaeger Tracer into Middleware
	c := jaegertracing.New(e, nil)
	defer c.Close()

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST},
		AllowCredentials: true,
	}))

	Cfg := env2Struct("Cfg", &cfg.Config{}).(*cfg.Config)
	port := Cfg.WebService.Port
	fullIP := localIP() + fSf(":%d", port)
	route := Cfg.Route
	file := Cfg.File
	mMtx := initMutex(Cfg.Route)

	defer e.Start(fSf(":%d", port))

	// *************************************** List all API, FILE *************************************** //

	path := route.HELP
	e.GET(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		return c.String(http.StatusOK,
			fSf("wget %-55s-> %s\n", fullIP+"/client-linux64", "Get Client(Linux64)")+
				fSf("wget %-55s-> %s\n", fullIP+"/client-mac", "Get Client(Mac)")+
				fSf("wget %-55s-> %s\n", fullIP+"/client-win64", "Get Client(Windows64)")+
				fSf("wget -O config.toml %-40s-> %s\n", fullIP+"/client-config", "Get Client Config")+
				fSf("\n")+
				fSf("POST %-40s-> %s\n"+
					"POST %-40s-> %s\n",
					fullIP+route.SIF2JSON, "Upload SIF(XML), return JSON. [sv]: SIF Spec. Version",
					fullIP+route.JSON2SIF, "Upload JSON, return SIF(XML). [sv]: SIF Spec. Version"))
	})

	// ------------------------------------------------------------------------------------ //

	mRouteRes := map[string]string{
		"/client-linux64": file.ClientLinux64,
		"/client-mac":     file.ClientMac,
		"/client-win64":   file.ClientWin64,
		"/client-config":  file.ClientConfig,
	}

	routeFun := func(rt, res string) func(c echo.Context) error {
		return func(c echo.Context) (err error) {
			if _, err = os.Stat(res); err == nil {
				fPln(rt, res)
				return c.File(res)
			}
			fPf("%v\n", warnOnErr("%v: [%s]  get [%s]", eg.FILE_NOT_FOUND, rt, res))
			return err
		}
	}

	for rt, res := range mRouteRes {
		e.GET(rt, routeFun(rt, res))
	}

	// ------------------------------------------------------------------------------------ //

	path = route.SIF2JSON
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		bytes, err := ioutil.ReadAll(c.Request().Body)
		xmlstr := string(bytes)
		// log("\n%s\n", xmlstr)

		if err != nil || !isXML(xmlstr) {
			return c.JSON(http.StatusBadRequest, result{
				Data:  nil,
				Info:  "",
				Error: err.Error() + " OR Is Request BODY Valid XML?",
			})
		}

		var (
			info      string
			errIntSvr error
		)

		pvalues := c.QueryParams()
		sv, pub2nats := "", false
		if ok, v := url1Value(pvalues, 0, "sv"); ok {
			sv = v
		}
		if ok, n := url1Value(pvalues, 0, "nats"); ok && n != "" {
			pub2nats = true
		}

		// json, svUsed, err := cvt2json.SIF2JSON(cfg.Cfg2JSON, xmlstr, sv, false)

		// Trace [cvt2json.SIF2JSON]
		// [cvt2json.SIF2JSON] uses (variadic parameter), must wrap it to [jaegertracing.TraceFunction]
		results := jaegertracing.TraceFunction(c, func() (string, string, error) {
			return cvt2json.SIF2JSON(Cfg.Cfg2JSON, xmlstr, sv, false)
		})
		json := results[0].Interface().(string)
		svUsed := results[1].Interface().(string)
		if !results[2].IsNil() {
			err = results[2].Interface().(error)
		} else {
			err = nil
		}

		info = "[cvt2json.SIF2JSON]"
		if err != nil {
			errIntSvr = err
			goto ERR_IS
		}

		// send a copy to NATS
		if pub2nats {
			url := Cfg.NATS.URL
			subj := Cfg.NATS.Subject
			timeout := time.Duration(Cfg.NATS.Timeout)

			info += fSf(" | To NATS@Subject: [%s@%s]", url, subj)
			nc, err := nats.Connect(url)
			if err != nil {
				errIntSvr = err
				goto ERR_IS
			}

			msg, err := nc.Request(subj, []byte(json), timeout*time.Millisecond)
			if msg != nil {
				info += fSf(" | NATS responded: [%s]", string(msg.Data))
			}
			if err != nil {
				errIntSvr = err
				goto ERR_IS
			}
		}

	ERR_IS:
		if errIntSvr != nil {
			return c.JSON(http.StatusInternalServerError, result{
				Data:  nil,
				Info:  info,
				Error: errIntSvr.Error(),
			})
		}

		info += fSf(" | SIF Ver: [%s]", svUsed)
		return c.JSON(http.StatusOK, result{
			Data:  &json,
			Info:  info,
			Error: "",
		})

		// return c.String(http.StatusOK, json) // json string is already JSON String, so return String
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
				warnOnErr("%v: \n%s", eg.HTTP_REQBODY_EMPTY, jsonstr)
			}
			if !isJSON(jsonstr) {
				warnOnErr("%v: \n%s", eg.JSON_INVALID, jsonstr)
				goto ERR
			}

			sv := ""
			if ok, ver := url1Value(c.QueryParams(), 0, "sv"); ok {
				sv = ver
			}

			// sif, svUsed, err := cvt2sif.JSON2SIF(cfg.Cfg2SIF, jsonstr, sv)

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
					Data:  nil,
					Info:  "",
					Error: err.Error(),
				})
			}
			return c.JSON(http.StatusOK, result{
				Data:  &sif,
				Info:  svUsed,
				Error: "",
			})
		}
	ERR:
		return c.JSON(http.StatusBadRequest, result{
			Data:  nil,
			Info:  "",
			Error: "JSON Data must be provided via Request BODY as Valid JSON",
		})
	})
}
