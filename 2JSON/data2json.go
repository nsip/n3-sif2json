package cvt2json

import (
	"io/ioutil"

	xj "github.com/basgys/goxml2json"
	cmn "github.com/cdutwhu/json-util/common"
	pp "github.com/cdutwhu/json-util/preprocess"
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

func xml2json(cfgPath, xmlPath, jsonPath string) {

	cfg := NewCfg(cfgPath)
	cmn.FailOnCondition(cfg == nil, "%v", fEf("ListAttribute Configuration File Couldn't Be Loaded"))
	cfgPrefix := cfg.(*XML2JSON)

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
	cmn.FailOnErr("That's embarrassing... %v", err)

	jsonfmt := jsonBuf.String()
	fPln(jsonfmt)
	jsonfmt = pp.FmtJSONStr(jsonfmt, cfgPrefix.JQDir)
	fPln(jsonfmt)
	jsonfmt = replDigCont(jsonfmt)
	fPln(jsonfmt)
	ioutil.WriteFile(jsonPath, []byte(jsonfmt), 0666)
}
