package cfg

import "testing"

func TestLoad(t *testing.T) {
	cfg := NewCfg("./config.toml")
	fPln(cfg.Path)
	fPln(cfg.ELog)
	fPln(cfg.Cfg2JSON)
	fPln(cfg.Cfg2SIF)
	fPln(cfg.WebService)
	fPln(cfg.Route)
}
