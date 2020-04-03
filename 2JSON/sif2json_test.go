package cvt2json

import (
	"io/ioutil"
	"sync"
	"testing"

	cmn "github.com/cdutwhu/json-util/common"
)

func TestJSONRoot(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/Activity.json")
	cmn.FailOnErr("%v", err)
	fPln(cmn.JSONRoot(string(bytes)))
}

func TestXMLRoot(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/Activity.xml")
	cmn.FailOnErr("%v", err)
	// fPln(string(bytes))
	fPln(cmn.XMLRoot(string(bytes)))
}

func TestEachFileContent(t *testing.T) {
	fPln(eachFileContent("../data/ListAttributes/PurchaseOrder", "json", cmn.Iter2Slc(10)...))
}

func TestSIF2JSON(t *testing.T) {
	cmn.SetLog("./error.log")
	defer cmn.ResetLog()

	// bytes, err := ioutil.ReadFile("/home/qmiao/Desktop/attribute_test.xml")
	// cmn.FailOnErr("%v", err)
	// obj := "Activity"
	// sv := "3.4.6"
	// json, sv, err := SIF2JSON("./config/SIF2JSON.toml", string(bytes), sv, false)
	// // fPln("end:", obj, sv, err)
	// cmn.FailOnErr("%v", err)
	// if json != "" {
	// 	cmn.MustWriteFile(fSf("../data/json/%s/%s.json", sv, obj), []byte(json))
	// }
	// return

	dir := `../data/examples/`
	files, err := ioutil.ReadDir(dir)
	cmn.FailOnErr("%v", err)
	cmn.FailOnErrWhen(len(files) == 0, "%v", fEf("no xml files prepared"))

	wg := sync.WaitGroup{}
	wg.Add(len(files))

	for _, file := range files {
		obj := cmn.RmTailFromLast(file.Name(), ".")

		// if cmn.XIn(obj, []string{"LearningStandardDocument", "StudentAttendanceTimeList"}) {
		// 	continue
		// }

		go func(obj string) {
			defer wg.Done()

			fPln("start:", obj)
			bytes, err := ioutil.ReadFile(fSf("../data/examples/%s.xml", obj))
			cmn.FailOnErr("%v", err)
			sv := "3.4.6"
			json, sv, err := SIF2JSON("./config/SIF2JSON.toml", string(bytes), sv, false)
			fPln("end:", obj, sv, err)
			cmn.FailOnErr("%v", err)
			if json != "" {
				cmn.MustWriteFile(fSf("../data/json/%s/%s.json", sv, obj), []byte(json))
			}

		}(obj)
	}

	wg.Wait()
}
