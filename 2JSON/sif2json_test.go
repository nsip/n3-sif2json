package cvt2json

import (
	"io/ioutil"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

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
	SIF2JSON("./config/sif2json.toml", "../data/test.xml", "../data/test.json")
}
