package main

import (
	"fmt"
	"log"

	g "github.com/nsip/n3-sif2json/Server/global"
	api "github.com/nsip/n3-sif2json/Server/webapi"
)

func main() {
	failOnErrWhen(!g.Init(), "%v", fmt.Errorf("Global Config Init Error"))
	log.Printf("Working on: [%v]", g.Cfg.WebService)
	done := make(chan string)
	go api.HostHTTPAsync()
	<-done
}
