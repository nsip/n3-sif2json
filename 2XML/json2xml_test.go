package cvt2xml

import (
	"io/ioutil"
	"testing"
)

func TestJSON2XML(t *testing.T) {
	xml1 := JSON2XML1("../data/Activity.json", "../data/Activity1.xml")
	xml2 := JSON2XML2("../SIFSpec/out.txt", xml1)
	xml3 := JSON2XML3(xml2, getReplMap("./SIFCfg/replace.json"))
	ioutil.WriteFile("../data/Activity2.xml", []byte(xml3), 0666)
}
