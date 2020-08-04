package goclient

import (
	"io/ioutil"
	"testing"

	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/davecgh/go-spew/spew"
)

func TestLoad(t *testing.T) {
	cfg := &Config{}
	n3cfg.New(cfg, nil, "./config.toml")
	spew.Dump(cfg)
}

func TestInit(t *testing.T) {
	cfg := &Config{}
	n3cfg.InitEnvVar(cfg, nil, "Cfg-Clt-S2J", "./config.toml")
	spew.Dump(cfg)
	fPln(" ------------------------------- ")
	ICfg := env2Struct("Cfg-Clt-S2J", &Config{}) //
	spew.Dump(ICfg.(*Config))
}

func TestDO(t *testing.T) {

	str, err := DO(
		"./config.toml",
		"HELP",
		nil,
	)
	fPln(str)
	fPln(err)
	fPln(" ------------------------------------ ")

	bytes, err := ioutil.ReadFile("../../data/examples347/NAPTest_0.xml")
	failOnErr("%v", err)
	str, err = DO(
		"./config.toml",
		"SIF2JSON",
		&Args{
			Data:   bytes,
			Ver:    "3.4.7",
			ToNATS: false,
		},
	)
	fPln(str)
	fPln(err)
	ioutil.WriteFile("./out.json", []byte(str), 0666)
}
