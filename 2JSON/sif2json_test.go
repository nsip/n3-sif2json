package cvt2json

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/cdutwhu/n3-util/n3err"
)

func TestJSONRoot(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/Activity.json")
	failOnErr("%v", err)
	fPln(jsonRoot(string(bytes)))
}

func TestXMLRoot(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/Activity.xml")
	failOnErr("%v", err)
	// fPln(string(bytes))
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
		file := files[i]
		obj := rmTailFromLast(file.Name(), ".")
		fPln("start:", obj)
		// if exist(obj, "LearningStandardDocument", "StudentAttendanceTimeList") {
		// 	continue
		// }
		bytes, err := ioutil.ReadFile(fSf("../data/examples347/%s.xml", obj))
		failOnErr("%v", err)
		json, sv, err := SIF2JSON("./config/config.toml", string(bytes), "3.4.7", false)
		fPln("end:", obj, sv, err)
		if json != "" {
			mustWriteFile(fSf("../data/json/%s/%s.json", sv, obj), []byte(json))
		}
	}
}

func TestSIF2JSON(t *testing.T) {
	enableLog2F(true, "./error.log")
	defer enableLog2F(false, "")

	// bytes, err := ioutil.ReadFile("/home/qmiao/Desktop/attribute_test.xml")
	// failOnErr("%v", err)
	// obj := "Activity"
	// json, sv, err := SIF2JSON("./config/Config.toml", string(bytes), "3.4.7", false)
	// // fPln("end:", obj, sv, err)
	// failOnErr("%v", err)
	// if json != "" {
	// 	mustWriteFile(fSf("../data/json/%s/%s.json", sv, obj), []byte(json))
	// }
	// return

	dir := `../data/examples347/`
	files, err := ioutil.ReadDir(dir)
	failOnErr("%v", err)
	failOnErrWhen(len(files) == 0, "%v", n3err.FILE_NOT_FOUND)

	// wg := sync.WaitGroup{}
	// wg.Add(len(files))
	// for _, file := range files {
	// 	obj := rmTailFromLast(file.Name(), ".")
	// 	// if exist(obj, "LearningStandardDocument", "StudentAttendanceTimeList") {
	// 	// 	continue
	// 	// }
	// 	go func(obj string) {
	// 		defer wg.Done()
	// 		fPln("start:", obj)
	// 		bytes, err := ioutil.ReadFile(fSf("../data/examples347/%s.xml", obj))
	// 		failOnErr("%v", err)
	// 		json, sv, err := SIF2JSON("./config/config.toml", string(bytes), "3.4.7", false)
	// 		fPln("end:", obj, sv, err)
	// 		if json != "" {
	// 			mustWriteFile(fSf("../data/json/%s/%s.json", sv, obj), []byte(json))
	// 		}
	// 	}(obj)
	// }
	// wg.Wait()

	// Go(1, s2j, files)
	Go(len(files), s2j, files)
	fPln("OK")
}
