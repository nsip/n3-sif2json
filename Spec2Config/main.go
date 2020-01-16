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
		goStruct += content + "\n"
	}
	return
}

func main() {
	GenTomlStruct(
		"./out.txt",
		"../2JSON/config.go.base",
		"../2JSON/config/list2json.toml.base",
		"../2JSON/config/num2json.toml.base",
		"../2JSON/config/bool2json.toml.base",
	)
}

// GenTomlStruct :
func GenTomlStruct(specTxtPath, goStructBasePath, tomlBasePath4LIST, tomlBasePath4NUM, tomlBasePath4BOOL string) {

	// appears in ./2JSON/ .base files
	const (
		ReplSignTOML    = "# AUTO-GEN #"
		ReplSignGO4LIST = "// # AUTO-GEN: LIST # //"
		ReplSignGO4NUM  = "// # AUTO-GEN: NUMERIC # //"
		ReplSignGO4BOOL = "// # AUTO-GEN: BOOLEAN # //"
	)

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

	bytes, err := ioutil.ReadFile(specTxtPath)
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
		toml4List, goStruct4List := PrintGrp4Cfg(mListAttr, "LIST")

		tb, err := ioutil.ReadFile(tomlBasePath4LIST)
		cmn.FailOnErr("%v", err)
		tomlOutPath := tomlBasePath4LIST[:len(tomlBasePath4LIST)-5]
		toml := sReplace(string(tb), ReplSignTOML, toml4List, 1)
		ioutil.WriteFile(tomlOutPath, []byte(toml), 0666)

		st, err := ioutil.ReadFile(goStructBasePath)
		cmn.FailOnErr("%v", err)
		goOutPath := goStructBasePath[:len(goStructBasePath)-5] // for below input
		goStruct4List = sReplace(string(st), ReplSignGO4LIST, goStruct4List, 1)
		ioutil.WriteFile(goOutPath, []byte(goStruct4List), 0666)

		// ------------------------------------ //

		mNumAttr := MapOfGrp(ObjGrp(SEP, numPathGrp...), SEP, numPathGrp...)
		toml4Num, goStruct4Num := PrintGrp4Cfg(mNumAttr, "NUMERIC")

		tb, err = ioutil.ReadFile(tomlBasePath4NUM)
		cmn.FailOnErr("%v", err)
		tomlOutPath = tomlBasePath4NUM[:len(tomlBasePath4NUM)-5]
		toml = sReplace(string(tb), ReplSignTOML, toml4Num, 1)
		ioutil.WriteFile(tomlOutPath, []byte(toml), 0666)

		st, err = ioutil.ReadFile(goOutPath) // read from previous go out path
		cmn.FailOnErr("%v", err)
		goStruct4Num = sReplace(string(st), ReplSignGO4NUM, goStruct4Num, 1)
		ioutil.WriteFile(goOutPath, []byte(goStruct4Num), 0666)

		// ------------------------------------ //

		mBoolAttr := MapOfGrp(ObjGrp(SEP, boolPathGrp...), SEP, boolPathGrp...)
		toml4Bool, goStruct4Bool := PrintGrp4Cfg(mBoolAttr, "BOOLEAN")

		tb, err = ioutil.ReadFile(tomlBasePath4BOOL)
		cmn.FailOnErr("%v", err)
		tomlOutPath = tomlBasePath4BOOL[:len(tomlBasePath4BOOL)-5]
		toml = sReplace(string(tb), ReplSignTOML, toml4Bool, 1)
		ioutil.WriteFile(tomlOutPath, []byte(toml), 0666)

		st, err = ioutil.ReadFile(goOutPath) // read from previous go out path
		cmn.FailOnErr("%v", err)
		goStruct4Bool = sReplace(string(st), ReplSignGO4BOOL, goStruct4Bool, 1)
		ioutil.WriteFile(goOutPath, []byte(goStruct4Bool), 0666)

	}
}
