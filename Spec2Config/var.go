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
	sSplit     = strings.Split
	sReplace   = strings.Replace
	sCount     = strings.Count
)
