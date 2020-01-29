package cvt2xml

import (
	"encoding/json"
	"io/ioutil"

	cmn "github.com/cdutwhu/json-util/common"
	"github.com/clbanning/mxj"
)

func json2xml(jsonPath, xmlPath string) {
	jsonBytes, err := ioutil.ReadFile(jsonPath)
	cmn.FailOnErr("%v", err)

	var f interface{}
	cmn.FailOnErr("%v", json.Unmarshal(jsonBytes, &f))
	fPln(f)

	b, err := mxj.AnyXmlIndent(f, "", "    ", "")
	cmn.FailOnErr("%v", err)

	xmlstr := string(b)
	xmlstr = sReplaceAll(xmlstr, "<>", "")
	xmlstr = sReplaceAll(xmlstr, "</>", "")
	xmlstr = re1.ReplaceAllString(xmlstr, "")
	xmlstr = re2.ReplaceAllString(xmlstr, "")
	xmlstr, _ = Indent(xmlstr, -4, false)
	xmlstr = sTrim(xmlstr, " \t\n")

	ioutil.WriteFile(xmlPath, []byte(xmlstr), 0666)
}

// ----------------------------------------- //

// Lookup Object From config.txt

// Next

// Check Exist

// Print 1 Level

//
