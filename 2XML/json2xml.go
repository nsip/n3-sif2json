package cvt2xml

import (
	"encoding/json"
	"io/ioutil"
	"regexp"

	cmn "github.com/cdutwhu/json-util/common"
	"github.com/clbanning/mxj"
)

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

// NextAttr : From Spec
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
func PrintXML(paper, line, mark string, iLine int, tag string) (string, bool) {
	if _, ok := mOAPrtLn[tag]; !ok {
		mOAPrtLn[tag] = -1
	}

	if iLine <= mOAPrtLn[tag] {
		return paper, false
	}
	mOAPrtLn[tag] = iLine

	if mark != "" {
		return paper + line + mark + "\n", true
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

	objIdx := fSf("%s@%d", obj, level)
	if _, ok := mObjIdxStart[objIdx]; !ok {
		mObjIdxStart[objIdx] = -1
	}
	if _, ok := mObjIdxEnd[objIdx]; !ok {
		mObjIdxEnd[objIdx] = -1
	}

	rewindAttrIter(mOAType[obj])
	PS, PE := -1, -1

	// ---------------------------------- //
	for i, l := range lines {
		if (sHasPrefix(l, OS1) || sHasPrefix(l, OS2)) && i > mObjIdxStart[objIdx] {
			if _, ok := PrintXML(paper, l, "", i, "*"+obj); ok { // [*+obj] is probe to detect Start line
				PS, mObjIdxStart[objIdx] = i, i
			}
		}
		if sHasPrefix(l, OS3) && i > mObjIdxEnd[objIdx] {
			if _, ok := PrintXML(paper, l, "", i, "*/"+obj); ok { // [*/+obj] is probe to detect End line
				PE, mObjIdxEnd[objIdx] = i, i
			}
		}
		if PS != -1 && PE != -1 { // if all found, break. PE may not be found when Single Line Attribute
			break
		}
	}
	// ---------------------------------- //

	paper, _ = PrintXML(paper, lines[PS], fSf("@%d#", PS), PS, obj)

	for attr, end := NextAttr(obj); !end; attr, end = NextAttr(obj) {
		AS1 := fSf("%s<%s ", indentAttr, attr)
		AS2 := fSf("%s<%s>", indentAttr, attr)
		AS3 := fSf("%s</%s>", indentAttr, attr)
		AE := fSf("</%s>", attr)

		for i, l := range lines {
			if i > PS && i < PE {
				switch {
				case (sHasPrefix(l, AS1) || sHasPrefix(l, AS2)) && sHasSuffix(l, AE): // one line
					if tempPaper, ok := PrintXML(paper, l, "", i, attr); ok {
						paper = tempPaper
					}
				case sHasPrefix(l, AS1) || sHasPrefix(l, AS2): // sub-object START
					if tempPaper, ok := PrintXML(paper, l, fSf("@%d#...", i), i, attr); ok {
						paper = tempPaper
					}
				case sHasPrefix(l, AS3): // sub-object END
					if tempPaper, ok := PrintXML(paper, l, "", i, "/"+attr); ok {
						paper = tempPaper
					}
				}
			}
		}
	}

	// Single Line Object has NO End Tag
	if PE != -1 {
		paper, _ = PrintXML(paper, lines[PE], "", PE, "/"+obj)
	}

	return
}

// ExtractOA :
func ExtractOA(xml, obj, parent string, lvl int) string {
	S := mkIndent(lvl+1) + "<"
	E := S + "/"

	lvlOAs := []string{} // Complex Object Tags
	xmlobj := sTrim(SortSimpleObject(xml, obj, lvl), "\n")
	for _, l := range sSplit(xmlobj, "\n") {
		sl := 0
		switch {
		case sHasPrefix(l, S) && !sHasPrefix(l, E) && sContains(l, "..."): // Complex Object Tags
			sl = len(S)
		default:
			continue
		}
		oa := cmn.RmTailFromFirstAny(l[sl:], " ", ">")
		lvlOAs = append(lvlOAs, oa)
	}

	path := parent + "~" + obj
	if _, ok := mPathIdx[path]; !ok {
		mPathIdx[path] = 0
	}

	ipath := fSf("%s@%d", path, mPathIdx[path])
	mPathIdx[path]++
	if parent == "" { // root is without @index
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

// ----------------------------------------------- //

// JSON2XML1 : Disordered, Formatted from JSON
func JSON2XML1(jsonPath, xmlPath string) string {
	jsonBytes, err := ioutil.ReadFile(jsonPath)
	cmn.FailOnErr("%v", err)
	cmn.FailOnErrWhen(!cmn.IsJSON(string(jsonBytes)), "", fEf("Input File is not a valid JSON File"))

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

	// ioutil.WriteFile(xmlPath, []byte(xmlstr), 0666)
	return xmlstr
}

// JSON2XML2 : Ordered, Some pieces are different
func JSON2XML2(xml1, SIFSpecPath string) string {
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

	// Init Spec Maps
	InitMapOfObjAttrs(xpathGrp, SEP)

	// Init "mIPathSubXML"
	root := cmn.XMLRoot(xml1)
	ExtractOA(xml1, root, "", 0)

	xmlobj := mIPathSubXML[root]
AGAIN:
	for k, subxml := range mIPathSubXML {
		mark := mIPathSubMark[k]
		xmlobj = sReplace(xmlobj, mark, subxml, 1)
	}
	if sContains(xmlobj, "...") {
		// ioutil.WriteFile(fSf("./%d.xml", nGoTo), []byte(xmlobj), 0666)
		nGoTo++
		cmn.FailOnErrWhen(nGoTo > maxGoTo, "%v", fEf("goto AGAIN deadlock"))
		goto AGAIN
	}

	return xmlobj
}

// JSON2XML3 : Pieces Replaced, should be almost identical to Original SIF
func JSON2XML3(xml2 string, mRepl map[string]string) string {

	// remove @Number#
	r := regexp.MustCompile("@([0-9]+)#")
	xml2 = string(r.ReplaceAll([]byte(xml2), []byte("")))

	// others from cfg
	for old, new := range mRepl {
		xml2 = sReplaceAll(xml2, old, new)
	}

	return xml2
}
