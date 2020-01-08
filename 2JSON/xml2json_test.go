package cvt2json

import "testing"

func TestGetEachFileContent(t *testing.T) {
	fPln(getEachFileContent("../ListAttr/PurchaseOrder", "json", 1, 2, 3, 4, 5))
}

func TestXML2JSON(t *testing.T) {
	xml2json("./config/XML2JSON.toml", "../data/test.xml", "../data/test.json")
}
