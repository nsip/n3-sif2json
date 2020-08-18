package main

import (
	"os"
	"os/signal"

	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/cdutwhu/n3-util/n3err"
)

func main() {
	Cfg := n3cfg.ToEnvN3sif2jsonServer(map[string]string{
		"[s]": "Service",
		"[v]": "Version",
	}, envKey)
	failOnErrWhen(Cfg == nil, "%v: Config Init Error", n3err.CFG_INIT_ERR)

	ws, service := Cfg.WebService, Cfg.Service.(string)
	os.Setenv("JAEGER_SERVICE_NAME", service)
	os.Setenv("JAEGER_SAMPLER_TYPE", "const")
	os.Setenv("JAEGER_SAMPLER_PARAM", "1")

	// --- LOGGLY --- //
	setLoggly(true, Cfg.Loggly.Token, service)

	enableWarnDetail(false)
	enableLog2F(true, Cfg.Log)
	logBind(logger, loggly("info")).Do(fSf("local log file @ [%s]", Cfg.Log))
	logBind(logger, loggly("info")).Do(fSf("[%s] Hosting on: [%v:%d], version [%v]", service, localIP(), ws.Port, Cfg.Version))

	done := make(chan string)
	c := make(chan os.Signal)
	signal.Notify(c, os.Kill, os.Interrupt)
	go HostHTTPAsync(c, done)

	logBind(logger, loggly("info")).Do(<-done)
}
