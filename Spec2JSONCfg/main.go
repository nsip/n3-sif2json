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
		"../2JSON/config/",
		"../2JSON/config/base-go/config",
		"../2JSON/config/base-toml/List2JSON",
		"../2JSON/config/base-toml/Num2JSON",
		"../2JSON/config/base-toml/Bool2JSON",
	)
}

// GenTomlAndStruct :
func GenTomlAndStruct(SIFSpecPath, SIFVer, outDir, basePath4GO, basePath4LIST, basePath4NUM, basePath4BOOL string) {

	// appears in ./2JSON/ .base files
	const (
		SignTOML    = "# AUTO-GEN #"
		SignGO4LIST = "// # AUTO-GEN: LIST # //"
		SignGO4NUM  = "// # AUTO-GEN: NUMERIC # //"
		SignGO4BOOL = "// # AUTO-GEN: BOOLEAN # //"
		SignSIFVer  = "# SIF-VER #"
	)

	// Check [base] file Replace Marks //

	bytes, err := ioutil.ReadFile(basePath4GO)
	cmn.FailOnErr("%v", err)
	goStruct := string(bytes)
	cmn.FailOnErrWhen(sCount(goStruct, SignGO4LIST+"\n") != 1, "%v", fEf("@config.go.base SignGO4LIST"))
	cmn.FailOnErrWhen(sCount(goStruct, SignGO4NUM+"\n") != 1, "%v", fEf("@config.go.base SignGO4NUM"))
	cmn.FailOnErrWhen(sCount(goStruct, SignGO4BOOL+"\n") != 1, "%v", fEf("@config.go.base SignGO4BOOL"))

	bytes, err = ioutil.ReadFile(basePath4LIST)
	cmn.FailOnErr("%v", err)
	tomlLIST := string(bytes)
	cmn.FailOnErrWhen(sCount(tomlLIST, SignTOML) != 1, "%v", fEf("@list2json.toml.base SignTOML"))
	cmn.FailOnErrWhen(sCount(tomlLIST, SignSIFVer) != 1, "%v", fEf("@list2json.toml.base SignSIFVer"))

	bytes, err = ioutil.ReadFile(basePath4NUM)
	cmn.FailOnErr("%v", err)
	tomlNUM := string(bytes)
	cmn.FailOnErrWhen(sCount(tomlNUM, SignTOML) != 1, "%v", fEf("@num2json.toml.base SignTOML"))
	cmn.FailOnErrWhen(sCount(tomlNUM, SignSIFVer) != 1, "%v", fEf("@num2json.toml.base SignSIFVer"))

	bytes, err = ioutil.ReadFile(basePath4BOOL)
	cmn.FailOnErr("%v", err)
	tomlBOOL := string(bytes)
	cmn.FailOnErrWhen(sCount(tomlBOOL, SignTOML) != 1, "%v", fEf("@bool2json.toml.base SignTOML"))
	cmn.FailOnErrWhen(sCount(tomlBOOL, SignSIFVer) != 1, "%v", fEf("@bool2json.toml.base SignSIFVer"))

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
		baseFile4LIST := cmn.RmHeadToLast(basePath4LIST, "/") + ".toml"
		cmn.FailOnErr("%v", ioutil.WriteFile(outDir+baseFile4LIST, []byte(toml), 0666))

		toml = sReplace(tomlNUM, SignSIFVer, SIFVer, 1)
		toml = sReplace(toml, SignTOML, toml4Num, 1)
		baseFile4NUM := cmn.RmHeadToLast(basePath4NUM, "/") + ".toml"
		cmn.FailOnErr("%v", ioutil.WriteFile(outDir+baseFile4NUM, []byte(toml), 0666))

		toml = sReplace(tomlBOOL, SignSIFVer, SIFVer, 1)
		toml = sReplace(toml, SignTOML, toml4Bool, 1)
		baseFile4BOOL := cmn.RmHeadToLast(basePath4BOOL, "/") + ".toml"
		cmn.FailOnErr("%v", ioutil.WriteFile(outDir+baseFile4BOOL, []byte(toml), 0666))

		goStruct = sReplace(goStruct, SignGO4LIST, goStruct4List, 1)
		goStruct = sReplace(goStruct, SignGO4NUM, goStruct4Num, 1)
		goStruct = sReplace(goStruct, SignGO4BOOL, goStruct4Bool, 1)
		baseFile4GO := cmn.RmHeadToLast(basePath4GO, "/") + ".go"
		cmn.FailOnErr("%v", ioutil.WriteFile(outDir+baseFile4GO, []byte(goStruct), 0666))
	}
}
