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
	fPf             = fmt.Printf
	fPln            = fmt.Println
	fSp             = fmt.Sprint
	fSf             = fmt.Sprintf
	sHasPrefix      = strings.HasPrefix
	sHasSuffix      = strings.HasSuffix
	sReplaceAll     = strings.ReplaceAll
	failOnErr       = fn.FailOnErr
	enableLog2F     = fn.EnableLog2F
	failOnErrWhen   = fn.FailOnErrWhen
	localIP         = net.LocalIP
	sTrim           = strings.Trim
	sCount          = strings.Count
	sSplit          = strings.Split
	sNewReader      = strings.NewReader
	sJoin           = strings.Join
	splitRev        = str.SplitRev
	replByPosGrp    = str.ReplByPosGrp
	rmTailFromLastN = str.RmTailFromLastN
	rmTailFromLast  = str.RmTailFromLast
	rmHeadToLast    = str.RmHeadToLast
	iter2Slc        = iter.Iter2Slc
	mustWriteFile   = io.MustWriteFile
	exist           = judge.Exist
	Go              = misc.Go
	xmlRoot         = n3xml.XMLRoot
	jsonRoot        = n3json.JSONRoot
	fmtJSON         = n3json.Fmt
	newJKV          = jkv.NewJKV
)
