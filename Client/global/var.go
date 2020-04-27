package global

import (
	cfg "github.com/nsip/n3-sif2json/Client/config"
)

var (
	// Cfg : global variable
	Cfg *cfg.Config
)

// Init : initialize the global variables
func Init() bool {
	Cfg = cfg.NewCfg("./config.toml", "./config/config.toml", "../config/config.toml", "../config.toml", "../../config.toml")
	return Cfg != nil
}
