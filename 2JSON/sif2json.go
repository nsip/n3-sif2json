package cvt2json

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	xj "github.com/basgys/goxml2json"
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
		failOnErr("%v", err)
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
		// jsoncfg = fmtJSON(jsoncfg, 2)

		json, _ = newJKV(json, "", false).Unfold(0, newJKV(jsoncfg, "", false))
		// make sure there is no double "[" OR "]"
		bytes := rRB.ReplaceAll(rLB.ReplaceAll([]byte(json), []byte("[")), []byte("]"))
		json = fmtJSON(string(bytes), 2)
	}
	return json
}

// SIF2JSON : if [sifver] is "", DefaultSIFVer applies
func SIF2JSON(xml, sifver string, enforced bool, subobj ...string) (string, string, error) {

	jsonBuf, err := xj.Convert(
		sNewReader(xml),
		// xj.WithTypeConverter(xj.Float, xj.Int, xj.Bool, xj.Null),
		// xj.WithAttrPrefix("-"),
		// xj.WithContentPrefix("#"),
	)
	failOnErr("That's embarrassing... %v", err)

	// json, sv := jsonBuf.String(), ""
	// return // --------------------------- test 3rd party lib --------------------------- //

	json, sv := fmtJSON(jsonBuf.String(), 2), ""

	// Deal with 'LF', 'TB', Part1 -------------------------------------------------------- //
	mRepl1 := map[string]string{"\n": "#LF#", "\t": "#TB#"}
	for k, v := range mRepl1 {
		posGrp, values := [][]int{}, []string{}
		re := regexp.MustCompile(fSf(`": "[^"]*[%s]+[^"]*"[,\n]{1}`, k))
		for _, pos := range re.FindAllStringIndex(json, -1) {
			start, end := pos[0]+4, pos[1]-2
			posGrp = append(posGrp, []int{start, end})
			values = append(values, sReplaceAll(json[start:end], k, v))
		}
		json = replByPosGrp(json, posGrp, values)
	}

	// Attributes Modification according to Config ---------------------------------------- //
	obj := xmlRoot(xml)              // infer object from xml root by default, use this object to search config json
	if enforced && len(subobj) > 0 { // if object is provided, ignore default, use 1st provided object to search
		obj = subobj[0]
	}

	ver, dft := DftSIFVer, "Default "
	if sifver != "" {
		ver, dft = sifver, ""
	}

	// Convert to real path
	old := "#V#"
	Dir2SIFLIST = sReplaceAll(Dir2SIFLIST, old, ver)
	Dir2SIFNUM = sReplaceAll(Dir2SIFNUM, old, ver)
	Dir2SIFBOOL = sReplaceAll(Dir2SIFBOOL, old, ver)

	// Check SIFCfg Version Directory
	svDir := rmTailFromLastN(Dir2SIFLIST, "/", 2)
	if _, err := os.Stat(svDir); err == nil {
		sv = ver
	} else {
		// failOnErr("%v", fmt.Errorf("No %sSIF Spec @Version %s", dft, ver))
		return "", "", fmt.Errorf("No %sSIF Spec @Version %s", dft, ver)
	}

	/////////////////////////////
	// "../SIFSpec/3.4.7/json/LIST/" + "Activity"
	/////////////////////////////

	// LIST
	rules := eachFileContent(Dir2SIFLIST+obj, "json", iter2Slc(10)...)
	json = enforceConfig(json, rules...)

	// NUMERIC
	rules = eachFileContent(Dir2SIFNUM+obj, "json", iter2Slc(2)...)
	json = enforceConfig(json, rules...)

	// BOOLEAN
	rules = eachFileContent(Dir2SIFBOOL+obj, "json", iter2Slc(2)...)
	json = enforceConfig(json, rules...)

	// Deal with 'LF', 'TB'  Part2 -------------------------------------------------------------
	mRepl2 := map[string]string{"#LF#": "\\n", "#TB#": "\\t"}
	for k, v := range mRepl2 {
		json = sReplaceAll(json, k, v)
	}

	// XML empty element(empty text) with Attributes -------------------------------------------
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
	json = replByPosGrp(json, emptyPosPair, []string{fSf("\"%s\": \"\",\n", mark)})
	json = fmtJSON(json, 2)
	return json, sv, nil
}
