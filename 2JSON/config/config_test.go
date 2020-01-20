package cfg

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if cfg := NewCfg("./List2JSON.toml"); cfg != nil {
		cfg := cfg.(*List2JSON)
		fPf("%+v\n", *cfg)
	}
	fPln()
	if cfg := NewCfg("./SIF2JSON.toml"); cfg != nil {
		cfg := cfg.(*SIF2JSON)
		fPf("%+v\n", *cfg)
	}
}
