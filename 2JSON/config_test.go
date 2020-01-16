package cvt2json

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if cfg := NewCfg("./config/list2json.toml"); cfg != nil {
		cfg := cfg.(*list2json)
		fPf("%+v\n", *cfg)
	}
	fPln()
	if cfg := NewCfg("./config/sif2json.toml"); cfg != nil {
		cfg := cfg.(*sif2json)
		fPf("%+v\n", *cfg)
	}
}
