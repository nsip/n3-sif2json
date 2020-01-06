package cvt2json

import "testing"

func TestXML2JSON(t *testing.T) {
	xml2json("./config/XML2JSON.toml", "../data/test.xml", "../data/test.json")
}
