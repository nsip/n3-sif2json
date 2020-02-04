package cvt2xml

import (
	"encoding/json"
	"io/ioutil"

	cmn "github.com/cdutwhu/json-util/common"
	"github.com/clbanning/mxj"
)

// JSON2XML1 :
func JSON2XML1(jsonPath, xmlPath string) string {
	jsonBytes, err := ioutil.ReadFile(jsonPath)
	cmn.FailOnErr("%v", err)
	cmn.FailOnErrWhen(!cmn.IsJSON(string(jsonBytes)), "", fEf("Input JSON Path is not a valid JSON File"))

	var f interface{}
	json.Unmarshal(jsonBytes, &f)
	// fPln(f)

	b, err := mxj.AnyXmlIndent(f, "", "    ", "")
	cmn.FailOnErr("%v", err)

	xmlstr := string(b)
	xmlstr = sReplaceAll(xmlstr, "<>", "")
	xmlstr = sReplaceAll(xmlstr, "</>", "")
	xmlstr = re1.ReplaceAllString(xmlstr, "")
	xmlstr = re2.ReplaceAllString(xmlstr, "")
	xmlstr, _ = Indent(xmlstr, -4, false)
	xmlstr = sTrim(xmlstr, " \t\n")

	ioutil.WriteFile(xmlPath, []byte(xmlstr), 0666)
	return xmlstr
}

// JSON2XML2 :
func JSON2XML2(SIFSpecPath, xml1 string) string {
	const (
		SEP       = "\t"
		XPATHTYPE = "XPATHTYPE:"
	)

	bytes, err := ioutil.ReadFile(SIFSpecPath)
	cmn.FailOnErr("%v", err)
	spec := string(bytes)

	for _, line := range sSplit(spec, "\n") {
		switch {
		case sHasPrefix(line, XPATHTYPE):
			l := sTrim(line[len(XPATHTYPE):], " \t\r")
			xpathGrp = append(xpathGrp, l)
		}
	}

	InitMapOfObjAttrs(xpathGrp, SEP)

	// value, end := NextAttr("SoftwareRequirement")
	// for ; !end; value, end = NextAttr("SoftwareRequirement") {
	// 	fPln(value)
	// }

	// bytes, err = ioutil.ReadFile("../data/Activity1.xml")
	// cmn.FailOnErr("%v", err)
	// sifCont := string(bytes)
	sifCont := xml1

	// fPln(SortSimpleObject(sifCont, "Evaluation", 1))

	// Init "mIPathSubXML"
	root := cmn.XMLRoot(sifCont)
	ExtractOA(sifCont, root, "", 0)

	// xmlobj := SortSimpleObject(sifCont, root, 0)
	xmlobj := mIPathSubXML[root]

AGAIN:
	for k, subxml := range mIPathSubXML {
		mark := mIPathSubMark[k]
		xmlobj = sReplace(xmlobj, mark, subxml, 1)
	}
	if sContains(xmlobj, "...") {
		goto AGAIN
	}

	fPln(xmlobj)
	return xmlobj
}

// JSON2XML3 :
func JSON2XML3(xml2 string) string {
	return xml2
}

// ----------------------------------------- //

// InitMapOfObjAttrs : xpathGrp is from SIF Spec txt
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

// SortSimpleObject : xml is 4 space formatted, level is obj level
// obj [level] = attribute [level-1]
// NextAttr is available
func SortSimpleObject(xml, obj string, level int) (paper string) {
	defer func() {
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

// ExtractOA :
func ExtractOA(xml, obj, parent string, lvl int) string {
	S := mkIndent(lvl+1) + "<"
	E := S + "/"

	lvlOAs := []string{}
	xmlobj := sTrim(SortSimpleObject(xml, obj, lvl), "\n")
	for _, l := range sSplit(xmlobj, "\n") {
		sl := 0
		switch {
		case sHasPrefix(l, S) && !sHasPrefix(l, E) && sHasSuffix(l, "..."):
			sl = len(S)
		default:
			continue
		}
		oa := cmn.RmTailFromFirstAny(l[sl:], " ", ">")
		if len(lvlOAs) == 0 {
			lvlOAs = append(lvlOAs, oa)
			continue
		}
		lastOA := lvlOAs[len(lvlOAs)-1]
		if oa != lastOA {
			lvlOAs = append(lvlOAs, oa)
		}
	}

	ipath := parent + "~" + obj
	if parent == "" {
		ipath = obj
	}

	mIPathSubXML[ipath] = xmlobj
	xmlobjLn1 := sSplit(xmlobj, "\n")[0]

	preBlank := mkIndent(sCount(ipath, "~"))
	mIPathSubMark[ipath] = fSf("%s...\n%s</%s>", xmlobjLn1, preBlank, obj)

	for _, subobj := range lvlOAs {
		ExtractOA(xml, subobj, ipath, lvl+1)
	}

	return xmlobj
}
