package webapi

import (
	"io/ioutil"
	"net/http"

	cmn "github.com/cdutwhu/json-util/common"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	cvt2json "github.com/nsip/n3-sif2json/2JSON"
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
		glb.WDCheck()
		return c.String(http.StatusOK,
			fSf("POST %-40s-> %s\n"+
				"POST %-40s-> %s\n",
				fullIP+route.SIF2JSON, "Upload SIF(XML), return JSON",
				fullIP+route.JSON2SIF, "Upload JSON, return SIF(XML) (Not IMPLEMENTED YET)"))
	})

	path = route.SIF2JSON
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()
		glb.WDCheck()
		if bytes, err := ioutil.ReadAll(c.Request().Body); err == nil {
			sifver := ""
			if ok, ver := url1Value(c.QueryParams(), 0, "sifver"); ok {
				sifver = ver
			}
			json := cvt2json.SIF2JSON("../2JSON/config/SIF2JSON.toml", string(bytes), sifver, false)
			// return c.JSON(http.StatusOK, result{
			// 	Data:  &json,
			// 	Error: "",
			// })
			// return c.JSON(http.StatusOK, json)
			return c.String(http.StatusOK, json)
		}
		return c.JSON(http.StatusBadRequest, result{
			Data:  nil,
			Error: "SIF Data must be provided via Request BODY",
		})
	})

	path = route.JSON2SIF

	e.Start(fSf(":%d", port))
}
