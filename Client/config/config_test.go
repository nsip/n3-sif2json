package cfg

import "testing"

func TestLoad(t *testing.T) {
	cfg := NewCfg("./config.toml")
	fPln(cfg.Path)
	fPln(cfg.ELog)
	fPln(cfg.Access)
	fPln(cfg.Server)
	fPln(cfg.Route)
}
