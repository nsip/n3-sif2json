package cvt2json

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/n3-util/common"
	jkv "github.com/cdutwhu/n3-util/jkv"
	"github.com/cdutwhu/n3-util/n3json"
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

	failOnErr       = cmn.FailOnErr
	failOnErrWhen   = cmn.FailOnErrWhen
	replByPosGrp    = cmn.ReplByPosGrp
	xmlRoot         = cmn.XMLRoot
	jsonRoot        = cmn.JSONRoot
	rmTailFromLastN = cmn.RmTailFromLastN
	rmTailFromLast  = cmn.RmTailFromLast
	rmHeadToLast    = cmn.RmHeadToLast
	iter2Slc        = cmn.Iter2Slc
	mustWriteFile   = cmn.MustWriteFile
	setLog          = cmn.SetLog
	resetLog        = cmn.ResetLog
	xin             = cmn.XIn
	fmtJSON         = n3json.Fmt
	newJKV          = jkv.NewJKV
)
