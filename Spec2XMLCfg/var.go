package main

import (
	"fmt"
	"strings"
)

var (
	fPln = fmt.Println
	fSf  = fmt.Sprintf
	fEf  = fmt.Errorf

	sHasPrefix = strings.HasPrefix
	sHasSuffix = strings.HasSuffix
	sSplit     = strings.Split
	sReplace   = strings.Replace
	sCount     = strings.Count
	sTrim      = strings.Trim
	sIndex     = strings.Index
	sContains  = strings.Contains
)

var (
	mObjAttrs     = make(map[string][]string) // key: obj-type
	mObjIdxOfAttr = make(map[string]int)      // key: obj-type
	mOAType       = make(map[string]string)   // key: obj
	mOAPrtLn      = make(map[string]int)      // key: obj
	objGrp        []string
)
