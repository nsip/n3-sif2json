package cfg

import "github.com/cdutwhu/n3-util/n3cfg"

// Config : AUTO Created From /home/qmiao/Desktop/temp/n3-sif2json/Config/config.toml
type Config struct {
	Service interface{}
	Log string
	Version interface{}
	SIF struct {
		DefaultVer string
	}
	Loggly struct {
		Token string
	}
	WebService struct {
		Port int
	}
	Route struct {
		ToJSON string
		ToSIF string
		Help string
	}
	NATS struct {
		Timeout int
		URL string
		Subject string
	}
	Server struct {
		Protocol string
		IP interface{}
		Port interface{}
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
