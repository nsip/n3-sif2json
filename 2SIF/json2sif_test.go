package cvt2sif

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/cdutwhu/n3-util/n3err"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-xmlfmt/xmlfmt"
)

func TestConfig(t *testing.T) {
	cfg := &Config{}
	n3cfg.New(cfg, nil, "./config.toml")
	spew.Dump(cfg)
}

func j2s(dim int, tid int, done chan int, params ...interface{}) {
	defer func() { done <- tid }()
	files := params[0].([]os.FileInfo)
	ver := params[1].(string)
	L := len(files)
	for i := tid; i < L; i += dim {
		ResetAll()

		obj := rmTailFromLast(files[i].Name(), ".")
		bytes, err := ioutil.ReadFile(fSf("../data/json/%s/%s.json", ver, obj))
		failOnErr("%v", err)

		sif, sv, err := JSON2SIF("./config.toml", string(bytes), ver)
		failOnErr("%v", err)

		sif = xmlfmt.FormatXML(sif, "", "    ")
		sif = sTrim(sif, " \t\n\r")

		fPln(obj, sv, " applied")
		if sif != "" {
			mustWriteFile(fSf("../data/sif/%s/%s_out.xml", sv, obj), []byte(sif))
		}
	}
}

func TestJSON2SIF(t *testing.T) {
	enableLog2F(true, "./error.log")
	defer enableLog2F(false, "")

	ver := "3.4.7"
	dir := `../data/json/` + ver
	files, err := ioutil.ReadDir(dir)
	failOnErr("%v", err)
	failOnErrWhen(len(files) == 0, "%v", n3err.FILE_NOT_FOUND)

	Go(1, j2s, files, ver) // only dispatch 1 goroutine, otherwise, crash

	// for _, file := range files {
	// 	ResetAll()

	// 	obj := rmTailFromLast(file.Name(), ".")

	// 	// if obj == "Activity2" {
	// 	// 	continue
	// 	// }

	// 	fPln("------------", obj)
	// 	bytes, err := ioutil.ReadFile(fSf("../data/json/%s/%s.json", ver, obj))
	// 	failOnErr("%v", err)

	// 	sif, sv, err := JSON2SIF("./config/config.toml", string(bytes), ver)
	// 	failOnErr("%v", err)

	// 	sif = xmlfmt.FormatXML(sif, "", "    ")
	// 	sif = sTrim(sif, " \t\n\r")

	// 	fPln(sv + " is used")
	// 	if sif != "" {
	// 		mustWriteFile(fSf("../data/sif/%s/%s_out.xml", sv, obj), []byte(sif))
	// 	}

	// 	// {
	// 	// 	// JSON2SIF4LF : deal with XML multiple-line content
	// 	// 	jsonWithCode, mCodeStr := JSON2SIF4LF(string(bytes))

	// 	// 	xml := JSON2SIF3RD(jsonWithCode)
	// 	// 	// ioutil.WriteFile(fSf("../data/sif/%s_0_out.xml", obj), []byte(xml), 0666)

	// 	// 	xml1 := JSON2SIFSpec(xml, "../SIFSpec/out.txt") // sv is here
	// 	// 	// ioutil.WriteFile(fSf("../data/sif/%s_1_out.xml", obj), []byte(xml1), 0666)

	// 	// 	mRepl := mapsMerge(getReplMap("./config/replace.json"), mCodeStr).(map[string]string)
	// 	// 	xml2 := JSON2SIFRepl(xml1, mRepl)
	// 	// 	ioutil.WriteFile(fSf("../data/sif/%s_out.xml", obj), []byte(xml2), 0666)
	// 	// }
	// }

	fPln("OK")
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
	failOnErr("%v", err)
	sifCont := string(jsonBytes)

	fPln(SortSimpleObject(sifCont, "AGRule", 4, "AGStatusReport/AGReportingObjectResponseList/AGReportingObjectResponse/AGRuleList/"))

	// ExtractOA(sifCont, "NAPStudentResponseSet", "", 0)
}
