package cvt2xml

import (
	"io/ioutil"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestJSON2XML(t *testing.T) {
	obj := "NAPStudentResponseSet"

	xml1 := JSON2XML1(fSf("../data/%s.json", obj))
	ioutil.WriteFile(fSf("../data/%s_1_out.xml", obj), []byte(xml1), 0666)

	xml2 := JSON2XML2(xml1, "../SIFSpec/out.txt")
	ioutil.WriteFile(fSf("../data/%s_2_out.xml", obj), []byte(xml2), 0666)

	xml3 := JSON2XML3(xml2, getReplMap("./SIFCfg/replace.json"))
	ioutil.WriteFile(fSf("../data/%s_3_out.xml", obj), []byte(xml3), 0666)
}

func TestSortSimpleObject(t *testing.T) {
	jsonBytes, err := ioutil.ReadFile("../data/AGAddressCollectionSubmission_1_out.xml")
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

	// "EntityContact"
	fPln(SortSimpleObject(sifCont, "Name", 4))
	// fPln(SortSimpleObject(sifCont, "ReportExclusionFlag", 1))
	// fPln("-----------------------")
	// fPln(SortSimpleObject(sifCont, "ItemResponseList", 3))
	// fPln(SortSimpleObject(sifCont, "ItemResponse", 4))
	// fPln(SortSimpleObject(sifCont, "ItemResponse", 4))
	// fPln(SortSimpleObject(sifCont, "ItemResponse", 4))

	// ExtractOA(sifCont, "NAPStudentResponseSet", "", 0)

}
