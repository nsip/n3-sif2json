package config

import (
	"os"
	"path/filepath"
	"reflect"
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

// NewCfg :
func NewCfg(configs ...string) *Config {
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
		if logfile, e := filepath.Abs(cfg.LogFile); e == nil {
			cfg.LogFile = logfile
		}

		// save
		cfg.save()

		// modify BUT not save
		return cfg.modCfg(map[string]interface{}{
			"[DATE]": time.Now().Format("2006-01-02"),
			"[v]":    cfg.WebService.Version,
		}) // *** replace version *** //
	}
	return nil
}

func (cfg *Config) save() {
	if f, e := os.OpenFile(cfg.Path, os.O_WRONLY|os.O_TRUNC, os.ModePerm); e == nil {
		defer f.Close()
		toml.NewEncoder(f).Encode(cfg)
	}
}

func (cfg *Config) modCfg(mRepl map[string]interface{}) *Config {
	if mRepl == nil || len(mRepl) == 0 {
		return cfg
	}

	cfgElem := reflect.ValueOf(cfg).Elem()
	for i, nField := 0, cfgElem.NumField(); i < nField; i++ {
		for key, value := range mRepl {
			field := cfgElem.Field(i)

			// string replace
			if oriVal, ok := field.Interface().(string); ok {
				if replaced := sReplaceAll(oriVal, key, value.(string)); replaced != oriVal {
					field.SetString(replaced)
				}
			}
			// TODO : SetInt ... if needed

			// go into struct, String replace
			if field.Kind() == reflect.Struct {
				for j, nFieldSub := 0, field.NumField(); j < nFieldSub; j++ {
					fieldSub := field.Field(j)

					// string replace
					if oriVal, ok := fieldSub.Interface().(string); ok {
						if replaced := sReplaceAll(oriVal, key, value.(string)); replaced != oriVal {
							fieldSub.SetString(replaced)
						}
					}
					// TODO : SetInt ... if needed
				}
			}
		}
	}
	return cfg
}
