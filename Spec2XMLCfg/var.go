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
	mIPathSubXML  = make(map[string]string) // key: path@index
	mIPathSubMark = make(map[string]string) // key: path@index

	rewindAttrIter = func(objType string) {
		mObjIdxOfAttr[objType] = 0
	}

	resetPrt = func() {
		mOAPrtLn = make(map[string]int)
	}

	mkIndent = func(n int) (indent string) {
		const INDENT = "    " // 4 space
		for i := 0; i < n; i++ {
			indent += INDENT
		}
		return
	}
)
