package cvt2xml

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
	sSpl        = strings.Split
	sTrim       = strings.Trim
	sTrimLeft   = strings.TrimLeft
	sJoin       = strings.Join
	sNewReader  = strings.NewReader
	sSplitRev   = func(s, sep string) []string {
		a := sSpl(s, sep)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		return a
	}
)

var (
	re1 = regexp.MustCompile("\n[ ]*<#content>")
	re2 = regexp.MustCompile("</#content>\n[ ]*")
)

// Indent :
func Indent(str string, n int, ign1stLn bool) (string, bool) {
	if n == 0 {
		return str, false
	}
	S := 0
	if ign1stLn {
		S = 1
	}
	lines := sSpl(str, "\n")
	if n > 0 {
		space := ""
		for i := 0; i < n; i++ {
			space += " "
		}
		for i := S; i < len(lines); i++ {
			if sTrim(lines[i], " \n\t") != "" {
				lines[i] = space + lines[i]
			}
		}
	} else {
		for i := S; i < len(lines); i++ {
			if len(lines[i]) == 0 { //                                         ignore empty string line
				continue
			}
			if len(lines[i]) <= -n || sTrimLeft(lines[i][:-n], " ") != "" { // cannot be indented as <n>, give up indent
				return str, false
			}
			lines[i] = lines[i][-n:]
		}
	}
	return sJoin(lines, "\n"), true
}
