package cfgreg

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"
	"testing"

	"github.com/cdutwhu/n3-util/n3cfg"
)

// Under n3-sif2json/Preprocess/CfgReg/
// echo password | sudo -S env "PATH=$PATH" go test -v -count=1 ./ -run TestRegCfg -args `whoami`
func TestRegCfg(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}
	user, _ := user.Current()
	osuser := user.Name
	if len(flag.Args()) > 0 {
		osuser = flag.Args()[0]
	}
	project := n3cfg.PrjName()
	// fmt.Println(osuser, project)

	mPkgConfig := map[string]string{
		"server":   "../../Server/config.toml",
		"goclient": "../../Server/goclient/config.toml",
		"cvt2json": "../../2JSON/config.toml",
		"cvt2sif":  "../../2SIF/config.toml",
	}

	for _, pkg := range flag.Args()[1:] {
		config := mPkgConfig[strings.ToLower(pkg)]
		if _, err := os.Stat(config); err == nil {
			ok, file := n3cfg.Register(osuser, config, project, pkg)
			fmt.Println(ok, file)
		}
	}
}
