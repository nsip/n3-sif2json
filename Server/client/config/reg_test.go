package config

import (
	"flag"
	"os"
	"os/user"
	"testing"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/n3-util/n3cfg"
)

// *** echo sudo-password | sudo -S env "PATH=$PATH" go test -v -count=1 ./ -run TestRegCfg -args `whoami`
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
	ok, file := n3cfg.Register(osuser, "./config.toml", prj, "goclient")
	fn.Logger("%v %v", ok, file)
}
