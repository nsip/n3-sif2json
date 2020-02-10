package cvt2xml

import (
	"encoding/json"
	"io/ioutil"
	"regexp"

	cmn "github.com/cdutwhu/json-util/common"
	"github.com/clbanning/mxj"
)

// ----------------------------------------- //

// InitOAs : trvsGrp is from SIF Spec txt
func InitOAs(trvsGrp []string, tblSep, pathSep string) {
	for _, trvs := range trvsGrp {
		path := sSplit(trvs, tblSep)[0]
		key := cmn.RmTailFromLast(path, pathSep)
		value := cmn.RmHeadToLast(path, pathSep)
		mPathAttrs[key] = append(mPathAttrs[key], value)
		mPathAttrIdx[key] = 0
	}
}

// NextAttr : From Spec
func NextAttr(obj, fullpath string) (value string, end bool) {
	if fullpath == "/" {
		fullpath = obj
	} else {
		fullpath += obj
	}
	idx := mPathAttrIdx[fullpath]
	if idx == len(mPathAttrs[fullpath]) {
		return "", true
	}
	defer func() {
		mPathAttrIdx[fullpath]++
	}()
	return mPathAttrs[fullpath][idx], false
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
func SortSimpleObject(xml, obj string, level int, trvsPath string) (paper string) {
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

	rewindAttrIter()
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

	attr, end := NextAttr(obj, trvsPath)
	for ; !end; attr, end = NextAttr(obj, trvsPath) {
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

// From : AGAddressCollectionSubmission~AddressCollectionReportingList@0~AddressCollectionReporting@0~EntityContact@0
// To :   AGAddressCollectionSubmission/AddressCollectionReportingList/AddressCollectionReporting/EntityContact/
func iPath2SpecPath(iPath, oldSep, newSep string) string {
	ss := sSpl(iPath, oldSep)
	for i, s := range ss {
		ss[i] = cmn.RmTailFromLast(s, "@")
	}
	return sJoin(ss, newSep) + newSep
}

// ExtractOA : root obj, path is ""
func ExtractOA(xml, obj, path string, lvl int) string {
	S := mkIndent(lvl+1) + "<"
	E := S + "/"

	iPathSep, specTrvsPathSep := "~", "/"

	// fPln(path)
	// fPln(iPath2SpecPath(path, iPathSep, specTrvsPathSep))
	specTrvsPath := iPath2SpecPath(path, iPathSep, specTrvsPathSep)

	lvlOAs := []string{} // Complex Object Tags
	xmlobj := sTrim(SortSimpleObject(xml, obj, lvl, specTrvsPath), "\n")
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

	if path != "" {
		path += (iPathSep + obj)
	} else {
		path = obj
	}
	if _, ok := mPathIdx[path]; !ok {
		mPathIdx[path] = 0
	}

	iPath := fSf("%s@%d", path, mPathIdx[path])
	mPathIdx[path]++
	if path == obj { // root is without @index
		iPath = obj
	}

	mIPathSubXML[iPath] = xmlobj

	xmlobjLn1 := sSplit(xmlobj, "\n")[0]
	preBlank := mkIndent(sCount(iPath, iPathSep))
	mIPathSubMark[iPath] = fSf("%s...\n%s</%s>", xmlobjLn1, preBlank, obj)

	for _, subobj := range lvlOAs {
		ExtractOA(xml, subobj, iPath, lvl+1)
	}

	return xmlobj
}

// ----------------------------------------------- //

// JSON2XML1 : Disordered, Formatted from JSON
func JSON2XML1(jsonPath string) string {
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

	return xmlstr
}

// JSON2XML2 : Ordered, Some pieces are different
func JSON2XML2(xml1, SIFSpecPath string) string {
	const (
		TRAVERSE = "TRAVERSE ALL, DEPTH ALL"
	)

	bytes, err := ioutil.ReadFile(SIFSpecPath)
	cmn.FailOnErr("%v", err)
	spec := string(bytes)

	for _, line := range sSplit(spec, "\n") {
		switch {
		case sHasPrefix(line, TRAVERSE):
			l := sTrim(line[len(TRAVERSE):], " \t\r")
			trvsGrp = append(trvsGrp, l)
		}
	}

	// Init Spec Maps
	InitOAs(trvsGrp, "\t", "/")

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
