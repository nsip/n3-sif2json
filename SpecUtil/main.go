package main

import (
	"fmt"
	"io/ioutil"
	"strings"

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

// ObjGrpWithList :
func ObjGrpWithList(sep string, listGrp ...string) []string {
	m := map[string]bool{}
	for _, lsPath := range listGrp {
		obj := strings.Split(lsPath, sep)[0]
		if _, ok := m[obj]; !ok {
			m[obj] = true
		}
	}
	return cmn.MapKeys(m).([]string)
}

// MapOfListPathGrp :
func MapOfListPathGrp(objs []string, sep string, listPathGrp ...string) map[string][]string {
	m := make(map[string][]string)
	for _, obj := range objs {
		prefix := obj + sep
		for _, lp := range listPathGrp {
			if strings.HasPrefix(lp, prefix) {
				m[obj] = append(m[obj], cmn.RmHeadToFirst(lp, sep))
			}
		}
	}
	return m
}

// PrintListPathGrp4Cfg :
func PrintListPathGrp4Cfg(m map[string][]string, attr string) (toml, gostruct string) {
	for obj, listPathGrp := range m {
		content := fmt.Sprintf("[%s]\n  %s = [", obj, attr)
		for _, path := range listPathGrp {
			content += fmt.Sprintf("\"%s\", ", path)
		}
		content = content[:len(content)-2] + "]"
		toml += content + "\n\n"

		// ------------------------- //
		content = fmt.Sprintf("%s struct { %s []string }", obj, attr)
		gostruct += content + "\n"
	}
	return
}

func main() {
	GenTomlStruct(
		"./out.txt",
		"../2JSON/config/cfg2json.toml.base",
		"../2JSON/config/cfg2json.toml",
		"../2JSON/config.go.base",
		"../2JSON/config.go",
	)
}

// GenTomlStruct :
func GenTomlStruct(specTxtPath, tomlTemplatePath, tomlOutPath, structTemplatePath, structGoPath string) {

	const (
		ReplSignTOML = "# AUTO-GEN #"
		ReplSignGO   = "// AUTO-GEN //"
	)

	const (
		SEP       = "/"
		OBJECT    = "OBJECT: "
		LIST      = "LIST: "
		NUMERIC   = "NUMERIC: "
		BOOLEAN   = "BOOLEAN: "
		XPATHTYPE = "XPATHTYPE: "
	)

	var (
		objGrp      []string
		listPathGrp []string
	)

	bytes, err := ioutil.ReadFile(specTxtPath)
	cmn.FailOnErr("%v", err)
	content := string(bytes)

	for _, line := range strings.Split(content, "\n") {
		switch {
		case strings.HasPrefix(line, OBJECT):
			objGrp = append(objGrp, line[len(OBJECT):])
		case strings.HasPrefix(line, LIST):
			listPathGrp = append(listPathGrp, cmn.RmTailFromLast(line[len(LIST):], "/"))

			// case others:
		}
	}

	// Println(true, objGrp...)
	// fmt.Println("-----------------------------")

	// Println(false, listPathGrp...)
	// fmt.Println("-----------------------------")

	m := MapOfListPathGrp(ObjGrpWithList(SEP, listPathGrp...), SEP, listPathGrp...)
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	toml, goStruct := PrintListPathGrp4Cfg(m, "ListAttrs")
	// fmt.Println(toml)
	// fmt.Println(goStruct)

	tt, err := ioutil.ReadFile(tomlTemplatePath)
	cmn.FailOnErr("%v", err)
	toml = strings.Replace(string(tt), ReplSignTOML, toml, 1)
	ioutil.WriteFile(tomlOutPath, []byte(toml), 0666)

	st, err := ioutil.ReadFile(structTemplatePath)
	cmn.FailOnErr("%v", err)
	goStruct = strings.Replace(string(st), ReplSignGO, goStruct, 1)
	ioutil.WriteFile(structGoPath, []byte(goStruct), 0666)
}
