package cvt2sif

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/endec"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/gotil/str"
	"github.com/cdutwhu/n3-util/n3xml"
)

var (
	fPln = fmt.Println
	fSp  = fmt.Sprint
	fSf  = fmt.Sprintf

	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
	sSplit      = strings.Split
	sReplace    = strings.Replace
	sCount      = strings.Count
	sTrim       = strings.Trim
	sTrimLeft   = strings.TrimLeft
	sIndex      = strings.Index
	sContains   = strings.Contains
	sReplaceAll = strings.ReplaceAll
	sSpl        = strings.Split
	sJoin       = strings.Join
	sNewReader  = strings.NewReader
	sSplitRev   = func(s, sep string) []string {
		a := sSpl(s, sep)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		return a
	}

	rmTailFromLast     = str.RmTailFromLast
	rmHeadToLast       = str.RmHeadToLast
	rmTailFromFirstAny = str.RmTailFromFirstAny
	replByPosGrp       = str.ReplByPosGrp
	failOnErrWhen      = fn.FailOnErrWhen
	failOnErr          = fn.FailOnErr
	failP1OnErrWhen    = fn.FailP1OnErrWhen
	failP1OnErr        = fn.FailP1OnErr
	enableLog2F        = fn.EnableLog2F
	isJSON             = judge.IsJSON
	md5Str             = endec.MD5Str
	mapsMerge          = rflx.MapMerge
	mustWriteFile      = io.MustWriteFile
	xmlRoot            = n3xml.XMLRoot
)

var nGoTo = 0

const maxGoTo = 100

var (
	re1 = regexp.MustCompile("\n[ ]*<#content>")
	re2 = regexp.MustCompile("</#content>\n[ ]*")

	TrvsGrpViaSpec []string                    // from SIF Spec
	mPathAttrs     = make(map[string][]string) // key: spec path, value: attribute-value
	mPathAttrIdx   = make(map[string]int)      // key: spec path, value: attribute-index

	mObjIdxStart  = make(map[string]int)    // key: obj-type@level, value: line-number
	mObjIdxEnd    = make(map[string]int)    // key: obj-type@level, value: line-number
	mOAPrtLn      = make(map[string]int)    // key: obj
	mIPathSubXML  = make(map[string]string) // key: path@index
	mIPathSubMark = make(map[string]string) // key: path@index
	mPathIdx      = make(map[string]int)    // key: path, for IPath

	RewindAttrIter = func() {
		for k := range mPathAttrIdx {
			mPathAttrIdx[k] = 0
		}
	}

	ResetPrt = func() {
		mOAPrtLn = make(map[string]int)
	}

	ResetAll = func() {
		mObjIdxStart = make(map[string]int)
		mObjIdxEnd = make(map[string]int)
		mOAPrtLn = make(map[string]int)
		mIPathSubXML = make(map[string]string)
		mIPathSubMark = make(map[string]string)
		mPathIdx = make(map[string]int)
	}

	mkIndent = func(n int) (indent string) {
		const INDENT = "    " // 4 space
		for i := 0; i < n; i++ {
			indent += INDENT
		}
		return
	}

	getReplMap = func(jsonPath string) (m map[string]string) {
		bytes, err := ioutil.ReadFile(jsonPath)
		fn.FailOnErr("%v", err)
		fn.FailOnErr("%v", json.Unmarshal(bytes, &m))
		return
	}
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
