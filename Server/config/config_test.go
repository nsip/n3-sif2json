package config

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestLoad(t *testing.T) {
	cfg := newCfg("./config.toml")
	spew.Dump(cfg)
}

func TestInit(t *testing.T) {
	InitEnvVarFromTOML("Cfg", "./config.toml")
	ICfg, err := env2Struct("Cfg", &Config{})
	failOnErr("%v", err)
	cfg := ICfg.(*Config)
	spew.Dump(cfg)
}
