package cvt2xml

import (
	"io/ioutil"
	"regexp"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func CountHeadSpace(s string, nGrp int) int {
	for i, c := range s {
		if c == ' ' {
			continue
		}
		return i / nGrp
	}
	return 0
}

func TagFromXMLLine(line string) (tag string, mKeyAttr map[string]string) {
	line = sTrim(line, " \t\n\r")
	cmn.FailOnErrWhen(line[0] != '<' || line[len(line)-1] != '>', "XML Err @ %v", fEf(line))
	if tag := regexp.MustCompile(`<.+[> ]`).FindString(line); tag != "" {
		tag = tag[1 : len(tag)-1] // remove '<' '>'
		ss := sSplit(tag, " ")    // cut fields
		mKeyAttr = make(map[string]string)
		for _, attr := range ss[1:] {
			if ak := regexp.MustCompile(`.+="`).FindString(attr); ak != "" {
				mKeyAttr[ak[:len(ak)-2]] = attr // remove '="'
			}
		}
		return ss[0], mKeyAttr
	}
	return "", nil
}

func Hierarchy(searchArea string, lvl int, hierarchy *[]string) {
	r := regexp.MustCompile(fSf(`\n[ ]{%d}<.*>`, (lvl-1)*4))
	if locGrp := r.FindAllStringIndex(searchArea, -1); locGrp != nil {
		loc := locGrp[len(locGrp)-1]
		start, end := loc[0], loc[1]
		find := searchArea[start:end]
		Hierarchy(searchArea[:start], lvl-1, hierarchy)
		tag, _ := TagFromXMLLine(find)
		*hierarchy = append(*hierarchy, tag)
	}
}

func SearchTagWithAttr(xml string) {
	root := cmn.XMLRoot(xml)
	TagOrAttr, minAttr := `[^ \t<>]+`, 2
	r := regexp.MustCompile(fSf(`[ ]*<%[1]s[ ]+(%[1]s="%[1]s"[ ]*){%d,}>`, TagOrAttr, minAttr))
	if loc := r.FindAllStringIndex(xml, -1); loc != nil {
		for _, l := range loc {
			fPln("---------------------------------------------")
			hierarchy := &[]string{root}
			start, end := l[0], l[1]
			withAttr := xml[start:end]

			fPln(start)

			Hierarchy(xml[:start], CountHeadSpace(withAttr, 4), hierarchy)
			tag, mka := TagFromXMLLine(withAttr)
			*hierarchy = append(*hierarchy, tag)
			fPln(sJoin(*hierarchy, "/"))
			fPln(mka)
		}
	}
}

func TestSearchTagWithAttr(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/xml/AGAddressCollectionSubmission_3_out.xml")
	cmn.FailOnErr("%v", err)
	xml := string(bytes)
	cmn.FailOnErrWhen(!cmn.IsXML(xml), "%v", fEf("Not XML"))
	SearchTagWithAttr(xml)
}

func TestGetPath(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/xml/AGAddressCollectionSubmission_3_out.xml")
	cmn.FailOnErr("%v", err)
	cmn.FailOnErrWhen(!cmn.IsXML(string(bytes)), "%v", fEf("Not XML"))

	// tag := "StudentAddress"

}
