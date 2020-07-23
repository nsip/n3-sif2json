package main

import (
	"fmt"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
)

var (
	fPln          = fmt.Println
	failOnErr     = fn.FailOnErr
	failOnErrWhen = fn.FailOnErrWhen
	enableLog2F   = fn.EnableLog2F
	logWhen       = fn.LoggerWhen
	logger        = fn.Logger
	localIP       = net.LocalIP
	env2Struct    = rflx.Env2Struct
)
