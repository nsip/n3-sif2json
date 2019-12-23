package cvt2json

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if cfg := NewCfg("./config/General.toml"); cfg != nil {
		cfg := cfg.(*General)
		fPln(*cfg)
	}
	if cfg := NewCfg("./config/List.toml"); cfg != nil {
		cfg := cfg.(*List)
		fPln(cfg.PurchaseOrder)
	}
	if cfg := NewCfg("./config/Prefix.toml"); cfg != nil {
		cfg := cfg.(*Prefix)
		fPln(*cfg)
	}
}
