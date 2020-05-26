package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	eg "github.com/cdutwhu/n3-util/n3errs"
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
	return mapKeys(m).([]string)
}

// MapOfGrp :
func MapOfGrp(objs []string, sep string, xxxPathGrp ...string) map[string][]string {
	m := make(map[string][]string)
	for _, obj := range objs {
		prefix := obj + sep
		for _, lp := range xxxPathGrp {
			if sHasPrefix(lp, prefix) {
				m[obj] = append(m[obj], rmHeadToFirst(lp, sep))
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
	GenTomlAndGoSrc(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6], os.Args[7])
	abs, err := filepath.Abs(os.Args[7])
	failOnErr("%v", err)
	fPf("Dumped [spec.go] [Bool2JSON.toml] [List2JSON.toml] [Num2JSON.toml] into %s\n", abs)
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
	failOnErr("%v", err)
	goStruct := string(bytes)
	failOnErrWhen(sCount(goStruct, SignGO4LIST+"\n") != 1, "%v: goStruct SignGO4LIST", eg.SRC_SIGN_MISSING)
	failOnErrWhen(sCount(goStruct, SignGO4NUM+"\n") != 1, "%v: goStruct SignGO4NUM", eg.SRC_SIGN_MISSING)
	failOnErrWhen(sCount(goStruct, SignGO4BOOL+"\n") != 1, "%v: goStruct SignGO4BOOL", eg.SRC_SIGN_MISSING)

	bytes, err = ioutil.ReadFile(baseToml4LIST)
	failOnErr("%v", err)
	tomlLIST := string(bytes)
	failOnErrWhen(sCount(tomlLIST, SignTOML) != 1, "%v: tomlLIST SignTOML", eg.CFG_SIGN_MISSING)
	failOnErrWhen(sCount(tomlLIST, SignSIFVer) != 1, "%v: tomlLIST SignSIFVer", eg.CFG_SIGN_MISSING)

	bytes, err = ioutil.ReadFile(baseToml4NUM)
	failOnErr("%v", err)
	tomlNUM := string(bytes)
	failOnErrWhen(sCount(tomlNUM, SignTOML) != 1, "%v: tomlNUM SignTOML", eg.CFG_SIGN_MISSING)
	failOnErrWhen(sCount(tomlNUM, SignSIFVer) != 1, "%v: tomlNUM SignSIFVer", eg.CFG_SIGN_MISSING)

	bytes, err = ioutil.ReadFile(baseToml4BOOL)
	failOnErr("%v", err)
	tomlBOOL := string(bytes)
	failOnErrWhen(sCount(tomlBOOL, SignTOML) != 1, "%v: tomlBOOL SignTOML", eg.CFG_SIGN_MISSING)
	failOnErrWhen(sCount(tomlBOOL, SignSIFVer) != 1, "%v: tomlBOOL SignSIFVer", eg.CFG_SIGN_MISSING)

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
	failOnErr("%v", err)
	content := string(bytes)

	SIFVer := ""

	for _, line := range sSplit(content, "\n") {
		switch {
		case sHasPrefix(line, VERSION):
			SIFVer = sTrim(line[len(VERSION):], " \t\r\n")
		case sHasPrefix(line, OBJECT):
			objGrp = append(objGrp, sTrim(line[len(OBJECT):], " \t\r\n"))
		case sHasPrefix(line, LIST):
			// listPathGrp = append(listPathGrp, rmTailFromLast(line[len(LIST):], "/")) // exclude last one
			listPathGrp = append(listPathGrp, sTrim(line[len(LIST):], " \t\r\n"))
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

		toml := sReplace(tomlLIST, SignSIFVer, SIFVer, 1)
		toml = sReplace(toml, SignTOML, toml4List, 1)
		baseFile4LIST := rmHeadToLast(baseToml4LIST, "/") + ".toml"
		failOnErr("%v", ioutil.WriteFile(outDir+baseFile4LIST, []byte(toml), 0666))

		toml = sReplace(tomlNUM, SignSIFVer, SIFVer, 1)
		toml = sReplace(toml, SignTOML, toml4Num, 1)
		baseFile4NUM := rmHeadToLast(baseToml4NUM, "/") + ".toml"
		failOnErr("%v", ioutil.WriteFile(outDir+baseFile4NUM, []byte(toml), 0666))

		toml = sReplace(tomlBOOL, SignSIFVer, SIFVer, 1)
		toml = sReplace(toml, SignTOML, toml4Bool, 1)
		baseFile4BOOL := rmHeadToLast(baseToml4BOOL, "/") + ".toml"
		failOnErr("%v", ioutil.WriteFile(outDir+baseFile4BOOL, []byte(toml), 0666))

		goStruct = sReplace(goStruct, SignGO4LIST, goStruct4List, 1)
		goStruct = sReplace(goStruct, SignGO4NUM, goStruct4Num, 1)
		goStruct = sReplace(goStruct, SignGO4BOOL, goStruct4Bool, 1)
		baseFile4GO := rmHeadToLast(baseGO, "/") + ".go"
		failOnErr("%v", ioutil.WriteFile(outDir+baseFile4GO, []byte(goStruct), 0666))
	}
}
