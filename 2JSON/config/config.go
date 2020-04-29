package cfg

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/burntsushi/toml"
)

var (
	fPf        = fmt.Printf
	fPln       = fmt.Println
	sHasSuffix = strings.HasSuffix
)

// !!! toml file name must be identical to config struct name !!!

// Config :
type Config struct {
	Path           string
	LogFile        string
	AttrPrefix     string
	ContentPrefix  string
	DefaultSIFVer  string
	SIFCfgDir4LIST string
	SIFCfgDir4NUM  string
	SIFCfgDir4BOOL string
}

var (
	// toml file name must be identical to Config struct name
	lsCfg = []interface{}{
		&Config{},
	}
)

// ------------------------------------------------- //

// NewCfg :
func NewCfg(configs ...string) interface{} {
	for _, f := range configs {
		if _, e := os.Stat(f); e == nil {
			if abs, e := filepath.Abs(f); e == nil {
				return set(f, abs)
			}
		}
	}
	return nil
}

func set(f, abs string) interface{} {
	for _, cfg := range lsCfg {
		name := reflect.TypeOf(cfg).Elem().Name()
		if sHasSuffix(f, "/"+name+".toml") {
			if _, e := toml.DecodeFile(f, cfg); e == nil {
				reflect.ValueOf(cfg).Elem().FieldByName("Path").SetString(abs)
				save(f, cfg)
				// modify for runtime
				return cfg
			}
		}
	}
	return nil
}

func save(path string, cfg interface{}) {
	if f, e := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, os.ModePerm); e == nil {
		defer f.Close()
		toml.NewEncoder(f).Encode(cfg)
	}
}
