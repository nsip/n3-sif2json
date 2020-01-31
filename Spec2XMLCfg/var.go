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

	S0  = mkIndent(0) + "<"
	S1  = mkIndent(1) + "<"
	S2  = mkIndent(2) + "<"
	S3  = mkIndent(3) + "<"
	S4  = mkIndent(4) + "<"
	S5  = mkIndent(5) + "<"
	S6  = mkIndent(6) + "<"
	S7  = mkIndent(7) + "<"
	S8  = mkIndent(8) + "<"
	S9  = mkIndent(9) + "<"
	S10 = mkIndent(10) + "<"
	S11 = mkIndent(11) + "<"

	E0  = S0 + "/"
	E1  = S1 + "/"
	E2  = S2 + "/"
	E3  = S3 + "/"
	E4  = S4 + "/"
	E5  = S5 + "/"
	E6  = S6 + "/"
	E7  = S7 + "/"
	E8  = S8 + "/"
	E9  = S9 + "/"
	E10 = S10 + "/"
	E11 = S11 + "/"
)
