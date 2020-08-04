package preprocess

import (
	"flag"
	"testing"

	"github.com/cdutwhu/n3-util/n3cfg"
)

// ****************************** //
// Auto generate config go struct file for [2JSON] [2SIF] [Server] [goclient].
func TestGenCfgStruct(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"2JSON", "2SIF", "Server"} // for this page unit test
	}
	for _, arg := range args {
		switch arg {
		case "2JSON":
			n3cfg.GenStruct("../2JSON/config.toml", "Config", "cvt2json", "../2JSON/config_auto.go")
		case "2SIF":
			n3cfg.GenStruct("../2SIF/config.toml", "Config", "cvt2sif", "../2SIF/config_auto.go")
		case "Server":
			n3cfg.GenStruct("../Server/config.toml", "Config", "main", "../Server/config_auto.go")
		case "goclient":
			n3cfg.GenStruct("../Server/goclient/config.toml", "Config", "goclient", "../Server/goclient/config_auto.go")
		}
	}
}
