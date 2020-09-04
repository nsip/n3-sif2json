package cvt2json

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/iter"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/misc"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/str"
	"github.com/cdutwhu/n3-util/jkv"
	"github.com/cdutwhu/n3-util/n3json"
	"github.com/cdutwhu/n3-util/n3xml"
)

var (
	fPf              = fmt.Printf
	fPln             = fmt.Println
	fSp              = fmt.Sprint
	fSf              = fmt.Sprintf
	sHasPrefix       = strings.HasPrefix
	sHasSuffix       = strings.HasSuffix
	sReplaceAll      = strings.ReplaceAll
	sToLower         = strings.ToLower
	sTrim            = strings.Trim
	sCount           = strings.Count
	sSplit           = strings.Split
	sNewReader       = strings.NewReader
	sJoin            = strings.Join
	failOnErr        = fn.FailOnErr
	enableLog2F      = fn.EnableLog2F
	failOnErrWhen    = fn.FailOnErrWhen
	enableWarnDetail = fn.EnableWarnDetail
	warnOnErr        = fn.WarnOnErr
	warner           = fn.Warner
	localIP          = net.LocalIP
	splitRev         = str.SplitRev
	replByPosGrp     = str.ReplByPosGrp
	rmTailFromLastN  = str.RmTailFromLastN
	rmTailFromLast   = str.RmTailFromLast
	rmHeadToLast     = str.RmHeadToLast
	iter2Slc         = iter.Iter2Slc
	mustWriteFile    = io.MustWriteFile
	exist            = judge.Exist
	Go               = misc.Go
	trackTime        = misc.TrackTime
	xmlRoot          = n3xml.XMLRoot
	jsonRoot         = n3json.JSONRoot
	fmtJSON          = n3json.Fmt
	newJKV           = jkv.NewJKV
)

var (
	DftSIFVer = "3.4.7"
)
