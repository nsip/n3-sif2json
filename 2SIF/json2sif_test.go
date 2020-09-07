package cvt2sif

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cdutwhu/n3-util/n3err"
	"github.com/go-xmlfmt/xmlfmt"
)

func TestJSONRoot(t *testing.T) {
	bytes, err := ioutil.ReadFile("../data/Activity.json")
	failOnErr("%v", err)
	fPln(jsonRoot(string(bytes)))
}

func j2s(dim int, tid int, done chan int, params ...interface{}) {
	defer func() { done <- tid }()
	ver := params[0].(string)
	files := params[1].([]os.FileInfo)
	dir := params[2].(string)

	for i := tid; i < len(files); i += dim {
		ResetAll()

		obj := rmTailFromLast(files[i].Name(), ".")
		bytes, err := ioutil.ReadFile(filepath.Join(dir, files[i].Name()))
		failOnErr("%v", err)

		sif, sv, err := JSON2SIF(string(bytes), ver)
		failOnErr("%v", err)

		sif = xmlfmt.FormatXML(sif, "", "    ")
		sif = sTrim(sif, " \t\n\r")

		fPf("%-40s%-10sapplied\n", obj, sv)
		if sif != "" {
			mustWriteFile(fSf("../data/output/%s/sif/%s.xml", sv, obj), []byte(sif))
		}
	}
}

func TestJSON2SIF(t *testing.T) {
	defer trackTime(time.Now())
	// enableLog2F(true, "./error.log")
	// defer enableLog2F(false, "")

	ver := "3.4.6"
	dir := `../data/output/` + ver + `/json`
	files, err := ioutil.ReadDir(dir)
	failOnErr("%v", err)
	failOnErrWhen(len(files) == 0, "%v", n3err.FILE_NOT_FOUND)
	Go(1, j2s, ver, files, dir) // only dispatch 1 goroutine, otherwise, error
	fPln("OK")
}
