package main

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/n3-util/common"
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

	failOnErr     = cmn.FailOnErr
	failOnErrWhen = cmn.FailOnErrWhen
	fmtJSON       = n3json.Fmt
)

var (
	lsObjects        = []string{}
	mObjPaths        = map[string][]string{}
	mObjMaxLenOfPath = map[string]int{}

	clearBuf = func() {
		lsObjects = []string{}
		mObjPaths = map[string][]string{}
		mObjMaxLenOfPath = map[string]int{}
	}
)
