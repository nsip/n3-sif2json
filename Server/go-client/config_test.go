package client

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestLoad(t *testing.T) {
	cfg := newCfg("./config.toml")
	spew.Dump(cfg)
}

func TestInit(t *testing.T) {
	initEnvVarFromTOML("Cfg-Clt-S2J", "./config.toml")
	ICfg, err := env2Struct("Cfg-Clt-S2J", &Config{}) //
	failOnErr("%v", err)
	spew.Dump(ICfg.(*Config))
}
