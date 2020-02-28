package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
func MapOfGrp(objs []string, sep string, xxxPathGrp ...string) map[string][]string {
	m := make(map[string][]string)
	for _, obj := range objs {
		prefix := obj + sep
		for _, lp := range xxxPathGrp {
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
	if len(os.Args) < 7 {
		fPln("You are not allowed to use this cli to create next build step resource unless fully understand what you are doing.\n" +
			"Project author or other admins are advised to do this for adding SIF Specifications.\n" +
			"If you still want to continue, input following arguments orderly:\n" +
			"  1. SIF Spec. file path. (a copy is /SIFSpec/out.txt                                        DO NOT edit!)\n" +
			"  2. path of config go-source base, (a copy exists in /2JSON/SpecCfgMaker/base-go/config     DO NOT edit!)\n" +
			"  3. path of List2JSON toml base, (a copy exists in /2JSON/SpecCfgMaker/base-toml/List2JSON  DO NOT edit!)\n" +
			"  4. path of Num2JSON toml base, (a copy exists in /2JSON/SpecCfgMaker/base-toml/Num2JSON    DO NOT edit!)\n" +
			"  5. path of Bool2JSON toml base, (a copy exists in /2JSON/SpecCfgMaker/base-toml/Bool2JSON  DO NOT edit!)\n" +
			"  6. auto-created go & toml configuration files output directory")
		return
	}
	SIFSpecName := os.Args[1]
	goBaseFile := os.Args[2]
	listTomlBaseFile := os.Args[3]
	numTomlBaseFile := os.Args[4]
	boolTomlBaseFile := os.Args[5]
	cfgOutputDir := os.Args[6]
	GenTomlAndGoSrc(SIFSpecName, goBaseFile, listTomlBaseFile, numTomlBaseFile, boolTomlBaseFile, cfgOutputDir)
	abs, err := filepath.Abs(cfgOutputDir)
	cmn.FailOnErr("%v", err)
	fPf("Dumped [config.go] [Bool2JSON.toml] [List2JSON.toml] [Num2JSON.toml] into %s\n", abs)
}

// GenTomlAndGoSrc :
func GenTomlAndGoSrc(SIFSpecPath, baseGO, baseToml4LIST, baseToml4NUM, baseToml4BOOL, outDir string) {

	// appears in ./2JSON/ .base files
	const (
		SignTOML    = "#AUTO-GEN#"
		SignGO4LIST = "// #AUTO-GEN: LIST# //"
		SignGO4NUM  = "// #AUTO-GEN: NUMERIC# //"
		SignGO4BOOL = "// #AUTO-GEN: BOOLEAN# //"
		SignSIFVer  = "#SIF-VER#"
	)

	// Check [base] file Replace Marks //

	bytes, err := ioutil.ReadFile(baseGO)
	cmn.FailOnErr("%v", err)
	goStruct := string(bytes)
	cmn.FailOnErrWhen(sCount(goStruct, SignGO4LIST+"\n") != 1, "%v", fEf("@SignGO4LIST"))
	cmn.FailOnErrWhen(sCount(goStruct, SignGO4NUM+"\n") != 1, "%v", fEf("@SignGO4NUM"))
	cmn.FailOnErrWhen(sCount(goStruct, SignGO4BOOL+"\n") != 1, "%v", fEf("@SignGO4BOOL"))

	bytes, err = ioutil.ReadFile(baseToml4LIST)
	cmn.FailOnErr("%v", err)
	tomlLIST := string(bytes)
	cmn.FailOnErrWhen(sCount(tomlLIST, SignTOML) != 1, "%v", fEf("@SignTOML"))
	cmn.FailOnErrWhen(sCount(tomlLIST, SignSIFVer) != 1, "%v", fEf("@SignSIFVer"))

	bytes, err = ioutil.ReadFile(baseToml4NUM)
	cmn.FailOnErr("%v", err)
	tomlNUM := string(bytes)
	cmn.FailOnErrWhen(sCount(tomlNUM, SignTOML) != 1, "%v", fEf("@SignTOML"))
	cmn.FailOnErrWhen(sCount(tomlNUM, SignSIFVer) != 1, "%v", fEf("@SignSIFVer"))

	bytes, err = ioutil.ReadFile(baseToml4BOOL)
	cmn.FailOnErr("%v", err)
	tomlBOOL := string(bytes)
	cmn.FailOnErrWhen(sCount(tomlBOOL, SignTOML) != 1, "%v", fEf("@SignTOML"))
	cmn.FailOnErrWhen(sCount(tomlBOOL, SignSIFVer) != 1, "%v", fEf("@SignSIFVer"))

	// ************************************** //

	const (
		SEP     = "/"
		VERSION = "VERSION: "
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

	SIFVer := ""

	for _, line := range sSplit(content, "\n") {
		switch {
		case sHasPrefix(line, VERSION):
			SIFVer = sTrim(line[len(VERSION):], " \t\r\n")
		case sHasPrefix(line, OBJECT):
			objGrp = append(objGrp, sTrim(line[len(OBJECT):], " \t\r\n"))
		case sHasPrefix(line, LIST):
			listPathGrp = append(listPathGrp, cmn.RmTailFromLast(line[len(LIST):], "/"))
		case sHasPrefix(line, NUMERIC):
			numPathGrp = append(numPathGrp, sTrim(line[len(NUMERIC):], " \t\r\n"))
		case sHasPrefix(line, BOOLEAN):
			boolPathGrp = append(boolPathGrp, sTrim(line[len(BOOLEAN):], " \t\r\n"))
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

		// fPln(SIFVer)
		// cmn.FailOnErrWhen(SIFVer != "3.4.5X", "%v", fEf("why?"))

		toml := sReplace(tomlLIST, SignSIFVer, SIFVer, 1)
		toml = sReplace(toml, SignTOML, toml4List, 1)
		baseFile4LIST := cmn.RmHeadToLast(baseToml4LIST, "/") + ".toml"
		cmn.FailOnErr("%v", ioutil.WriteFile(outDir+baseFile4LIST, []byte(toml), 0666))

		toml = sReplace(tomlNUM, SignSIFVer, SIFVer, 1)
		toml = sReplace(toml, SignTOML, toml4Num, 1)
		baseFile4NUM := cmn.RmHeadToLast(baseToml4NUM, "/") + ".toml"
		cmn.FailOnErr("%v", ioutil.WriteFile(outDir+baseFile4NUM, []byte(toml), 0666))

		toml = sReplace(tomlBOOL, SignSIFVer, SIFVer, 1)
		toml = sReplace(toml, SignTOML, toml4Bool, 1)
		baseFile4BOOL := cmn.RmHeadToLast(baseToml4BOOL, "/") + ".toml"
		cmn.FailOnErr("%v", ioutil.WriteFile(outDir+baseFile4BOOL, []byte(toml), 0666))

		goStruct = sReplace(goStruct, SignGO4LIST, goStruct4List, 1)
		goStruct = sReplace(goStruct, SignGO4NUM, goStruct4Num, 1)
		goStruct = sReplace(goStruct, SignGO4BOOL, goStruct4Bool, 1)
		baseFile4GO := cmn.RmHeadToLast(baseGO, "/") + ".go"
		cmn.FailOnErr("%v", ioutil.WriteFile(outDir+baseFile4GO, []byte(goStruct), 0666))
	}
}
