package webapi

import (
	"fmt"
	"strings"
	"sync"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/n3-util/n3log"
	"github.com/cdutwhu/n3-util/rest"
)

var (
	fPf         = fmt.Printf
	fPln        = fmt.Println
	fSf         = fmt.Sprintf
	sReplaceAll = strings.ReplaceAll

	localIP       = net.LocalIP
	isXML         = judge.IsXML
	isJSON        = judge.IsJSON
	enableLog2F   = fn.EnableLog2F
	logger        = fn.Logger
	warnOnErr     = fn.WarnOnErr
	failOnErr     = fn.FailOnErr
	mustWriteFile = io.MustWriteFile
	struct2Map    = rflx.Struct2Map
	env2Struct    = rflx.Env2Struct
	url1Value     = rest.URL1Value
	loggly        = n3log.Loggly
	logBind       = n3log.Bind
)

func initMutex(route interface{}) map[string]*sync.Mutex {
	mMtx := make(map[string]*sync.Mutex)
	for _, v := range struct2Map(route) {
		mMtx[v.(string)] = &sync.Mutex{}
	}
	return mMtx
}

type result struct {
	Data  string `json:"data"`
	Info  string `json:"info"`
	Error string `json:"error"`
}
