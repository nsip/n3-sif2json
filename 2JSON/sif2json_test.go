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
	bytes, err := ioutil.ReadFile("../data/AGAddressCollectionSubmission.xml")
	cmn.FailOnErr("%v", err)
	json, sv, err := SIF2JSON("./config/SIF2JSON.toml", string(bytes), "3.4.5X", false)
	fPln(sv, err)
	ioutil.WriteFile("../data/AGAddressCollectionSubmission.json", []byte(json), 0666)
}
