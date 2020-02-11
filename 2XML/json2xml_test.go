package cvt2xml

import (
	"io/ioutil"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestJSON2XML(t *testing.T) {
	cmn.SetLog("./error.log")

	dir := `../data/json/`
	files, err := ioutil.ReadDir(dir)
	cmn.FailOnErr("%v", err)

	for _, file := range files {
		obj := cmn.RmTailFromLast(file.Name(), ".")
		fPln("------------", obj)

		resetAll()

		// obj := "LearningResource"

		xml1 := JSON2XML1(fSf("../data/json/%s.json", obj))
		// ioutil.WriteFile(fSf("../data/xml/%s_1_out.xml", obj), []byte(xml1), 0666)

		xml2 := JSON2XML2(xml1, "../SIFSpec/out.txt")
		// ioutil.WriteFile(fSf("../data/xml/%s_2_out.xml", obj), []byte(xml2), 0666)

		xml3 := JSON2XML3(xml2, getReplMap("./SIFCfg/replace.json"))
		ioutil.WriteFile(fSf("../data/xml/%s_3_out.xml", obj), []byte(xml3), 0666)
	}
}

func TestSortSimpleObject(t *testing.T) {
	const (
		TRAVERSE = "TRAVERSE ALL, DEPTH ALL"
	)

	bytes, err := ioutil.ReadFile("../SIFSpec/out.txt")
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

	fPln(NextAttr("ParentName", "AGAddressCollectionSubmission/AddressCollectionReportingList/AddressCollectionReporting/AddressCollectionStudentList/AddressCollectionStudent/Parent1/"))
	fPln(NextAttr("ParentName", "AGAddressCollectionSubmission/AddressCollectionReportingList/AddressCollectionReporting/AddressCollectionStudentList/AddressCollectionStudent/Parent1/"))
	fPln(NextAttr("ParentName", "AGAddressCollectionSubmission/AddressCollectionReportingList/AddressCollectionReporting/AddressCollectionStudentList/AddressCollectionStudent/Parent1/"))
	fPln(NextAttr("ParentName", "AGAddressCollectionSubmission/AddressCollectionReportingList/AddressCollectionReporting/AddressCollectionStudentList/AddressCollectionStudent/Parent1/"))

	// fPln(NextAttr("Name", "FinancialQuestionnaireSubmission/FQReportingList/FQReporting/EntityContact/"))
	// fPln(NextAttr("Name", "FinancialQuestionnaireSubmission/FQReportingList/FQReporting/EntityContact/"))
	// fPln(NextAttr("Name", "FinancialQuestionnaireSubmission/FQReportingList/FQReporting/EntityContact/"))
	// fPln(NextAttr("Name", "FinancialQuestionnaireSubmission/FQReportingList/FQReporting/EntityContact/"))
	// fPln(NextAttr("Name", "FinancialQuestionnaireSubmission/FQReportingList/FQReporting/EntityContact/"))

	return

	// jsonBytes, err := ioutil.ReadFile("../data/AGAddressCollectionSubmission_1_out.xml")
	// cmn.FailOnErr("%v", err)
	// sifCont := string(jsonBytes)

	// fPln(SortSimpleObject(sifCont, "Name", 4))
	// fPln(SortSimpleObject(sifCont, "ReportExclusionFlag", 1))
	// fPln("-----------------------")
	// fPln(SortSimpleObject(sifCont, "ItemResponseList", 3))
	// fPln(SortSimpleObject(sifCont, "ItemResponse", 4))
	// fPln(SortSimpleObject(sifCont, "ItemResponse", 4))
	// fPln(SortSimpleObject(sifCont, "ItemResponse", 4))

	// ExtractOA(sifCont, "NAPStudentResponseSet", "", 0)

}
