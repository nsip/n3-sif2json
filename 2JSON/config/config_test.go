package config

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestConfig(t *testing.T) {
	cfg := NewCfg("./config.toml")
	spew.Dump(cfg)
}
