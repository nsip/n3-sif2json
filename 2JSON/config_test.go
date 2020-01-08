package cvt2json

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if cfg := NewCfg("./config/cfg2json.toml"); cfg != nil {
		cfg := cfg.(*cfg2json)
		fPf("%+v\n", *cfg)
	}
	fPln()
	if cfg := NewCfg("./config/sif2json.toml"); cfg != nil {
		cfg := cfg.(*sif2json)
		fPf("%+v\n", *cfg)
	}
}
