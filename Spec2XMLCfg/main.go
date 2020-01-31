package main

import (
	"io/ioutil"

	cmn "github.com/cdutwhu/json-util/common"
)

// InitMapOfObjAttrs :
func InitMapOfObjAttrs(xpathGrp []string, sep string) {
	for _, xpath := range xpathGrp {
		ss := sSplit(xpath, sep)
		attr, attrType, objType := ss[0], ss[1], ss[4]
		mObjAttrs[objType] = append(mObjAttrs[objType], attr)
		mObjIdxOfAttr[objType] = 0
		mOAType[attr] = attrType
	}
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
	if _, ok := mOAPrtLn[tag]; !ok {
		mOAPrtLn[tag] = -1
	}

	if iLine <= mOAPrtLn[tag] {
		return paper, false
	}
	mOAPrtLn[tag] = iLine

	if contentHolder != "" {
		return paper + line + contentHolder + "\n", true
	}
	return paper + line + "\n", true
}

// SortSimpleObject : xml is 4 space formatted, level is obj's level
// obj [level] = attribute [level-1]
// NextAttr is available
func SortSimpleObject(xml, obj string, level int) (paper string) {
	defer func() {
		fPln(paper)
		resetPrt()
	}()

	const INDENT = "    " // 4 space
	indentObj, indentAttr := "", INDENT
	for i := 0; i < level; i++ {
		indentObj += INDENT
		indentAttr += INDENT
	}

	OS1 := fSf("%s<%s ", indentObj, obj)
	OS2 := fSf("%s<%s>", indentObj, obj)
	OS3 := fSf("%s</%s>", indentObj, obj)

	lines := sSplit(xml, "\n")

	// Find nObj
	nObj := sCount(xml, OS1)
	if n := sCount(xml, OS2); n > nObj {
		nObj = n
	}

	//	NEXTOBJ:
	for t := 0; t < nObj; t++ {

		rewindAttrIter(mOAType[obj])
		PS, PE := 0, 0

		for i, l := range lines {
			if sHasPrefix(l, OS1) || sHasPrefix(l, OS2) {
				if tempPaper, prt := PrintXML(paper, l, "", i, obj); !prt {
					continue
				} else {
					paper = tempPaper
					PS = i
					break
				}
			}
		}

		for i, l := range lines {
			if sHasPrefix(l, OS3) {
				if _, prt := PrintXML(paper, l, "", i, "//"+obj); !prt { // [//+obj] is probe to detect End Position
					continue
				} else {
					PE = i
					break
				}
			}
		}

		attr, end := NextAttr(obj)
	NEXTATTR:
		for ; !end; attr, end = NextAttr(obj) {
			// fPln(attr)

			AS1 := fSf("%s<%s ", indentAttr, attr)
			AS2 := fSf("%s<%s>", indentAttr, attr)
			AS3 := fSf("%s</%s>", indentAttr, attr)
			AE := fSf("</%s>", attr)

			// fPln(AS1, "|", AS2, "|", AS3, "| ------------------------------- ")

			for i, l := range lines {

				if i > PS && i < PE {

					// fPln(i, l)

					switch {
					case (sHasPrefix(l, AS1) || sHasPrefix(l, AS2)) && sHasSuffix(l, AE): // one line
						if tempPaper, prt := PrintXML(paper, l, "", i, attr); !prt {
							continue
						} else {
							paper = tempPaper
							continue NEXTATTR
						}
					case sHasPrefix(l, AS1) || sHasPrefix(l, AS2): // sub-object START
						if tempPaper, prt := PrintXML(paper, l, "...", i, attr); !prt {
							continue
						} else {
							paper = tempPaper
							continue
						}
					case sHasPrefix(l, AS3): // sub-object END
						if tempPaper, prt := PrintXML(paper, l, "", i, "/"+attr); !prt {
							continue
						} else {
							paper = tempPaper
							continue NEXTATTR
						}
					}
				}
			}
		}

		for i, l := range lines {
			if sHasPrefix(l, OS3) {
				if tempPaper, prt := PrintXML(paper, l, "", i, "/"+obj); !prt {
					continue
				} else {
					paper = tempPaper
					break
				}
			}
		}

	} // end of [nObj] loop

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
			l := sTrim(line[len(XPATHTYPE):], " \t\r")
			xpathGrp = append(xpathGrp, l)
		}
	}

	// for _, obj := range objGrp {
	// 	fPln(obj, mObjAttrs[obj])
	// }

	InitMapOfObjAttrs(xpathGrp, SEP)

	// value, end := NextAttr("SoftwareRequirement")
	// for ; !end; value, end = NextAttr("SoftwareRequirement") {
	// 	fPln(value)
	// }

	bytes, err = ioutil.ReadFile("../data/Activity1.xml")
	cmn.FailOnErr("%v", err)
	// SortSimpleObject(string(bytes), "SoftwareRequirement", 2)

	ScanOA(string(bytes))
}

// --------------------------------------- //

// ScanOA :
func ScanOA(xml string) {

	var (
		mLvlOAs = make(map[int][]string)
	)

	ss := sSplit(xml, "\n")
	for _, l := range ss {
		lvl, sl := 0, 0
		switch {
		case sHasPrefix(l, S0) && !sHasPrefix(l, E0):
			lvl, sl = 0, len(S0)
		case sHasPrefix(l, S1) && !sHasPrefix(l, E1):
			lvl, sl = 1, len(S1)
		case sHasPrefix(l, S2) && !sHasPrefix(l, E2):
			lvl, sl = 2, len(S2)
		case sHasPrefix(l, S3) && !sHasPrefix(l, E3):
			lvl, sl = 3, len(S3)
		case sHasPrefix(l, S4) && !sHasPrefix(l, E4):
			lvl, sl = 4, len(S4)
		case sHasPrefix(l, S5) && !sHasPrefix(l, E5):
			lvl, sl = 5, len(S5)
		case sHasPrefix(l, S6) && !sHasPrefix(l, E6):
			lvl, sl = 6, len(S6)
		case sHasPrefix(l, S7) && !sHasPrefix(l, E7):
			lvl, sl = 7, len(S7)
		case sHasPrefix(l, S8) && !sHasPrefix(l, E8):
			lvl, sl = 8, len(S8)
		case sHasPrefix(l, S9) && !sHasPrefix(l, E9):
			lvl, sl = 9, len(S9)
		case sHasPrefix(l, S10) && !sHasPrefix(l, E10):
			lvl, sl = 10, len(S10)
		case sHasPrefix(l, S11) && !sHasPrefix(l, E11):
			lvl, sl = 11, len(S11)
		default:
			continue
		}

		oa := cmn.RmTailFromFirstAny(l[sl:], " ", ">")
		mLvlOAs[lvl] = append(mLvlOAs[lvl], oa)
	}

	for i := 0; i < 11; i++ {
		fPln(i, mLvlOAs[i])
		for _, oa := range mLvlOAs[i] {
			SortSimpleObject(xml, oa, i)
		}
	}
}
