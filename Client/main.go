package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

	timeout := time.After(time.Duration(glb.Cfg.Access.Timeout) * time.Second)
	done := make(chan bool)

	go func() {
		var resp *http.Response = nil
		var err error = nil
		var data []byte = nil
		url := mFnURL[os.Args[1]] // http://ip:port/

		cmd := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
		iPtr := cmd.String("i", "", "Path of original SIF/JSON file to be uploaded")
		vPtr := cmd.String("v", "", "SIF Version (optional), format like (1.2.3)")
		oPtr := cmd.String("o", "", "Path of outcome file (optional)")
		// infoPtr := cmd.String("info", "", "Only dump the request info (optional)")
		cmd.Parse(os.Args[2:])

		if *vPtr != "" {
			url += fSf("?sv=%s", *vPtr)
		}
		// fPln("accessing ... " + url)

		if *oPtr != "" {
			dir := filepath.Dir(*oPtr)
			if _, err = os.Stat(dir); os.IsNotExist(err) {
				os.Mkdir(dir, os.ModePerm)
			}
		}

		switch os.Args[1] { // Config - Route - each Field
		case "API":
			resp, err = http.Get(url)

		case "SIF2JSON", "JSON2SIF":
			cmn.FailOnErrWhen(*iPtr == "", "%v", fEf("[-i] must be provided"))
			data, err := ioutil.ReadFile(*iPtr)
			cmn.FailOnErr("%v: %v", err, "Is [-i] provided correctly?")
			if os.Args[1] == "SIF2JSON" {
				cmn.FailOnErrWhen(!cmn.IsXML(string(data)), "%v", fEf("input file is not valid XML, Abort"))
				if *oPtr != "" && !sHasSuffix(*oPtr, ".json") {
					*oPtr += ".json"
				}
			} else {
				cmn.FailOnErrWhen(!cmn.IsJSON(string(data)), "%v", fEf("input file is not valid JSON, Abort"))
				if *oPtr != "" && !sHasSuffix(*oPtr, ".xml") {
					*oPtr += ".xml"
				}
			}
			resp, err = http.Post(url, "application/json", bytes.NewBuffer(data))

		default:
			if e := cmn.WarnOnErr("%v", fEf("Unsupported Subcommand: %v", os.Args[1])); e != nil {
				fPln(e.Error())
				goto QUIT
			}
		}

		cmn.FailOnErrWhen(resp == nil, "HTTP Access Fatal: %v OR %v", err, fEf("Couldn't get Response"))
		defer resp.Body.Close()
		data, err = ioutil.ReadAll(resp.Body)
		cmn.FailOnErr("resp Body fatal: %v", err)
		if data != nil {
			if os.Args[1] == "API" {
				fPt(string(data))
			} else {
				m := make(map[string]interface{})
				cmn.FailOnErr("json.Unmarshal ... %v", json.Unmarshal(data, &m))
				cmn.FailOnErrWhen(m["error"] != nil && m["error"] != "", "%v", fEf("ERROR: %v\n", m["error"]))

				if m["info"] != nil && m["info"] != "" {
					fPf("INFO: %v\n", m["info"])
				}

				fPln(" ----------------------------- ")

				if m["data"] != nil && m["data"] != "" {
					if *oPtr != "" {
						ioutil.WriteFile(*oPtr, []byte(m["data"].(string)), 0666)
					}
					fPf("%s\n", m["data"])
				}
			}
		}

	QUIT:
		done <- true
	}()

	select {
	case <-timeout:
		cmn.FailOnErr("%v", fEf("Didn't Get Response in time. %d(s)", glb.Cfg.Access.Timeout))
	case <-done:
	}
}
