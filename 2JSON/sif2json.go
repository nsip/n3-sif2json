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
func enforceConfig(json string, lsJSONCfg ...string) string {

	rLB := regexp.MustCompile(`\[[ \t\r\n]*\[`)
	rRB := regexp.MustCompile(`\][ \t\r\n]*\]`)

	for _, jsoncfg := range lsJSONCfg {
		// make sure [jsoncfg] is formatted; Otherwise, do Fmt firstly
		// jsoncfg = jkv.FmtJSON(jsoncfg, 2)

		json, _ = jkv.NewJKV(json, "", false).Unfold(0, jkv.NewJKV(jsoncfg, "", false))
		// make sure there is no double "[" OR "]"
		bytes := rRB.ReplaceAll(rLB.ReplaceAll([]byte(json), []byte("[")), []byte("]"))
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

	cmn.FailOnErrWhen(sCount(s2j.SIFCfgDir4LIST, SignSIFVer) == 0, "Missing SignSIFVer @ %s, %v", cfgPath, fEf(""))
	cmn.FailOnErrWhen(sCount(s2j.SIFCfgDir4NUM, SignSIFVer) == 0, "Missing SignSIFVer @ %s, %v", cfgPath, fEf(""))
	cmn.FailOnErrWhen(sCount(s2j.SIFCfgDir4BOOL, SignSIFVer) == 0, "Missing SignSIFVer @ %s, %v", cfgPath, fEf(""))

	xmlReader := sNewReader(xml)
	jsonBuf, err := xj.Convert(
		xmlReader,
		// xj.WithTypeConverter(xj.Float, xj.Int, xj.Bool, xj.Null),
		// xj.WithAttrPrefix("-"),
		// xj.WithContentPrefix("#"),
	)
	cmn.FailOnErr("That's embarrassing... %v", err)

	// json = jsonBuf.String()
	// return // --------------------------- test 3rd party lib --------------------------- //

	json = jkv.FmtJSON(jsonBuf.String(), 2)

	// Deal with 'LF', 'TB', Part1 --------------------------------------------------------------------------
	mRepl1 := map[string]string{"\n": "#LF#", "\t": "#TB#"}
	for k, v := range mRepl1 {
		posGrp, values := [][]int{}, []string{}
		re := regexp.MustCompile(fSf(`": "[^"]*[%s]+[^"]*"[,\n]{1}`, k))
		for _, pos := range re.FindAllStringIndex(json, -1) {
			start, end := pos[0]+4, pos[1]-2
			posGrp = append(posGrp, []int{start, end})
			values = append(values, sReplaceAll(json[start:end], k, v))
		}
		json = cmn.ReplByPosGrp(json, posGrp, values)
	}

	// Attributes Modification according to Config ----------------------------------------------------------
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

	// LIST
	rules := eachFileContent(s2j.SIFCfgDir4LIST+obj, "json", cmn.Iter2Slc(10)...)
	json = enforceConfig(json, rules...)

	// NUMERIC
	rules = eachFileContent(s2j.SIFCfgDir4NUM+obj, "json", cmn.Iter2Slc(2)...)
	json = enforceConfig(json, rules...)

	// BOOLEAN
	rules = eachFileContent(s2j.SIFCfgDir4BOOL+obj, "json", cmn.Iter2Slc(2)...)
	json = enforceConfig(json, rules...)

	// Deal with 'LF', 'TB'  Part2 --------------------------------------------------------------------------
	mRepl2 := map[string]string{"#LF#": "\\n", "#TB#": "\\t"}
	for k, v := range mRepl2 {
		json = sReplaceAll(json, k, v)
	}

	// XML empty element(empty text) with Attributes --------------------------------------------------------
	emptyPosPair := [][]int{}

	re1 := regexp.MustCompile(`": \{\n([ ]+"-.+": .+,\n)*([ ]+"-.+": .+\n)[ ]+\}`) // one empty object
	for _, pos := range re1.FindAllStringIndex(json, -1) {
		emptyPosPair = append(emptyPosPair, []int{pos[0] + 6, pos[0] + 6})
	}

	re2 := regexp.MustCompile(`[\[,]\n[ ]+\{\n([ ]+"-.+": .+,\n)*([ ]+"-.+": .+\n)[ ]+\}`) // empty object in array
	for _, pos := range re2.FindAllStringIndex(json, -1) {
		remain, offset := json[pos[0]:], 0
		for i, c := range remain {
			if c == '{' {
				offset = i + 1
				break
			}
		}
		emptyPosPair = append(emptyPosPair, []int{pos[0] + offset, pos[0] + offset})
	}

	const mark = "value" // "#content"
	json = cmn.ReplByPosGrp(json, emptyPosPair, []string{fSf("\"%s\": \"\",\n", mark)})
	json = jkv.FmtJSON(json, 2)

	return
}
