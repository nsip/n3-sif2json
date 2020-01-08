package cvt2json

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if cfg := NewCfg("./config/General.toml"); cfg != nil {
		cfg := cfg.(*General)
		fPln(*cfg)
	}
	fPln()
	if cfg := NewCfg("./config/Cfg2JSON.toml"); cfg != nil {
		cfg := cfg.(*Cfg2JSON)
		fPf("%+v\n", *cfg)
	}
	fPln()
	if cfg := NewCfg("./config/XML2JSON.toml"); cfg != nil {
		cfg := cfg.(*XML2JSON)
		fPln(*cfg)
	}
}
