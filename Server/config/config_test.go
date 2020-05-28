package config

import "testing"

func TestLoad(t *testing.T) {
	cfg := newCfg("./config.toml")
	// fPf("%+v\n", *cfg)
	fPln(cfg.Path)
	fPln(cfg.LogFile)
	fPln(cfg.Cfg2JSON)
	fPln(cfg.Cfg2SIF)
	fPln(cfg.WebService)
	fPln(cfg.Route)
	fPln(cfg.NATS)
	fPln(cfg.File)
	fPln(cfg.ServiceName)
}

func TestInit(t *testing.T) {
	InitEnvVarFromTOML("Cfg", "./config.toml")
	cfg := env2Struct("Cfg", &Config{}).(*Config)
	fPln(cfg.Path)
	fPln(cfg.LogFile)
	fPln(cfg.Cfg2JSON)
	fPln(cfg.Cfg2SIF)
	fPln(cfg.WebService)
	fPln(cfg.Route)
	fPln(cfg.NATS)
	fPln(cfg.File)
	fPln(cfg.ServiceName)
}
