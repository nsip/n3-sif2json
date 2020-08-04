package main

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/gotil/str"
)

var (
	fPln           = fmt.Println
	fPf            = fmt.Printf
	fSf            = fmt.Sprintf
	sHasPrefix     = strings.HasPrefix
	sSplit         = strings.Split
	sReplace       = strings.Replace
	sCount         = strings.Count
	sTrim          = strings.Trim
	mapKeys        = rflx.MapKeys
	rmHeadToFirst  = str.RmHeadToFirst
	rmHeadToLast   = str.RmHeadToLast
	rmTailFromLast = str.RmTailFromLast
	failOnErr      = fn.FailOnErr
	failOnErrWhen  = fn.FailOnErrWhen
)
