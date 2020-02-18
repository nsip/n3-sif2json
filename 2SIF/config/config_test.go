package config

import "testing"

func TestConfig(t *testing.T) {
	if cfg := NewCfg("./JSON2SIF.toml"); cfg != nil {
		cfg := cfg.(*JSON2SIF)
		fPf("%+v\n", *cfg)
	}
}
