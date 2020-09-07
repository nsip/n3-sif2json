package cfggen

import (
	"flag"
	"os"
	"testing"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/cdutwhu/n3-util/n3cfg/attrim"
	"github.com/cdutwhu/n3-util/n3err"
)

// ****************************** //
// go test -v -timeout 1s -count=1 $workpath/cfggen -run TestMkCltCfg -args "Path" "Service" "Route" "Server" "Access"
// working in build.sh
func TestMkCfg4Clt(t *testing.T) {
	cfg := n3cfg.ToEnvN3sif2jsonAll(
		map[string]string{
			"[v]":    "Version",
			"[s]":    "Service",
			"[port]": "WebService.Port",
		},
		"TestKey",
		"../../Config/config.toml",
	)
	fn.FailOnErrWhen(cfg == nil, "%v", n3err.CFG_INIT_ERR)

	temp := "./temp.toml"
	defer os.Remove(temp)
	n3cfg.Save(temp, cfg)

	// A Copy for Server Executable
	attrim.RmCfgAttrL1(temp, "../config_rel.toml")

	// A Copy for Server's goclient
	if !flag.Parsed() {
		flag.Parse()
	}
	attrim.SelCfgAttrL1(temp, "../client/config/config.toml", flag.Args()...)
}
