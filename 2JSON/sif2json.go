package cvt2json

import (
	"io/ioutil"
	"os"

	xj "github.com/basgys/goxml2json"
	cmn "github.com/cdutwhu/json-util/common"
	jkv "github.com/cdutwhu/json-util/jkv"
	pp "github.com/cdutwhu/json-util/preprocess"
	cfg "github.com/nsip/n3-sif2json/2JSON/config"
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

// SIF2JSON : if [SIFVer] is "", use config's DefaultSIFVer
func SIF2JSON(cfgPath, xml, SIFVer string, enforced bool, subobj ...string) (json string) {
	const (
		SignSIFVer = "# SIFVER #"
	)

	ICfg := cfg.NewCfg(cfgPath)
	cmn.FailOnCondition(ICfg == nil, "%v", fEf("SIF2JSON config couldn't be Loaded"))
	s2j := ICfg.(*cfg.SIF2JSON)

	cmn.FailOnCondition(sCount(s2j.SIFCfgDir4LIST, SignSIFVer) == 0, "SignSIFVer is missing @ %s, %v", cfgPath, fEf(""))
	cmn.FailOnCondition(sCount(s2j.SIFCfgDir4NUM, SignSIFVer) == 0, "SignSIFVer is missing @ %s, %v", cfgPath, fEf(""))
	cmn.FailOnCondition(sCount(s2j.SIFCfgDir4BOOL, SignSIFVer) == 0, "SignSIFVer is missing @ %s, %v", cfgPath, fEf(""))

	xmlReader := sNewReader(xml)
	jsonBuf, err := xj.Convert(
		xmlReader,
		// xj.WithTypeConverter(xj.Float, xj.Int, xj.Bool, xj.Null),
		// xj.WithAttrPrefix("-"),
		// xj.WithContentPrefix("#"),
	)
	cmn.FailOnErr("That's embarrassing... %v", err)

	json = pp.FmtJSONStr(jsonBuf.String(), s2j.JQDir)

	// Digital string to number
	// json := replaceDigCont(json, s2j.JQDir)

	// Attributes Modification
	obj := xmlroot(xml)              // infer object from xml root by default, use this object to search config json
	if enforced && len(subobj) > 0 { // if object is provided, ignore default, use 1st provided object to search
		obj = subobj[0]
	}

	if SIFVer != "" {
		s2j.DefaultSIFVer = SIFVer
	}

	// LIST
	s2j.SIFCfgDir4LIST = sReplaceAll(s2j.SIFCfgDir4LIST, SignSIFVer, s2j.DefaultSIFVer)
	LISTRules := getEachFileContent(s2j.SIFCfgDir4LIST+obj, "json", cmn.Iter2Slc(10)...)
	json = enforceConfig(json, s2j.JQDir, LISTRules...)

	// NUMERIC
	s2j.SIFCfgDir4NUM = sReplaceAll(s2j.SIFCfgDir4NUM, SignSIFVer, s2j.DefaultSIFVer)
	NUMRules := getEachFileContent(s2j.SIFCfgDir4NUM+obj, "json", cmn.Iter2Slc(2)...)
	json = enforceConfig(json, s2j.JQDir, NUMRules...)

	// BOOLEAN
	s2j.SIFCfgDir4BOOL = sReplaceAll(s2j.SIFCfgDir4BOOL, SignSIFVer, s2j.DefaultSIFVer)
	BOOLRules := getEachFileContent(s2j.SIFCfgDir4BOOL+obj, "json", cmn.Iter2Slc(2)...)
	json = enforceConfig(json, s2j.JQDir, BOOLRules...)

	return
}
