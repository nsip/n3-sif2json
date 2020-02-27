package cfg

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if cfg := NewCfg("./SIF2JSON.toml"); cfg != nil {
		cfg := cfg.(*SIF2JSON)
		fPf("%+v\n", *cfg)
	}
}
