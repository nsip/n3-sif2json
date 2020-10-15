package cvt2json

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cdutwhu/n3-util/n3err"
)

func TestXMLRoot(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/examples347/Activity_0.xml")
	failOnErr("%v", err)
	fPln(xmlRoot(string(bytes)))
}

func s2j(dim int, tid int, done chan int, params ...interface{}) {
	defer func() { done <- tid }()

	files := params[0].([]os.FileInfo)
	dir := params[1].(string)
	ver := params[2].(string)

	for i := tid; i < len(files); i += dim {
		obj := rmTailFromLast(files[i].Name(), ".")
		fPln("start:", obj)
		// if exist(obj, "LearningStandardDocument", "StudentAttendanceTimeList") {
		// 	continue
		// }
		bytes, err := ioutil.ReadFile(filepath.Join(dir, files[i].Name()))
		failOnErr("%v", err)
		json, sv, err := SIF2JSON(string(bytes), ver, false)
		fPln("end:", obj, sv, err)
		if json != "" {
			mustWriteFile(fSf("../data/output/%s/json/%s.json", sv, obj), []byte(json))
		}
	}
}

func TestSIF2JSON(t *testing.T) {
	defer trackTime(time.Now())
	// enableLog2F(true, "./error.log")
	// defer enableLog2F(false, "")
	// defer enableWarnDetail(true)
	enableWarnDetail(false)

	ver := "3.4.7"
	dir := `../data/examples/` + ver
	files, err := ioutil.ReadDir(dir)
	failOnErr("%v", err)
	failOnErrWhen(len(files) == 0, "%v", n3err.FILE_NOT_FOUND)
	syncParallel(4, s2j, files, dir, ver)
	fPln("OK")
}
