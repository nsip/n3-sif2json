package cvt2json

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/cdutwhu/n3-util/n3err"
	sif346 "github.com/nsip/n3-sif2json/SIFSpec/3.4.6"
	sif347 "github.com/nsip/n3-sif2json/SIFSpec/3.4.7"
)

func TestXMLRoot(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/examples347/Activity_0.xml")
	failOnErr("%v", err)
	fPln(xmlRoot(string(bytes)))
}

func TestEachFileContent(t *testing.T) {
	fPln(eachFileContent("../data/ListAttributes/PurchaseOrder", "json", iter2Slc(10)...))
}

func s2j(dim int, tid int, done chan int, params ...interface{}) {
	defer func() { done <- tid }()
	files := params[0].([]os.FileInfo)
	L := len(files)
	for i := tid; i < L; i += dim {
		obj := rmTailFromLast(files[i].Name(), ".")
		fPln("start:", obj)
		// if exist(obj, "LearningStandardDocument", "StudentAttendanceTimeList") {
		// 	continue
		// }
		bytes, err := ioutil.ReadFile(fSf("../data/examples/3.4.7/%s.xml", obj))
		failOnErr("%v", err)
		json, sv, err := SIF2JSON(string(bytes), "3.4.7", false)
		fPln("end:", obj, sv, err)
		if json != "" {
			mustWriteFile(fSf("../data/json/%s/%s.json", sv, obj), []byte(json))
		}
	}
}

func TestSIF2JSON(t *testing.T) {
	enableLog2F(true, "./error.log")
	defer enableLog2F(false, "")

	// Test Resource
	fPln(string(sif346.JSON["BOOLEAN_Activity_4"]))
	fPln(string(sif347.JSON["BOOLEAN_Activity_4"]))
	// Test End

	dir := `../data/examples/3.4.7/`
	files, err := ioutil.ReadDir(dir)
	failOnErr("%v", err)
	failOnErrWhen(len(files) == 0, "%v", n3err.FILE_NOT_FOUND)

	Go(1, s2j, files)
	fPln("OK")
}
