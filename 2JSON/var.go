package cvt2json

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	cmn "github.com/cdutwhu/json-util/common"
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
	sSplitRev   = func(s, sep string) []string {
		a := sSplit(s, sep)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		return a
	}
)

var (
	lsObjects         = []string{}
	mObjLAttrs        = map[string][]string{}
	mObjMaxLenOfLAttr = map[string]int{}
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

	xmlroot = func(xml string) (root string) {
		xml = sTrim(xml, " \t\n")
		start, end := 0, 0
		for i := len(xml) - 1; i >= 0; i-- {
			switch xml[i] {
			case '>':
				end = i
			case '/':
				start = i + 1
			}
			if start != 0 && end != 0 {
				break
			}
		}
		root = xml[start:end]

		// check, flag (?s) let . includes "NewLine"
		re1 := regexp.MustCompile(fSf(`(?s)^<%s .+</%s>$`, root, root))
		re2 := regexp.MustCompile(fSf(`(?s)^<%s>.+</%s>$`, root, root))
		cmn.FailOnCondition(!re1.MatchString(xml) && !re2.MatchString(xml), "%v", fEf("Invalid XML"))
		return
	}

	jsonroot = func(jsonstr string) string {
		x := make(map[string]interface{})
		cmn.FailOnErr("%v", json.Unmarshal([]byte(jsonstr), &x))
		for k := range x {
			return k
		}
		return ""
	}
)
