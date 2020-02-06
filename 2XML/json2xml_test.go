package cvt2xml

import (
	"io/ioutil"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestJSON2XML(t *testing.T) {
	xml1 := JSON2XML1("../data/NAPStudentResponseSet.json", "../data/NAPStudentResponseSet1.xml")
	xml2 := JSON2XML2(xml1, "../SIFSpec/out.txt")
	xml3 := JSON2XML3(xml2, getReplMap("./SIFCfg/replace.json"))
	ioutil.WriteFile("../data/NAPStudentResponseSet3.xml", []byte(xml3), 0666)
}

func TestSortSimpleObject(t *testing.T) {
	jsonBytes, err := ioutil.ReadFile("../data/NAPStudentResponseSet3.xml")
	cmn.FailOnErr("%v", err)
	sifCont := string(jsonBytes)

	const (
		SEP       = "\t"
		XPATHTYPE = "XPATHTYPE:"
	)

	bytes, err := ioutil.ReadFile("../SIFSpec/out.txt")
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

	fPln(SortSimpleObject(sifCont, "NAPStudentResponseSet", 0))
	// fPln(SortSimpleObject(sifCont, "ReportExclusionFlag", 1))
	// fPln("-----------------------")
	// fPln(SortSimpleObject(sifCont, "ItemResponseList", 3))
	// fPln(SortSimpleObject(sifCont, "ItemResponse", 4))
	// fPln(SortSimpleObject(sifCont, "ItemResponse", 4))
	// fPln(SortSimpleObject(sifCont, "ItemResponse", 4))

	// ExtractOA(sifCont, "NAPStudentResponseSet", "", 0)

}
