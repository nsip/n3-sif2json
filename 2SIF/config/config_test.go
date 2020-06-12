package config

import "testing"

func TestConfig(t *testing.T) {
	cfg := NewCfg("./config.toml")
	fPln(cfg)
}
