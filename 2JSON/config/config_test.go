package cfg

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if cfg := NewCfg("./Config.toml"); cfg != nil {
		cfg := cfg.(*Config)
		fPf("%+v\n", *cfg)
	}
}
