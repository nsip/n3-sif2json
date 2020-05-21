package main

import (
	"fmt"

	g "github.com/nsip/n3-sif2json/Server/global"
	api "github.com/nsip/n3-sif2json/Server/webapi"
)

func main() {
	failOnErrWhen(!g.Init(), "%v", fmt.Errorf("Global Config Init Error"))

	cfg := g.Cfg
	ws, logfile, servicename := cfg.WebService, cfg.LogFile, cfg.ServiceName

	setLog(logfile)
	fPln(logWhen(true, "[%s] Hosting on: [%v:%d], version [%v]", servicename, localIP(), ws.Port, ws.Version))

	done := make(chan string)
	go api.HostHTTPAsync()
	<-done
}
