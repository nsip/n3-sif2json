package cvt2json

import (
	"io/ioutil"
	"os"
	"regexp"

	xj "github.com/basgys/goxml2json"
	cmn "github.com/cdutwhu/json-util/common"
	jkv "github.com/cdutwhu/json-util/jkv"
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

func eachFileContent(dir, ext string, indices ...int) (rt []string) {
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

	rLB := regexp.MustCompile(`\[[ \t\r\n]*\[`)
	rRB := regexp.MustCompile(`\][ \t\r\n]*\]`)

	for _, jsoncfg := range lsJSONCfg {
		// make sure [jsoncfg] is formatted
		// otherwise, do Fmt firstly
		// jsoncfg = pp.FmtJSONStr(jsoncfg, jqDir)

		json, _ = jkv.NewJKV(json, "", false).Unfold(0, jkv.NewJKV(jsoncfg, "", false))
		// make sure there is no double "[" OR "]"
		bytes := rRB.ReplaceAll(rLB.ReplaceAll([]byte(json), []byte("[")), []byte("]"))
		// json = pp.FmtJSONStr(string(bytes), jqDir)
		json = jkv.FmtJSON(string(bytes), 2)
	}
	return json
}

// SIF2JSON : if [SIFVer] is "", use config's DefaultSIFVer
func SIF2JSON(cfgPath, xml, SIFVer string, enforced bool, subobj ...string) (json, sv string, err error) {
	const (
		SignSIFVer = "#SIFVER#"
	)

	ICfg := cfg.NewCfg(cfgPath)
	cmn.FailOnErrWhen(ICfg == nil, "%v", fEf("SIF2JSON config couldn't be Loaded"))
	s2j := ICfg.(*cfg.SIF2JSON)

	cmn.FailOnErrWhen(sCount(s2j.SIFCfgDir4LIST, SignSIFVer) == 0, "SignSIFVer is missing @ %s, %v", cfgPath, fEf(""))
	cmn.FailOnErrWhen(sCount(s2j.SIFCfgDir4NUM, SignSIFVer) == 0, "SignSIFVer is missing @ %s, %v", cfgPath, fEf(""))
	cmn.FailOnErrWhen(sCount(s2j.SIFCfgDir4BOOL, SignSIFVer) == 0, "SignSIFVer is missing @ %s, %v", cfgPath, fEf(""))

	xmlReader := sNewReader(xml)
	jsonBuf, err := xj.Convert(
		xmlReader,
		// xj.WithTypeConverter(xj.Float, xj.Int, xj.Bool, xj.Null),
		// xj.WithAttrPrefix("-"),
		// xj.WithContentPrefix("#"),
	)
	cmn.FailOnErr("That's embarrassing... %v", err)

	// json = jsonBuf.String()
	// return // --------------------------------------- test 3rd party lib --------------------------------------- //

	// json = pp.FmtJSONStr(jsonBuf.String(), s2j.JQDir)
	json = jkv.FmtJSON(jsonBuf.String(), 2)

	// Deal with 'LF', 'TB', P1
	posGrp, values := [][]int{}, []string{}
	for _, pos := range regexp.MustCompile(`": "[^"]*[\n]+[^"]*"[,\n]{1}`).FindAllStringIndex(json, -1) {
		start, end := pos[0]+4, pos[1]-2
		posGrp = append(posGrp, []int{start, end})
		values = append(values, sReplaceAll(json[start:end], "\n", "#LF#"))
	}
	json = cmn.ReplByPosGrp(json, posGrp, values)

	posGrp, values = [][]int{}, []string{}
	for _, pos := range regexp.MustCompile(`": "[^"]*[\t]+[^"]*"[,\n]{1}`).FindAllStringIndex(json, -1) {
		start, end := pos[0]+4, pos[1]-2
		posGrp = append(posGrp, []int{start, end})
		values = append(values, sReplaceAll(json[start:end], "\t", "#TB#"))
	}
	json = cmn.ReplByPosGrp(json, posGrp, values)

	// AGAIN1:
	// 	for _, pos := range regexp.MustCompile(`": "[^"]*[\n]+[^"]*"[,\n]{1}`).FindAllStringIndex(json, 1) {
	// 		start, end := pos[0]+4, pos[1]-2
	// 		value := json[start:end]
	// 		value = sReplaceAll(value, "\n", "#LF#")
	// 		json = sReplByPos(json, start, end, value)
	// 		goto AGAIN1
	// 	}

	// AGAIN2:
	// 	for _, pos := range regexp.MustCompile(`": "[^"]*[\t]+[^"]*"[,\n]{1}`).FindAllStringIndex(json, 1) {
	// 		start, end := pos[0]+4, pos[1]-2
	// 		value := json[start:end]
	// 		value = sReplaceAll(value, "\t", "#TB#")
	// 		json = sReplByPos(json, start, end, value)
	// 		goto AGAIN2
	// 	}
	// End Dealing with 'LF', 'TB', P1

	// return // --------------------------------------- test 3rd party lib --------------------------------------- //

	// Attributes Modification
	obj := cmn.XMLRoot(xml)          // infer object from xml root by default, use this object to search config json
	if enforced && len(subobj) > 0 { // if object is provided, ignore default, use 1st provided object to search
		obj = subobj[0]
	}

	dft := "Default "
	if SIFVer != "" {
		s2j.DefaultSIFVer = SIFVer
		dft = ""
	}

	// SIFCfgDir Version Set
	s2j.SIFCfgDir4LIST = sReplaceAll(s2j.SIFCfgDir4LIST, SignSIFVer, s2j.DefaultSIFVer)
	s2j.SIFCfgDir4NUM = sReplaceAll(s2j.SIFCfgDir4NUM, SignSIFVer, s2j.DefaultSIFVer)
	s2j.SIFCfgDir4BOOL = sReplaceAll(s2j.SIFCfgDir4BOOL, SignSIFVer, s2j.DefaultSIFVer)

	// Check SIFCfg Version Directory
	svDir := cmn.RmTailFromLastN(s2j.SIFCfgDir4LIST, "/", 2)
	if _, err := os.Stat(svDir); err == nil {
		sv = cmn.RmHeadToLast(svDir, "/")
	} else {
		return "", "", fEf("No %sSIF Spec @Version %s", dft, s2j.DefaultSIFVer)
	}
	// End Checking

	// LIST
	rules := eachFileContent(s2j.SIFCfgDir4LIST+obj, "json", cmn.Iter2Slc(10)...)
	json = enforceConfig(json, s2j.JQDir, rules...)

	// NUMERIC
	rules = eachFileContent(s2j.SIFCfgDir4NUM+obj, "json", cmn.Iter2Slc(2)...)
	json = enforceConfig(json, s2j.JQDir, rules...)

	// BOOLEAN
	rules = eachFileContent(s2j.SIFCfgDir4BOOL+obj, "json", cmn.Iter2Slc(2)...)
	json = enforceConfig(json, s2j.JQDir, rules...)

	// Deal with 'LF', 'TB'  P2
	json = sReplaceAll(json, "#LF#", "\\n")
	json = sReplaceAll(json, "#TB#", "\\t")
	// End Dealing with 'LF', 'TB', P2

	return
}
