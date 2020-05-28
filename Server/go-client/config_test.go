package client

import "testing"

func TestLoad(t *testing.T) {
	cfg := newCfg("./config.toml")
	fPln(cfg.Path)
	fPln(cfg.LogFile)
	fPln(cfg.Access)
	fPln(cfg.Server)
	fPln(cfg.Route)
	fPln(cfg.ServiceName)
}

func TestInit(t *testing.T) {
	initEnvVarFromTOML("Cfg", "./config.toml")
	cfg := env2Struct("Cfg", &config{}).(*config)
	fPln(cfg.Path)
	fPln(cfg.LogFile)
	fPln(cfg.Access)
	fPln(cfg.Server)
	fPln(cfg.Route)
	fPln(cfg.ServiceName)
}
