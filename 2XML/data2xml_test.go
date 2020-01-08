package cvt2xml

import "testing"

func TestJSON2XML(t *testing.T) {
	json2xml("../data/test.json", "../data/test2.xml")
}
