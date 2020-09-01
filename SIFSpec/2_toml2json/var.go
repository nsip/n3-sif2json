package main

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/embres"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/str"
	"github.com/cdutwhu/n3-util/n3json"
)

var (
	fPln           = fmt.Println
	fPf            = fmt.Printf
	fSp            = fmt.Sprint
	fSf            = fmt.Sprintf
	sHasPrefix     = strings.HasPrefix
	sHasSuffix     = strings.HasSuffix
	sTrim          = strings.Trim
	sCount         = strings.Count
	sContains      = strings.Contains
	sReplaceAll    = strings.ReplaceAll
	sSplit         = strings.Split
	sNewReader     = strings.NewReader
	sJoin          = strings.Join
	splitRev       = str.SplitRev
	mustWriteFile  = io.MustWriteFile
	failOnErr      = fn.FailOnErr
	failOnErrWhen  = fn.FailOnErrWhen
	createDirBytes = embres.CreateDirBytes
	printFileBytes = embres.PrintFileBytes
	fmtJSON        = n3json.Fmt
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
