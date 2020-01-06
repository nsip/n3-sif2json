package cvt2json

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if cfg := NewCfg("./config/General.toml"); cfg != nil {
		cfg := cfg.(*General)
		fPln(*cfg)
	}
	if cfg := NewCfg("./config/Path2JSON.toml"); cfg != nil {
		cfg := cfg.(*Path2JSON)
		fPln(*cfg)
	}
	if cfg := NewCfg("./config/XML2JSON.toml"); cfg != nil {
		cfg := cfg.(*XML2JSON)
		fPln(*cfg)
	}
}
