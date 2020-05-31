package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	eg "github.com/cdutwhu/n3-util/n3errs"
	clt "github.com/nsip/n3-sif2json/Server/go-client"
)

func main() {
	fns := structFields(clt.Config{}.Route)
	failOnErrWhen(len(os.Args) < 3, "%v: need [config.toml] %v", eg.CLI_SUBCMD_ERR, fns)

	cltcfg, fn := os.Args[1], os.Args[2]

	cmd := flag.NewFlagSet(fn, flag.ExitOnError)
	iPtr := cmd.String("i", "", "Path of original SIF/JSON file to be uploaded")
	vPtr := cmd.String("v", "", "SIF Version (optional), format like (1.2.3)")
	wPtr := cmd.Bool("w", false, "whole dump flag: Print INFO & ERROR")    // true: print INFO & ERROR from Server
	nPtr := cmd.Bool("n", false, "indicate server to send a copy to NATS") // true: indicate server
	cmd.Parse(os.Args[3:])

	data, err := ioutil.ReadFile(*iPtr)
	failOnErrWhen(fn == "SIF2JSON" || fn == "JSON2SIF", "%v: %s", err, *iPtr)

	str, err := clt.DO(
		cltcfg,
		fn,
		clt.Args{
			Data:   data,
			Ver:    *vPtr,
			ToNATS: *nPtr,
		})
	failOnErr("Access SIF2JSON Service Failed: %v", err)

	if fn == "HELP" {
		fPt(str)
		return
	}

	m := make(map[string]interface{})
	failOnErr("json.Unmarshal ... %v", json.Unmarshal([]byte(str), &m))
	if *wPtr {
		fPf("INFO: %v\n", m["info"])
		fPln("-----------------------------")
		fPf("ERROR: %v\n", m["error"])
		fPln("-----------------------------")
	}
	fPf("%s\n", m["data"])
}
