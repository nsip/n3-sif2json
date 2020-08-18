package goclient

import (
	"io/ioutil"
	"testing"

	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/davecgh/go-spew/spew"
)

func TestLoad(t *testing.T) {
	cfg := n3cfg.ToEnvN3sif2jsonGoclient(nil, "TestKey", "./config.toml")
	spew.Dump(cfg)
}

func TestInit(t *testing.T) {
	cfg := n3cfg.ToEnvN3sif2jsonGoclient(nil, "TestKey", "./config.toml")
	spew.Dump(cfg)
	fPln(" ------------------------------- ")
	cfg1 := n3cfg.FromEnvN3sif2jsonGoclient("TestKey")
	spew.Dump(cfg1)
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
