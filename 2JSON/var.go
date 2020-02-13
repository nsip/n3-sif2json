package cvt2json

import (
	"fmt"
	"strings"
)

var (
	fPln        = fmt.Println
	fPf         = fmt.Printf
	fSp         = fmt.Sprint
	fSf         = fmt.Sprintf
	fEf         = fmt.Errorf
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

// var (
// cont = `"#content"`
// reContDigVal = regexp.MustCompile(fSf(`%s: "[0-9.]+"`, cont))
// contDigVal   = func(str string) (string, string) {
// 	if sHasPrefix(str, cont) {
// 		dig := str[len(cont)+3 : len(str)-1]
// 		return dig, fSf(`%s: %s`, cont, dig)
// 	}
// 	return "", ""
// }
// contValRepl = func(lsStr []string) (m map[string]string) {
// 	m = make(map[string]string)
// 	for _, str := range lsStr {
// 		_, modstr := contDigVal(str)
// 		m[str] = modstr
// 	}
// 	return m
// }
// )
