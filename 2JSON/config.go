package cvt2json

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/burntsushi/toml"
)

// General : toml
type General struct {
	Path    string
	PathAbs string
	ErrLog  string
}

// Prefix : toml
type Prefix struct {
	Path          string
	PathAbs       string
	AttrPrefix    string
	ContentPrefix string
}

// List : toml
type List struct {
	Path          string
	PathAbs       string
	JQDir         string
	Sep           string
	PurchaseOrder struct {
		L1 string
		L2 string
		L3 string
	}
	Test struct {
		L1 string
		L2 string
		L3 string
	}
	Test1 struct {
		L1 string
	}
}

var (
	// toml file name must be identical to config struct definition name
	lsCfg = []interface{}{
		&General{},
		&Prefix{},
		&List{},
	}
)

// NewCfg :
func NewCfg(cfgPaths ...string) interface{} {
	for _, f := range cfgPaths {
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
				reflect.ValueOf(cfg).Elem().FieldByName("Path").SetString(f)
				reflect.ValueOf(cfg).Elem().FieldByName("PathAbs").SetString(abs)
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
