package config

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/json-util/common"
)

var (
	fPf         = fmt.Printf
	fPln        = fmt.Println
	sHasSuffix  = strings.HasSuffix
	sReplaceAll = strings.ReplaceAll

	failOnErr = cmn.FailOnErr
	localIP   = cmn.LocalIP
	cfgRepl   = cmn.CfgRepl
)
