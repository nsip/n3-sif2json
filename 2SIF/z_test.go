package cvt2xml

import (
	"io/ioutil"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestSearchTagWithAttr(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/xml/AGAddressCollectionSubmission_3_out.xml")
	cmn.FailOnErr("%v", err)
	xml := string(bytes)
	cmn.FailOnErrWhen(!cmn.IsXML(xml), "%v", fEf("Not XML"))
	posGrp, pathGrp, mAttrGrp := SearchTagWithAttr(xml)
	for i, path := range pathGrp {
		fPln("--------------------------------------------")
		fPln(xml[posGrp[i][0]:posGrp[i][1]])
		fPln(path)
		fPln(mAttrGrp[i])
	}
}

func TestGetPath(t *testing.T) {
	cmn.SetLog("./error.log")
	bytes, err := ioutil.ReadFile("../data/xml/AGAddressCollectionSubmission_3_out.xml")
	cmn.FailOnErr("%v", err)
	cmn.FailOnErrWhen(!cmn.IsXML(string(bytes)), "%v", fEf("Not XML"))

	// tag := "StudentAddress"

	fPln(sReplByPos("abcdefg", 5, 6, "AAAA"))
}
