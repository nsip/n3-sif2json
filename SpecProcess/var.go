package main

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/json-util/common"
)

var (
	fPln = fmt.Println
	fPf  = fmt.Printf
	fSf  = fmt.Sprintf

	sHasPrefix = strings.HasPrefix
	sSplit     = strings.Split
	sReplace   = strings.Replace
	sCount     = strings.Count
	sTrim      = strings.Trim

	mapKeys        = cmn.MapKeys
	rmHeadToFirst  = cmn.RmHeadToFirst
	rmHeadToLast   = cmn.RmHeadToLast
	rmTailFromLast = cmn.RmTailFromLast
	failOnErr      = cmn.FailOnErr
	failOnErrWhen  = cmn.FailOnErrWhen
)
