package main

import (
	"io/ioutil"

	cmn "github.com/cdutwhu/json-util/common"
)

func initMapOfObjAttrs(xpathGrp []string, sep string) {
	for _, xpath := range xpathGrp {
		ss := sSplit(xpath, sep)
		attr, attrType, objType := ss[0], ss[1], ss[len(ss)-1]
		mObjAttrs[objType] = append(mObjAttrs[objType], attr)
		mObjIdxOfAttr[objType] = 0
		mOAType[attr] = attrType
	}
}

func rewindAttrIter(objType string) {
	mObjIdxOfAttr[objType] = 0
}

// NextAttr :
func NextAttr(obj string) (value string, end bool) {
	if objType, ok := mOAType[obj]; ok {
		obj = objType
	}
	idx := mObjIdxOfAttr[obj]
	if idx == len(mObjAttrs[obj]) {
		return "", true
	}
	defer func() {
		mObjIdxOfAttr[obj]++
	}()
	return mObjAttrs[obj][idx], false
}

// PrintXML : append print to a string
func PrintXML(paper, line, contentHolder string, iLine int, tag string) (string, bool) {
	if iLine <= mOAPrtLine[tag] {
		return paper, false
	}
	mOAPrtLine[tag] = iLine

	if contentHolder != "" {
		return paper + line + "\n" + contentHolder + "\n", true
	}
	return paper + line + "\n", true
}

// SortAttr : xml is 4 space formatted
// obj [level] = attribute [level-1]
// NextAttr is available
func SortAttr(xml, obj string, level, times int) (paper string) {
	defer func() {
		fPln(paper)
	}()

	lines := sSplit(xml, "\n")

	const INDENT = "    " // 4 space
	indentObj, indentAttr := "", ""
	for i := 0; i < level; i++ {
		if i > 0 {
			indentObj += INDENT
		}
		indentAttr += INDENT
	}

	S1 := fSf("%s<%s ", indentObj, obj)
	S2 := fSf("%s<%s>", indentObj, obj)
	S3 := fSf("%s</%s>", indentObj, obj)

	for t := 0; t < times; t++ {

		for i, l := range lines {
			if sHasPrefix(l, S1) || sHasPrefix(l, S2) {
				if tempPaper, prt := PrintXML(paper, l, "", i, obj); !prt {
					continue
				} else {
					paper = tempPaper
					break
				}
			}
		}

		attr, end := NextAttr(obj)
	NEXTATTR:
		for ; !end; attr, end = NextAttr(obj) {
			// fPln(attr)

			S1 := fSf("%s<%s ", indentAttr, attr)
			S2 := fSf("%s<%s>", indentAttr, attr)
			S3 := fSf("%s</%s>", indentAttr, attr)
			E := fSf("</%s>", attr)

			for i, l := range lines {
				switch {
				case (sHasPrefix(l, S1) || sHasPrefix(l, S2)) && sHasSuffix(l, E): // one line
					if tempPaper, prt := PrintXML(paper, l, "", i, attr); !prt {
						continue
					} else {
						paper = tempPaper
						continue NEXTATTR
					}
				case sHasPrefix(l, S1) || sHasPrefix(l, S2): // sub-object START
					if tempPaper, prt := PrintXML(paper, l, "****", i, attr); !prt {
						continue
					} else {
						paper = tempPaper
						continue NEXTATTR
					}
				case sHasPrefix(l, S3): // sub-object END
					if tempPaper, prt := PrintXML(paper, l, "", i, "/"+attr); !prt {
						continue
					} else {
						paper = tempPaper
						continue NEXTATTR
					}
				}
			}
		}

		for i, l := range lines {
			if sHasPrefix(l, S3) {
				if tempPaper, prt := PrintXML(paper, l, "", i, "/"+obj); !prt {
					continue
				} else {
					paper = tempPaper
					break
				}
			}
		}
	}

	return
}

func main() {

	SIFSpecPath := "./out.txt"

	const (
		SEP       = "\t"
		XPATHTYPE = "XPATHTYPE:"
		OBJECT    = "OBJECT: "
	)

	var (
		xpathGrp []string
	)

	bytes, err := ioutil.ReadFile(SIFSpecPath)
	cmn.FailOnErr("%v", err)
	content := string(bytes)

	for _, line := range sSplit(content, "\n") {
		switch {
		case sHasPrefix(line, OBJECT):
			objGrp = append(objGrp, line[len(OBJECT):])
		case sHasPrefix(line, XPATHTYPE):
			l := sTrim(line[len(XPATHTYPE):], " \t")
			xpathGrp = append(xpathGrp, l)
		}
	}

	// for _, obj := range objGrp {
	// 	fPln(obj, mObjAttrs[obj])
	// }

	initMapOfObjAttrs(xpathGrp, SEP)

	// value, end := NextAttr("SoftwareRequirement")
	// for ; !end; value, end = NextAttr("SoftwareRequirement") {
	// 	fPln(value)
	// }

	bytes, err = ioutil.ReadFile("../data/Activity1.xml")
	cmn.FailOnErr("%v", err)
	SortAttr(string(bytes), "SoftwareRequirement", 3, 3)
}
