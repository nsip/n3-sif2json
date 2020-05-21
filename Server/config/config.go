package config

import (
	"os"
	"path/filepath"
	"reflect"

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
		// modify BUT not save
		ver := fSf("%s", cfg.WebService.Version)
		return cfg.modCfg(map[string]string{"#v": ver}) // *** replace version *** //
	}
	return nil
}

func (cfg *Config) save() {
	if f, e := os.OpenFile(cfg.Path, os.O_WRONLY|os.O_TRUNC, os.ModePerm); e == nil {
		defer f.Close()
		toml.NewEncoder(f).Encode(cfg)
	}
}

func (cfg *Config) modCfg(mRepl map[string]string) *Config {
	if mRepl == nil || len(mRepl) == 0 {
		return cfg
	}
	nField := reflect.ValueOf(cfg.Route).NumField()
	for i := 0; i < nField; i++ {
		for key, value := range mRepl {
			replaced := sReplaceAll(reflect.ValueOf(cfg.Route).Field(i).Interface().(string), key, value)
			reflect.ValueOf(&cfg.Route).Elem().Field(i).SetString(replaced)
		}
	}
	return cfg
}
