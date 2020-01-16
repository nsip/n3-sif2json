package cvt2xml

import "testing"

func TestJSON2XML(t *testing.T) {
	json2xml("../data/Activity.json", "../data/Activity1.xml")
}
