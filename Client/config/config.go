package cfg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/burntsushi/toml"
)

var (
	fPln = fmt.Println
)

// Config is toml
type Config struct {
	Path    string
	LogFile string
	Server  struct {
		Protocol string
		IP       string
		Port     int
	}
	Access struct {
		Timeout int
	}
	Route struct {
		ROOT     string
		SIF2JSON string
		JSON2SIF string
	}
}

// NewCfg :
func NewCfg(configs ...string) *Config {
	for _, f := range configs {
		if _, e := os.Stat(f); e == nil {
			cfg := &Config{Path: f}
			return cfg.set()
		}
	}
	return nil
}

// set is
func (cfg *Config) set() *Config {
	f := cfg.Path /* make a copy of original for restoring */
	if _, e := toml.DecodeFile(f, cfg); e == nil {
		// modify some to save
		cfg.Path = f
		if abs, e := filepath.Abs(f); e == nil {
			cfg.Path = abs
		}
		if logfile, e := filepath.Abs(cfg.LogFile); e == nil {
			cfg.LogFile = logfile
		}
		// save
		cfg.save()
		return cfg
	}
	return nil
}

func (cfg *Config) save() {
	if f, e := os.OpenFile(cfg.Path, os.O_WRONLY|os.O_TRUNC, os.ModePerm); e == nil {
		defer f.Close()
		toml.NewEncoder(f).Encode(cfg)
	}
}
