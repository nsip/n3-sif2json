package client

import "github.com/cdutwhu/n3-util/n3cfg"

// Config : AUTO Created From /home/qmiao/Desktop/temp/n3-sif2json/Server/client/config.toml
type Config struct {
	Service string
	Route struct {
		ToJSON string
		ToSIF string
		Help string
	}
	Server struct {
		Port int
		Protocol string
		IP string
	}
	Access struct {
		Timeout int
	}
}

// NewCfg :
func NewCfg(cfgStruName string, mReplExpr map[string]string, cfgPaths ...string) interface{} {
	var cfg interface{}
	switch cfgStruName {
	case "Config":
		cfg = &Config{}
	default:
		return nil
	}
	return n3cfg.InitEnvVar(cfg, mReplExpr, cfgStruName, cfgPaths...)
}
