package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/burntsushi/toml"
)

// Config is toml
type Config struct {
	Path        string
	LogFile     string
	Cfg2JSON    string
	Cfg2SIF     string
	ServiceName string
	WebService  struct {
		Port    int
		Version string
	}
	Route struct {
		HELP     string
		SIF2JSON string
		JSON2SIF string
	}
	NATS struct {
		URL     string
		Subject string
		Timeout int
	}
	File struct {
		ClientLinux64 string
		ClientMac     string
		ClientWin64   string
		ClientConfig  string
	}
}

// newCfg :
func newCfg(configs ...string) *Config {
	for _, f := range configs {
		if _, e := os.Stat(f); e == nil {
			return (&Config{Path: f}).set()
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

		// save
		cfg.save()

		ICfg, e := cfgRepl(cfg, map[string]interface{}{
			"[DATE]": time.Now().Format("2006-01-02"),
			"[v]":    cfg.WebService.Version,
		})
		failOnErr("%v", e)
		return ICfg.(*Config)
	}
	return nil
}

func (cfg *Config) save() {
	if f, e := os.OpenFile(cfg.Path, os.O_WRONLY|os.O_TRUNC, os.ModePerm); e == nil {
		defer f.Close()
		toml.NewEncoder(f).Encode(cfg)
	}
}

// InitEnvVarFromTOML : initialize the global variables
func InitEnvVarFromTOML(key string, configs ...string) bool {
	configs = append(configs, "./config.toml")
	Cfg := newCfg(configs...)
	if Cfg == nil {
		return false
	}
	struct2Env(key, Cfg)
	return true
}
