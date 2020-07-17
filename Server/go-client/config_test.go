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
	ICfg := env2Struct("Cfg-Clt-S2J", &Config{}) //
	spew.Dump(ICfg.(*Config))
}
