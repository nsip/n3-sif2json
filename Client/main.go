package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	cmn "github.com/cdutwhu/json-util/common"
	glb "github.com/nsip/n3-sif2json/Client/global"
)

func main() {
	cmn.FailOnErrWhen(!glb.Init(), "%v", fEf("Config File Init Failed"))
	cmn.SetLog(glb.Cfg.ELog)
	if e := cmn.WarnOnErrWhen(len(os.Args) < 2, "%v", fEf("Need Subcommands: ["+sJoin(getCfgRouteFields(), " ")+"]")); e != nil {
		fPln(e.Error())
		return
	}
	cmn.FailOnErrWhen(!initMapFnURL(glb.Cfg.Server.Protocol, glb.Cfg.Server.IP, glb.Cfg.Server.Port), "%v", fEf("initMapFnURL failed"))

	done := make(chan bool)

	go func() {
		var (
			resp *http.Response = nil
			err  error          = nil
		)

		url := mFnURL[os.Args[1]] // http://ip:port/
		cmd := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
		iPtr := cmd.String("i", "", "Path of original SIF/JSON file to be uploaded")
		vPtr := cmd.String("v", "", "SIF Version (optional), format like (1.2.3)")
		fPtr := cmd.Bool("f", false, "full dump flag: Print INFO & ERROR")     // true: print INFO & ERROR from Server
		nPtr := cmd.Bool("n", false, "indicate server to send a copy to NATS") // true: indicate server
		cmd.Parse(os.Args[2:])

		psV, psN := "", ""
		if *vPtr != "" {
			psV = fSf("sv=%s", *vPtr)
		}
		if *nPtr {
			psN = fSf("nats=true")
		}
		url = fSf("%s?%s&%s", url, psV, psN)
		url = sReplace(url, "?&", "?", 1)
		url = sTrimRight(url, "?&")

		if *fPtr {
			fPln("accessing ... " + url)
			fPln("-----------------------------")
		}

		switch os.Args[1] { // Config - Route - each Field
		case "API":
			resp, err = http.Get(url)

		case "SIF2JSON", "JSON2SIF":
			cmn.FailOnErrWhen(*iPtr == "", "%v", fEf("[-i] must be provided"))
			data, err := ioutil.ReadFile(*iPtr)
			cmn.FailOnErr("%v: %v", err, "Is [-i] provided correctly?")
			str := string(data)

			if os.Args[1] == "SIF2JSON" {
				cmn.FailOnErrWhen(!cmn.IsXML(str), "%v Abort", fEf("input file is invalid XML,"))
			} else {
				cmn.FailOnErrWhen(!cmn.IsJSON(str), "%v About", fEf("input file is invalid JSON,"))
			}
			resp, err = http.Post(url, "application/json", bytes.NewBuffer(data))

		default:
			if e := cmn.WarnOnErr("%v", fEf("Unsupported Subcommand: %v", os.Args[1])); e != nil {
				fPln(e.Error())
				done <- true
				return
			}
		}

		cmn.FailOnErrWhen(resp == nil, "HTTP Access Fatal: %v OR %v", err, fEf("Couldn't get Response."), url)
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		cmn.FailOnErr("resp Body fatal: %v", err)
		if data != nil {
			if os.Args[1] == "API" {
				fPt(string(data))
			} else {
				m := make(map[string]interface{})
				cmn.FailOnErr("json.Unmarshal ... %v", json.Unmarshal(data, &m))

				if *fPtr {
					if m["info"] != nil && m["info"] != "" {
						fPf("INFO: %v\n", m["info"])
					}
					fPln("-----------------------------")
					if m["error"] != nil && m["error"] != "" {
						fPf("ERROR: %v\n", m["error"])
					}
					fPln("-----------------------------")
				}
				if m["data"] != nil && m["data"] != "" {
					fPf("%s\n", m["data"])
				}
			}
		}

		done <- true
	}()

	select {
	case <-time.After(time.Duration(glb.Cfg.Access.Timeout) * time.Second):
		cmn.FailOnErr("%v", fEf("Didn't Get Response in time. %d(s)", glb.Cfg.Access.Timeout))
	case <-done:
	}
}
