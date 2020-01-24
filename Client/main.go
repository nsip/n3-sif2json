package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	cmn "github.com/cdutwhu/json-util/common"
	glb "github.com/nsip/n3-sif2json/Client/global"
)

func main() {
	cmn.FailOnErrWhen(!glb.Init(), "%v", fEf("Config File Init Failed"))
	cmn.SetLog(glb.Cfg.ELog)
	if err := cmn.WarnOnErrWhen(len(os.Args) < 2, "%v", fEf("Need Subcommands: ["+sJoin(getCfgRouteFields(), " ")+"]")); err != nil {
		fPln(err.Error())
		return
	}
	cmn.FailOnErrWhen(!initMapFnURL(glb.Cfg.Server.Protocol, glb.Cfg.Server.IP, glb.Cfg.Server.Port), "%v", fEf("initMapFnURL failed"))

	fPln(mFnURL)

	timeout := time.After(time.Duration(glb.Cfg.Access.Timeout) * time.Second)
	done := make(chan bool)

	go func() {
		var resp *http.Response = nil
		var err error = nil
		var data []byte = nil
		url := mFnURL[os.Args[1]]
		cmd := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
		sv := cmd.String("sv", "", "SIF Version, like (1.2.3)")
		cmd.Parse(os.Args[2:])

		switch os.Args[1] {
		case "ROOT":
			fPln("hello")
		default:
			cmn.FailOnErr("%v", fEf("unknown subcommand: %v", os.Args[1]))
		}

		fPln(resp, err, data, url, cmd, sv)
	}()

	select {
	case <-timeout:
		cmn.FailOnErr("%v", fEf("Didn't Get Server Response in time. %d(s)", glb.Cfg.Access.Timeout))
	case <-done:
	}

}
