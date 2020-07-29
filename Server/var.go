package main

import (
	"fmt"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/n3-util/n3log"
)

var (
	fPln           = fmt.Println
	fSf            = fmt.Sprintf
	failOnErr      = fn.FailOnErr
	failOnErrWhen  = fn.FailOnErrWhen
	enableLog2F    = fn.EnableLog2F
	logWhen        = fn.LoggerWhen
	logger         = fn.Logger
	localIP        = net.LocalIP
	env2Struct     = rflx.Env2Struct
	lrInit         = n3log.LrInit
	loggly         = n3log.Loggly
	logBind        = n3log.Bind
	enableLoggly   = n3log.EnableLoggly
	setLogglyToken = n3log.SetLogglyToken
)
