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
		svPtr := cmd.String("sv", "", "SIF Version (optional), format like (1.2.3)")
		// sifPtr := cmd.String("sif", "", "the path of SIF file to be uploaded")
		sifPtr, sifName := cmd.String("sif", "", "the path of SIF file to be uploaded"), ""
		// jsonPtr := cmd.String("json", "", "the path of JSON file to be uploaded")
		cmd.Parse(os.Args[2:])

		switch os.Args[1] { // Config - Route - each Field
		// case "API":
		// 	// fPln("accessing ... " + url)
		// 	// resp, err = http.Get(url)
		// 	goto QUIT

		case "SIF2JSON":
			if *svPtr != "" {
				url += fSf("?sv=%s", *svPtr)
			}
			// fPln("accessing ... " + url)
			cmn.FailOnErrWhen(*sifPtr == "", "%v", fEf("[-sif] must be provided"))
			sif, err := ioutil.ReadFile(*sifPtr)
			cmn.FailOnErr("%v: %v", err, "Is [-sif] provided correctly?")
			cmn.FailOnErrWhen(!cmn.IsXML(string(sif)), "%v", fEf("sif is not valid XML file, abort"))
			sifName = filepath.Base(*sifPtr)
			resp, err = http.Post(url, "application/json", bytes.NewBuffer(sif))

		case "JSON2SIF":
			cmn.FailOnErr("%v", fEf("JSON2SIF is not implemented"))

		default:
			if e := cmn.WarnOnErr("%v", fEf("Unsupported Subcommand: %v", os.Args[1])); e != nil {
				fPln(e.Error())
				goto QUIT
			}
		}

		cmn.FailOnErr("http access fatal: %v", err)
		defer resp.Body.Close()
		data, err = ioutil.ReadAll(resp.Body)
		cmn.FailOnErr("resp Body fatal: %v", err)
		if data != nil {
			m := make(map[string]interface{})
			cmn.FailOnErr("json.Unmarshal ... %v", json.Unmarshal(data, &m))
			cmn.FailOnErrWhen(m["error"] != nil && m["error"] != "", "%v", fEf("ERROR: %v\n", m["error"]))
			if m["info"] != nil && m["info"] != "" {
				fPf("INFO: %v\n", m["info"])
			}
			if m["data"] != nil && m["data"] != "" {
				path := fSf("./data/%s.json", sifName)
				fPf("Saving to %s\n", path)
				ioutil.WriteFile(path, []byte(m["data"].(string)), 0666)
				fPf("%s\n", m["data"])
			}
		}

	QUIT:
		done <- true
	}()

	select {
	case <-timeout:
		cmn.FailOnErr("%v", fEf("Didn't Get Server Response in time. %d(s)", glb.Cfg.Access.Timeout))
	case <-done:
	}
}
