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
		// fPln("------------", obj)

		resetAll()

		// if obj != "AGStatusReport" {
		// 	continue
		// }

		// JSON2XML0 : deal with XML multiple-line content
		jsonWithCode, mCodeStr := JSON2SIF4LF(fSf("../data/json/%s.json", obj))

		xml := JSON2SIF3RD(jsonWithCode)
		ioutil.WriteFile(fSf("../data/xml/%s_0_out.xml", obj), []byte(xml), 0666)

		xml1 := JSON2SIFSpec(xml, "../SIFSpec/out.txt")
		mRepl := cmn.MapsMerge(getReplMap("./SIFCfg/replace.json"), mCodeStr).(map[string]string)
		xml2 := JSON2SIFRepl(xml1, mRepl)
		ioutil.WriteFile(fSf("../data/xml/%s_2_out.xml", obj), []byte(xml2), 0666)
	}
}

func TestSortSimpleObject(t *testing.T) {
	// Init Spec Maps
	InitOAs("../SIFSpec/out.txt", "\t", "/")

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

func TestSearchTagWithAttr(t *testing.T) {
	cmn.SetLog("./error.log")
	bytes, err := ioutil.ReadFile("../data/xml/AGAddressCollectionSubmission_2_out.xml")
	cmn.FailOnErr("%v", err)
	xml := string(bytes)
	cmn.FailOnErrWhen(!cmn.IsXML(xml), "%v", fEf("Not XML"))

	fPln(sReplByPos("abcdefg", 5, 6, "AAAA"))
}
