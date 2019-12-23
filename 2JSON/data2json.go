package cvt2json

import (
	"io/ioutil"

	pp "../preprocess"
	xj "github.com/basgys/goxml2json"
)

func replDigCont(json string) string {
	m4repl := contValRepl(reContDigVal.FindAllString(json, -1))
	for oldstr, newstr := range m4repl {
		json = sReplaceAll(json, oldstr, newstr)
	}
	return json

	// for _, pair := range reContDigVal.FindAllStringIndex(jsonfmt, -1) {
	// 	str := jsonfmt[pair[0]:pair[1]]
	// 	fPln(contDigVal(str))
	// }
}

func xml2json(xmlPath string) {

	b, _ := ioutil.ReadFile(xmlPath)
	xmlstr := string(b)
	fPln(xmlstr)

	// xml is an io.Reader
	xmlReader := sNewReader(xmlstr)
	jsonBuf, err := xj.Convert(
		xmlReader,
		xj.WithTypeConverter(xj.Float, xj.Int, xj.Bool, xj.Null),
		// xj.WithAttrPrefix("-"),
		// xj.WithContentPrefix("#"),
	)
	if err != nil {
		panic("That's embarrassing...")
	}

	jsonfmt := jsonBuf.String()
	fPln(jsonfmt)
	jsonfmt = pp.FmtJSONStr(jsonfmt, "../preprocess/utils")
	fPln(jsonfmt)
	jsonfmt = replDigCont(jsonfmt)
	fPln(jsonfmt)
	ioutil.WriteFile("../data/test.json", []byte(jsonfmt), 0666)
}
