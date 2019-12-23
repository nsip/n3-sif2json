package cvt2json

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	fPln        = fmt.Println
	fSp         = fmt.Sprint
	fSf         = fmt.Sprintf
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sCount      = strings.Count
	sReplaceAll = strings.ReplaceAll
	sSplit      = strings.Split
	sNewReader  = strings.NewReader
	sSplitRev   = func(s, sep string) []string {
		a := sSplit(s, sep)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		return a
	}
)

var (
	lsObjects       = []string{}
	mObjLAttrs      = map[string][]string{}
	mObjMLenOfLAttr = map[string]int{}
)

var (
	cont         = `"#content"`
	reContDigVal = regexp.MustCompile(fSf(`%s: "[0-9.]+"`, cont))
	contDigVal   = func(str string) (string, string) {
		if sHasPrefix(str, cont) {
			dig := str[len(cont)+3 : len(str)-1]
			return dig, fSf(`%s: %s`, cont, dig)
		}
		return "", ""
	}
	contValRepl = func(lsStr []string) (m map[string]string) {
		m = make(map[string]string)
		for _, str := range lsStr {
			_, modstr := contDigVal(str)
			m[str] = modstr
		}
		return m
	}
)
