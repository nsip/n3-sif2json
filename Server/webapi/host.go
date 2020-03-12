package webapi

import (
	"io/ioutil"
	"net/http"
	"time"

	cmn "github.com/cdutwhu/json-util/common"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nats-io/nats.go"
	cvt2json "github.com/nsip/n3-sif2json/2JSON"
	cvt2sif "github.com/nsip/n3-sif2json/2SIF"
	glb "github.com/nsip/n3-sif2json/Server/global"
)

// HostHTTPAsync : Host a HTTP Server for SIF or JSON
func HostHTTPAsync() {
	e := echo.New()

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

	port := glb.Cfg.WebService.Port
	fullIP := cmn.LocalIP() + fSf(":%d", port)
	route := glb.Cfg.Route
	initMutex()

	path := "/"
	e.GET(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		return c.String(http.StatusOK,
			fSf("POST %-40s-> %s\n"+
				"POST %-40s-> %s\n",
				fullIP+route.SIF2JSON, "Upload SIF(XML), return JSON. [sv]: SIF Spec. Version",
				fullIP+route.JSON2SIF, "Upload JSON, return SIF(XML). [sv]: SIF Spec. Version"))
	})

	path = route.SIF2JSON
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		bytes, err := ioutil.ReadAll(c.Request().Body)
		if err != nil || !cmn.IsXML(string(bytes)) {
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

		json, svUsed, err := cvt2json.SIF2JSON(glb.Cfg.Cfg2JSON, string(bytes), sv, false)
		info = "[cvt2json.SIF2JSON]"
		if err != nil {
			errIntSvr = err
			goto ERR_IS
		}

		// send a copy to NATS
		if pub2nats {
			url := glb.Cfg.NATS.URL
			subj := glb.Cfg.NATS.Subject
			timeout := time.Duration(glb.Cfg.NATS.Timeout)

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

	path = route.JSON2SIF
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		if bytes, err := ioutil.ReadAll(c.Request().Body); err == nil {
			if !cmn.IsJSON(string(bytes)) {
				goto ERR
			}
			sv := ""
			if ok, ver := url1Value(c.QueryParams(), 0, "sv"); ok {
				sv = ver
			}
			sif, svUsed, err := cvt2sif.JSON2SIF(glb.Cfg.Cfg2SIF, string(bytes), sv)
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

	e.Start(fSf(":%d", port))
}
