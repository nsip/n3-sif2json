package cvt2json

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/iter"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/misc"
	"github.com/cdutwhu/gotil/str"
	"github.com/cdutwhu/n3-util/jkv"
	"github.com/cdutwhu/n3-util/n3json"
	"github.com/cdutwhu/n3-util/n3xml"
)

var (
	fPln        = fmt.Println
	fPf         = fmt.Printf
	fSp         = fmt.Sprint
	fSf         = fmt.Sprintf
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sTrim       = strings.Trim
	sCount      = strings.Count
	sReplaceAll = strings.ReplaceAll
	sSplit      = strings.Split
	sNewReader  = strings.NewReader
	sJoin       = strings.Join
	sSplitRev   = func(s, sep string) []string {
		a := sSplit(s, sep)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		return a
	}

	enableLog2F     = fn.EnableLog2F
	failOnErr       = fn.FailOnErr
	failOnErrWhen   = fn.FailOnErrWhen
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
