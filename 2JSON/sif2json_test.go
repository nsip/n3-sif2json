package cvt2json

import (
	"io/ioutil"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestJSONRoot(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/Activity.json")
	cmn.FailOnErr("%v", err)
	fPln(cmn.JSONRoot(string(bytes)))
}

func TestXMLRoot(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/Activity.xml")
	cmn.FailOnErr("%v", err)
	// fPln(string(bytes))
	fPln(cmn.XMLRoot(string(bytes)))
}

func TestEachFileContent(t *testing.T) {
	fPln(eachFileContent("../data/ListAttributes/PurchaseOrder", "json", cmn.Iter2Slc(10)...))
}

func TestSIF2JSON(t *testing.T) {

	dir := `../data/examples/`
	files, err := ioutil.ReadDir(dir)
	cmn.FailOnErr("%v", err)

	for _, file := range files {
		obj := cmn.RmTailFromLast(file.Name(), ".")
		fPln(obj)

		if obj == "LearningStandardDocument" ||
			obj == "LearningStandardItem" ||
			obj == "NAPCodeFrame" ||
			obj == "NAPTestItem" ||
			obj == "StudentAttendanceTimeList" {
			continue
		}

		// obj := "LearningStandardDocument"
		bytes, err := ioutil.ReadFile(fSf("../data/examples/%s.xml", obj))
		cmn.FailOnErr("%v", err)
		json, sv, err := SIF2JSON("./config/SIF2JSON.toml", string(bytes), "3.4.5X", false)
		fPln(obj, sv, err)
		ioutil.WriteFile(fSf("../data/%s.json", obj), []byte(json), 0666)

	}
}
