package cvt2xml

import (
	"io/ioutil"
	"testing"
)

func TestJSON2XML(t *testing.T) {
	xml1 := JSON2XML1("../data/Activity.json", "../data/Activity1.xml")
	xml2 := JSON2XML2("./out.txt", xml1)
	ioutil.WriteFile("../data/Activity2.xml", []byte(xml2), 0666)
}
