package main

import (
	"flag"
	"os"
	"testing"

	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/cdutwhu/n3-util/n3err"
	"github.com/davecgh/go-spew/spew"
)

func TestMain(t *testing.T) {
	main()
}

func TestLoad(t *testing.T) {
	cfg := &Config{}
	failOnErrWhen(n3cfg.New(cfg, nil, "./config.toml") == "", "%v", n3err.CFG_INIT_ERR)
	spew.Dump(cfg)
}

func TestInit(t *testing.T) {
	cfg := &Config{}
	failOnErr("%v", n3cfg.InitEnvVar(cfg, nil, "Cfg", "./config.toml"))
	ICfg := env2Struct("Cfg", cfg)
	spew.Dump(ICfg.(*Config))
}

// ****************************** //
// Create a copy of config for Client. Excluding some attributes.
// Once building, move it to Client config Directory.
func TestGenCltCfg(t *testing.T) {
	cfg := &Config{}
	failOnErrWhen(
		n3cfg.New(cfg,
			map[string]string{
				"[v]":    "Version",
				"[s]":    "Service",
				"[port]": "WebService.Port",
			}, "./config.toml") == "",
		"%v",
		n3err.CFG_INIT_ERR,
	)

	temp := "./temp.toml"
	defer func() { os.Remove(temp) }()

	n3cfg.Save(temp, cfg)
	if !flag.Parsed() {
		flag.Parse()
	}
	n3cfg.SelCfgAttrL1(temp, "./goclient/config.toml", flag.Args()...)
}
