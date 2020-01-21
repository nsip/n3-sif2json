package cvt2json

import (
	"io/ioutil"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestJSONRoot(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/test.json")
	cmn.FailOnErr("%v", err)
	fPln(jsonroot(string(bytes)))
}

func TestXMLRoot(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/test.xml")
	cmn.FailOnErr("%v", err)
	// fPln(string(bytes))
	fPln(xmlroot(string(bytes)))
}

func TestGetEachFileContent(t *testing.T) {
	fPln(getEachFileContent("../data/ListAttributes/PurchaseOrder", "json", cmn.Iter2Slc(10)...))
}

func TestSIF2JSON(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/Activity.xml")
	cmn.FailOnErr("%v", err)
	json := SIF2JSON("./config/SIF2JSON.toml", string(bytes), "0.0.1", false)
	ioutil.WriteFile("../data/Activity.json", []byte(json), 0666)
}
