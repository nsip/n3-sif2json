package config

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/n3-util/n3cfg"
)

var (
	fPf         = fmt.Printf
	fPln        = fmt.Println
	sHasSuffix  = strings.HasSuffix
	sReplaceAll = strings.ReplaceAll

	failOnErr = fn.FailOnErr
	localIP   = net.LocalIP
	cfgRepl   = n3cfg.Modify
)
