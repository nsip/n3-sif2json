package cvt2sif

import (
	"io/ioutil"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestJSON2SIF(t *testing.T) {
	cmn.SetLog("./error.log")
	defer cmn.ResetLog()

	dir := `../data/json/`
	files, err := ioutil.ReadDir(dir)
	cmn.FailOnErr("%v", err)
	cmn.FailOnErrWhen(len(files) == 0, "%v", fEf("no json files prepared"))

	for _, file := range files {
		ResetAll()

		obj := cmn.RmTailFromLast(file.Name(), ".")
		// fPln("------------", obj)

		bytes, err := ioutil.ReadFile(fSf("../data/json/%s.json", obj))
		cmn.FailOnErr("%v", err)

		if sif, sv, err := JSON2SIF("./config/JSON2SIF.toml", string(bytes), "3.4.5X"); err == nil {
			fPln(sv + " is used")
			ioutil.WriteFile(fSf("../data/sif/%s_out.xml", obj), []byte(sif), 0666)
		} else {
			fPln(err.Error())
		}

		// {
		// 	// JSON2SIF4LF : deal with XML multiple-line content
		// 	jsonWithCode, mCodeStr := JSON2SIF4LF(string(bytes))

		// 	xml := JSON2SIF3RD(jsonWithCode)
		// 	// ioutil.WriteFile(fSf("../data/sif/%s_0_out.xml", obj), []byte(xml), 0666)

		// 	xml1 := JSON2SIFSpec(xml, "../SIFSpec/out.txt") // sv is here
		// 	// ioutil.WriteFile(fSf("../data/sif/%s_1_out.xml", obj), []byte(xml1), 0666)

		// 	mRepl := cmn.MapsMerge(getReplMap("./config/replace.json"), mCodeStr).(map[string]string)
		// 	xml2 := JSON2SIFRepl(xml1, mRepl)
		// 	ioutil.WriteFile(fSf("../data/sif/%s_out.xml", obj), []byte(xml2), 0666)
		// }
	}
}

func TestSortSimpleObject(t *testing.T) {
	// Init Spec Maps
	InitOAs("../SIFSpec/out.txt", "\t", "/")

	// fPln(NextAttr("ParentName", "AGAddressCollectionSubmission/AddressCollectionReportingList/AddressCollectionReporting/AddressCollectionStudentList/AddressCollectionStudent/Parent1/"))
	// fPln(NextAttr("ParentName", "AGAddressCollectionSubmission/AddressCollectionReportingList/AddressCollectionReporting/AddressCollectionStudentList/AddressCollectionStudent/Parent1/"))
	// fPln(NextAttr("ParentName", "AGAddressCollectionSubmission/AddressCollectionReportingList/AddressCollectionReporting/AddressCollectionStudentList/AddressCollectionStudent/Parent1/"))
	// fPln(NextAttr("ParentName", "AGAddressCollectionSubmission/AddressCollectionReportingList/AddressCollectionReporting/AddressCollectionStudentList/AddressCollectionStudent/Parent1/"))

	// fPln(NextAttr("Name", "FinancialQuestionnaireSubmission/FQReportingList/FQReporting/EntityContact/"))
	// fPln(NextAttr("Name", "FinancialQuestionnaireSubmission/FQReportingList/FQReporting/EntityContact/"))
	// fPln(NextAttr("Name", "FinancialQuestionnaireSubmission/FQReportingList/FQReporting/EntityContact/"))
	// fPln(NextAttr("Name", "FinancialQuestionnaireSubmission/FQReportingList/FQReporting/EntityContact/"))
	// fPln(NextAttr("Name", "FinancialQuestionnaireSubmission/FQReportingList/FQReporting/EntityContact/"))

	jsonBytes, err := ioutil.ReadFile("../data/sif/AGStatusReport_0_out.xml")
	cmn.FailOnErr("%v", err)
	sifCont := string(jsonBytes)

	fPln(SortSimpleObject(sifCont, "AGRule", 4, "AGStatusReport/AGReportingObjectResponseList/AGReportingObjectResponse/AGRuleList/"))

	// ExtractOA(sifCont, "NAPStudentResponseSet", "", 0)
}
