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

		// obj := "LearningResource"
		// if !cmn.XIn(obj, []string{"LearningStandardDocument", "StudentAttendanceTimeList"}) {
		// 	continue
		// }

		// JSON2XML0 : deal with XML multiple-line content
		jsonWithCode, mCodeStr := JSON2SIF4LF(fSf("../data/json/%s.json", obj))

		xml := JSON2SIF3RD(jsonWithCode)
		// ioutil.WriteFile(fSf("../data/xml/%s_0_out.xml", obj), []byte(xml), 0666)

		xml1 := JSON2SIFViaSpec(xml, "../SIFSpec/out.txt")
		// ioutil.WriteFile(fSf("../data/xml/%s_1_out.xml", obj), []byte(xml1), 0666)

		// doing ...
		posGrp, pathGrp, mAttrGrp := SearchTagWithAttr(xml)
		for i, path := range pathGrp {
			fPln("--------------------------------------------")
			attrs := []string{}
			for _, trvs := range TrvsGrpViaSpec {
				if sHasPrefix(trvs, path+"/@") {
					fPln(trvs)
					attrs = append(attrs, sSplit(trvs, "\t")[2][1:]) // from Spec format
				}
			}
			fPln(attrs)

			xmlLine := xml[posGrp[i][0]:posGrp[i][1]]
			fPln(xmlLine)
			tag, _ := TagFromXMLLine(xmlLine)
			fPln(tag)
			n := CountHeadSpace(xmlLine, 1)
			fPln(n)

			m := mAttrGrp[i]
			fPln(m)

			out := fSf("")
		}
		// doing ...

		mRepl := cmn.MapsMerge(getReplMap("./SIFCfg/replace.json"), mCodeStr).(map[string]string)
		xml2 := JSON2SIFRepl(xml1, mRepl)
		ioutil.WriteFile(fSf("../data/xml/%s_2_out.xml", obj), []byte(xml2), 0666)
	}
}

func TestSortSimpleObject(t *testing.T) {
	const TRAVERSE = "TRAVERSE ALL, DEPTH ALL"

	bytes, err := ioutil.ReadFile("../SIFSpec/out.txt")
	cmn.FailOnErr("%v", err)
	spec := string(bytes)

	for _, line := range sSplit(spec, "\n") {
		switch {
		case sHasPrefix(line, TRAVERSE):
			l := sTrim(line[len(TRAVERSE):], " \t\r")
			TrvsGrpViaSpec = append(TrvsGrpViaSpec, l)
		}
	}

	// Init Spec Maps
	InitOAs(TrvsGrpViaSpec, "\t", "/")

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
