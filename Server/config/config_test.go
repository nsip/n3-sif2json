package cfg

import "testing"

func TestLoad(t *testing.T) {
	cfg := NewCfg("./config.toml")
	fPln(cfg.Path)
	fPln(cfg.PathAbs)
	fPln(cfg.ErrLog)
	fPln(cfg.WebService)
	fPln(cfg.Route)
}
