package config

import (
	"flag"
	"os"
	"testing"

	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/cdutwhu/n3-util/n3err"
	"github.com/davecgh/go-spew/spew"
)

func TestLoad(t *testing.T) {
	cfg := newCfg("./config.toml")
	spew.Dump(cfg)
}

func TestInit(t *testing.T) {
	InitEnvVarFromTOML("Cfg", "./config.toml")
	ICfg := env2Struct("Cfg", &Config{})
	spew.Dump(ICfg.(*Config))
}

// ****************************** //
// Create a copy of config for Client. Excluding some attributes.
// Once building, move it to Client config Directory.
func TestGenCltCfg(t *testing.T) {
	cfg := newCfg("./config.toml")
	failOnErrWhen(cfg == nil, "%v", n3err.CFG_INIT_ERR)
	temp := "./temp.toml"
	cfg.SaveAs(temp)
	if !flag.Parsed() {
		flag.Parse()
	}
	n3cfg.SelFileAttrL1(temp, "../goclient/config.toml", flag.Args()...)
	os.Remove(temp)
}
