package cvt2xml

import (
	"encoding/json"
	"io/ioutil"

	"github.com/clbanning/mxj"
)

func json2xml(jsonPath, xmlPath string) {
	jsonBytes, err := ioutil.ReadFile(jsonPath)
	var f interface{}
	if err = json.Unmarshal(jsonBytes, &f); err != nil {
		panic("1")
	}

	fPln(f)

	b, err := mxj.AnyXmlIndent(f, "", "    ", "")
	xmlstr := string(b)
	xmlstr = sReplaceAll(xmlstr, "<>", "")
	xmlstr = sReplaceAll(xmlstr, "</>", "")
	xmlstr = re1.ReplaceAllString(xmlstr, "")
	xmlstr = re2.ReplaceAllString(xmlstr, "")
	xmlstr, _ = Indent(xmlstr, -4, false)
	xmlstr = sTrim(xmlstr, " \t\n")

	// var f1 interface{}
	// if b, err = xml.Marshal(&f1); err != nil {
	// 	panic("2")
	// }
	ioutil.WriteFile(xmlPath, []byte(xmlstr), 0666)

	// return
}
