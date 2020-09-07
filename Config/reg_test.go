package config

import (
	"flag"
	"os"
	"os/user"
	"testing"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/n3-util/n3cfg"
)

// *** echo password | sudo -S env "PATH=$PATH" go test -v -count=1 ./ -run TestRegCfg
func TestRegCfg(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}
	user, _ := user.Current()
	osuser := user.Name
	if len(flag.Args()) > 0 {
		osuser = flag.Args()[0] // `whoami`
	}
	prj := n3cfg.PrjName()
	config := "./config.toml"
	_, err := os.Stat(config)
	fn.FailOnErr("%v @ [%s]", err, config)
	ok, file := n3cfg.Register(osuser, "./config.toml", prj, "all")
	fn.Logger("%v %v", ok, file)
}
