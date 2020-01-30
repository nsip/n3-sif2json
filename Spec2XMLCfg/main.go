package main

import (
	"io/ioutil"

	cmn "github.com/cdutwhu/json-util/common"
)

func initMapOfObjAttrs(xpathGrp []string, sep string) {
	for _, xpath := range xpathGrp {
		ss := sSplit(xpath, sep)
		attr, attrType, objType := ss[0], ss[1], ss[4]
		mObjAttrs[objType] = append(mObjAttrs[objType], attr)
		mObjIdxOfAttr[objType] = 0
		mOAType[attr] = attrType
	}
}

// RewindAttrIter :
func RewindAttrIter(objType string) {
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

		RewindAttrIter(mOAType[obj])
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

	initMapOfObjAttrs(xpathGrp, SEP)

	// value, end := NextAttr("SoftwareRequirement")
	// for ; !end; value, end = NextAttr("SoftwareRequirement") {
	// 	fPln(value)
	// }

	bytes, err = ioutil.ReadFile("../data/Activity1.xml")
	cmn.FailOnErr("%v", err)
	// SortSimpleObject(string(bytes), "ActivityTime", 1)
	SortSimpleObject(string(bytes), "Activity", 0)
}
