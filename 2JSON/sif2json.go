package cvt2json

import (
	"io/ioutil"
	"os"

	xj "github.com/basgys/goxml2json"
	cmn "github.com/cdutwhu/json-util/common"
	jkv "github.com/cdutwhu/json-util/jkv"
	pp "github.com/cdutwhu/json-util/preprocess"
)

// func replaceDigCont(json, jqDir string) string {
// 	json = pp.FmtJSONStr(json, jqDir)
// 	m4repl := contValRepl(reContDigVal.FindAllString(json, -1))
// 	for oldstr, newstr := range m4repl {
// 		json = sReplaceAll(json, oldstr, newstr)
// 	}
// 	return json

// 	// for _, pair := range reContDigVal.FindAllStringIndex(jsonfmt, -1) {
// 	// 	str := jsonfmt[pair[0]:pair[1]]
// 	// 	fPln(contDigVal(str))
// 	// }
// }

func getEachFileContent(dir, ext string, indices ...int) (rt []string) {
	if dir[len(dir)-1] != '/' {
		dir += "/"
	}
	if ext[0] == '.' {
		ext = ext[1:]
	}
	files := []string{}
	for _, index := range indices {
		file := fSf("%s%d.%s", dir, index, ext)
		if _, err := os.Stat(file); err == nil {
			files = append(files, file)
		}
	}
	for _, f := range files {
		bytes, err := ioutil.ReadFile(f)
		cmn.FailOnErr("%v", err)
		rt = append(rt, string(bytes))
	}
	return
}

// enforceConfig : LIST config must be from low Level to high level
func enforceConfig(json, jqDir string, lsJSONCfg ...string) string {
	for _, jsoncfg := range lsJSONCfg {
		// make sure [jsoncfg] is formatted
		// otherwise, do Fmt firstly
		// jsoncfg = pp.FmtJSONStr(jsoncfg, jqDir)
		maskroot, _ := jkv.NewJKV(json, "", false).Unfold(0, jkv.NewJKV(jsoncfg, "", false))
		json = pp.FmtJSONStr(maskroot, jqDir)
	}
	return json
}

// SIF2JSON :
func SIF2JSON(cfgPath, xmlPath, jsonPath string, enforced ...string) {
	ICfg := NewCfg(cfgPath)
	cmn.FailOnCondition(ICfg == nil, "%v", fEf("sif2json.toml couldn't be Loaded"))
	cfg := ICfg.(*sif2json)

	bytesXML, err := ioutil.ReadFile(xmlPath)
	cmn.FailOnErr("%v", err)
	// fPln(string(bytesXML))

	// xml is an io.Reader
	xml := string(bytesXML)
	xmlReader := sNewReader(xml)
	jsonBuf, err := xj.Convert(
		xmlReader,
		// xj.WithTypeConverter(xj.Float, xj.Int, xj.Bool, xj.Null),
		// xj.WithAttrPrefix("-"),
		// xj.WithContentPrefix("#"),
	)
	cmn.FailOnErr("That's embarrassing... %v", err)

	json := pp.FmtJSONStr(jsonBuf.String(), cfg.JQDir)

	// Digital string to number
	// json := replaceDigCont(json, cfg.JQDir)

	// List Attributes Modification
	obj := xmlroot(xml)    // infer object from xml root by default, use this object to search config json
	if len(enforced) > 0 { // if object is provided, ignore default, use 1st provided object to search config json
		obj = enforced[0]
	}
	lsAttrRule := getEachFileContent(cfg.CfgJSONDir4LIST+obj, "json", cmn.Iter2Slc(10)...)
	json = enforceConfig(json, cfg.JQDir, lsAttrRule...)

	ioutil.WriteFile(jsonPath, []byte(json), 0666)
}
