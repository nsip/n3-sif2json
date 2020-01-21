package main

import (
	"fmt"
	"io/ioutil"

	cmn "github.com/cdutwhu/json-util/common"
)

// Println :
func Println(num bool, slc ...string) {
	if num {
		for i, v := range slc {
			fmt.Printf("%d: %v\n", i, v)
		}
	} else {
		for _, v := range slc {
			fmt.Println(v)
		}
	}
}

// ObjGrp :
func ObjGrp(sep string, listGrp ...string) []string {
	m := map[string]bool{}
	for _, lsPath := range listGrp {
		obj := sSplit(lsPath, sep)[0]
		if _, ok := m[obj]; !ok {
			m[obj] = true
		}
	}
	return cmn.MapKeys(m).([]string)
}

// MapOfGrp :
func MapOfGrp(objs []string, sep string, listPathGrp ...string) map[string][]string {
	m := make(map[string][]string)
	for _, obj := range objs {
		prefix := obj + sep
		for _, lp := range listPathGrp {
			if sHasPrefix(lp, prefix) {
				m[obj] = append(m[obj], cmn.RmHeadToFirst(lp, sep))
			}
		}
	}
	return m
}

// PrintGrp4Cfg :
func PrintGrp4Cfg(m map[string][]string, attr string) (toml, goStruct string) {
	for obj, grp := range m {
		content := fmt.Sprintf("[%s]\n  %s = [", obj, attr)
		for _, path := range grp {
			content += fmt.Sprintf("\"%s\", ", path)
		}
		content = content[:len(content)-2] + "]"
		toml += content + "\n\n"

		// ------------------------- //
		content = fmt.Sprintf("%s struct { %s []string }", obj, attr)
		goStruct += content + "\n\t"
	}
	return
}

func main() {
	GenTomlAndStruct(
		"./out.txt",
		"0.0.1",
		"../2JSON/config/config.go.base",
		"../2JSON/config/List2JSON.toml.base",
		"../2JSON/config/Num2JSON.toml.base",
		"../2JSON/config/Bool2JSON.toml.base",
	)
}

// GenTomlAndStruct :
func GenTomlAndStruct(SIFSpecPath, SIFVer, basePath4GO, basePath4LIST, basePath4NUM, basePath4BOOL string) {

	// appears in ./2JSON/ .base files
	const (
		SignTOML    = "# AUTO-GEN #"
		SignGO4LIST = "// # AUTO-GEN: LIST # //"
		SignGO4NUM  = "// # AUTO-GEN: NUMERIC # //"
		SignGO4BOOL = "// # AUTO-GEN: BOOLEAN # //"
		SignSIFVer  = "# SIFVER #"
	)

	// Check [base] file Replace Marks //

	bytes, err := ioutil.ReadFile(basePath4GO)
	cmn.FailOnErr("%v", err)
	goStruct := string(bytes)
	cmn.FailOnCondition(sCount(goStruct, SignGO4LIST+"\n") != 1, "%v", fEf("@config.go.base SignGO4LIST"))
	cmn.FailOnCondition(sCount(goStruct, SignGO4NUM+"\n") != 1, "%v", fEf("@config.go.base SignGO4NUM"))
	cmn.FailOnCondition(sCount(goStruct, SignGO4BOOL+"\n") != 1, "%v", fEf("@config.go.base SignGO4BOOL"))

	bytes, err = ioutil.ReadFile(basePath4LIST)
	cmn.FailOnErr("%v", err)
	tomlLIST := string(bytes)
	cmn.FailOnCondition(sCount(tomlLIST, SignTOML) != 1, "%v", fEf("@list2json.toml.base SignTOML"))
	cmn.FailOnCondition(sCount(tomlLIST, SignSIFVer) != 1, "%v", fEf("@list2json.toml.base SignSIFVer"))

	bytes, err = ioutil.ReadFile(basePath4NUM)
	cmn.FailOnErr("%v", err)
	tomlNUM := string(bytes)
	cmn.FailOnCondition(sCount(tomlNUM, SignTOML) != 1, "%v", fEf("@num2json.toml.base SignTOML"))
	cmn.FailOnCondition(sCount(tomlNUM, SignSIFVer) != 1, "%v", fEf("@num2json.toml.base SignSIFVer"))

	bytes, err = ioutil.ReadFile(basePath4BOOL)
	cmn.FailOnErr("%v", err)
	tomlBOOL := string(bytes)
	cmn.FailOnCondition(sCount(tomlBOOL, SignTOML) != 1, "%v", fEf("@bool2json.toml.base SignTOML"))
	cmn.FailOnCondition(sCount(tomlBOOL, SignSIFVer) != 1, "%v", fEf("@bool2json.toml.base SignSIFVer"))

	// ************************************** //

	const (
		SEP     = "/"
		OBJECT  = "OBJECT: "
		LIST    = "LIST: "
		NUMERIC = "NUMERIC: "
		BOOLEAN = "BOOLEAN: "
	)

	var (
		objGrp      []string
		listPathGrp []string
		numPathGrp  []string
		boolPathGrp []string
	)

	bytes, err = ioutil.ReadFile(SIFSpecPath)
	cmn.FailOnErr("%v", err)
	content := string(bytes)

	for _, line := range sSplit(content, "\n") {
		switch {
		case sHasPrefix(line, OBJECT):
			objGrp = append(objGrp, line[len(OBJECT):])
		case sHasPrefix(line, LIST):
			listPathGrp = append(listPathGrp, cmn.RmTailFromLast(line[len(LIST):], "/"))
		case sHasPrefix(line, NUMERIC):
			numPathGrp = append(numPathGrp, line[len(NUMERIC):])
		case sHasPrefix(line, BOOLEAN):
			boolPathGrp = append(boolPathGrp, line[len(BOOLEAN):])
		}
	}

	// Println(true, objGrp...)
	// fmt.Println("-----------------------------")

	// Println(false, listPathGrp...)
	// fmt.Println("-----------------------------")

	{
		mListAttr := MapOfGrp(ObjGrp(SEP, listPathGrp...), SEP, listPathGrp...)
		mNumAttr := MapOfGrp(ObjGrp(SEP, numPathGrp...), SEP, numPathGrp...)
		mBoolAttr := MapOfGrp(ObjGrp(SEP, boolPathGrp...), SEP, boolPathGrp...)

		toml4List, goStruct4List := PrintGrp4Cfg(mListAttr, "LIST")
		toml4Num, goStruct4Num := PrintGrp4Cfg(mNumAttr, "NUMERIC")
		toml4Bool, goStruct4Bool := PrintGrp4Cfg(mBoolAttr, "BOOLEAN")

		toml := sReplace(tomlLIST, SignSIFVer, SIFVer, 1)
		toml = sReplace(toml, SignTOML, toml4List, 1)
		tomlOutPath := basePath4LIST[:len(basePath4LIST)-5] // remove ".base" ext-name
		cmn.FailOnErr("%v", ioutil.WriteFile(tomlOutPath, []byte(toml), 0666))

		toml = sReplace(tomlNUM, SignSIFVer, SIFVer, 1)
		toml = sReplace(toml, SignTOML, toml4Num, 1)
		tomlOutPath = basePath4NUM[:len(basePath4NUM)-5] // remove ".base" ext-name
		cmn.FailOnErr("%v", ioutil.WriteFile(tomlOutPath, []byte(toml), 0666))

		toml = sReplace(tomlBOOL, SignSIFVer, SIFVer, 1)
		toml = sReplace(toml, SignTOML, toml4Bool, 1)
		tomlOutPath = basePath4BOOL[:len(basePath4BOOL)-5] // remove ".base" ext-name
		cmn.FailOnErr("%v", ioutil.WriteFile(tomlOutPath, []byte(toml), 0666))

		goStruct = sReplace(goStruct, SignGO4LIST, goStruct4List, 1)
		goStruct = sReplace(goStruct, SignGO4NUM, goStruct4Num, 1)
		goStruct = sReplace(goStruct, SignGO4BOOL, goStruct4Bool, 1)
		goOutPath := basePath4GO[:len(basePath4GO)-5] // for below INPUT
		cmn.FailOnErr("%v", ioutil.WriteFile(goOutPath, []byte(goStruct), 0666))
	}
}
