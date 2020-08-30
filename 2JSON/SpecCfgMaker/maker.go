package main

import (
	"encoding/json"
	"os"
	"reflect"
	"sort"

	"github.com/cdutwhu/n3-util/n3err"
	"github.com/peterbourgon/mergemap"
)

// initGlobalMaps :
func initGlobalMaps(oneObjPathList interface{}, name, sep string) {
	// nameType := reflect.TypeOf(oneObjPathList).Name()
	value := reflect.ValueOf(oneObjPathList)
	nField := value.NumField()

	// for [****] version,
	// [nField] should be 1 as all paths have been wrapped into [****] Array
	for i := 0; i < nField; i++ {
		// [****] version
		lsPath := fSp(value.Field(i).Interface())
		lsPath = lsPath[1 : len(lsPath)-1]
		mObjPaths[name] = append(mObjPaths[name], sSplit(lsPath, " ")...)
		for _, path := range mObjPaths[name] {
			if n := sCount(path, sep) + 1; mObjMaxLenOfPath[name] < n {
				mObjMaxLenOfPath[name] = n
			}
		}
	}
	sort.SliceStable(mObjPaths[name], func(i, j int) bool {
		return sCount(mObjPaths[name][i], sep) < sCount(mObjPaths[name][j], sep)
	})
}

// InitCfgBuf :
func InitCfgBuf(cfg interface{}, sep string) {
	clearBuf()
	value := reflect.ValueOf(cfg)
	nField, valType := value.NumField(), value.Type()
	for i := 0; i < nField; i++ {
		fVal, fValTyp := value.Field(i), valType.Field(i)
		// nameType := reflect.TypeOf(fVal.Interface()).Name()
		// fPln(nameType)
		if fVal.Kind() == reflect.Struct {
			initGlobalMaps(fVal.Interface(), fValTyp.Name, sep)
			lsObjects = append(lsObjects, fValTyp.Name)
		}
	}
}

// GetLoadedObjects :
func GetLoadedObjects() []string {
	return append([]string{}, lsObjects...)
}

// GetAllFullPaths :
func GetAllFullPaths(obj, sep string) (paths []string) {
	for _, path := range mObjPaths[obj] {
		// fPln(path)
		paths = append(paths, obj+sep+path)
	}
	return
}

// GetLvlFullPaths :
func GetLvlFullPaths(obj, sep string, lvl int) (paths []string, valid bool) {
	if lvl > mObjMaxLenOfPath[obj] {
		return nil, false
	}
	for _, path := range mObjPaths[obj] {
		if lvl == sCount(path, sep)+1 {
			paths = append(paths, obj+sep+path)
		}
	}
	return paths, true
}

// -------------------------------------------------- //

// MakeBasicMap :
func MakeBasicMap(field string, value interface{}) map[string]interface{} {
	return map[string]interface{}{field: value}
}

// MakeOneMap :
func MakeOneMap(path, sep, valsymbol string) map[string]interface{} {
	var v interface{}
	for i, seg := range splitRev(path, sep) {
		if i == 0 {
			v = valsymbol
		}
		v = MakeBasicMap(seg, v)
	}
	return v.(map[string]interface{})
}

// MergeMaps :
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	var v map[string]interface{}
	for i, m := range maps {
		if i == 0 {
			v = m
		} else {
			v = mergemap.Merge(v, m)
		}
	}
	return v
}

// MakeMap :
func MakeMap(paths []string, sep, valsymbol string) map[string]interface{} {
	maps := []map[string]interface{}{}
	for _, path := range paths {
		maps = append(maps, MakeOneMap(path, sep, valsymbol))
	}
	return MergeMaps(maps...)
}

// MakeJSON :
func MakeJSON(m map[string]interface{}) string {
	jsonbytes, e := json.Marshal(m)
	failOnErr("MakeJSON Fatal: %v", e)
	return string(jsonbytes)
}

// ----------------------------------------------- //

// YieldJSON4OneCfg :
func YieldJSON4OneCfg(obj, sep, outDir, jsonVal string, levelized, extContent bool) {
	if outDir[len(outDir)-1] != '/' {
		outDir += "/"
	}
	path := outDir + obj + "/"

	// delete all obsolete json files when new config-json files are coming
	failOnErr("%v", os.RemoveAll(path))
	fPf("%s is removed\n", path)
	failOnErr("%v", os.MkdirAll(path, os.ModePerm))
	fPf("%s is created\n", path)

	if levelized {
		for lvl := 1; lvl < 100; lvl++ {
			if paths, valid := GetLvlFullPaths(obj, sep, lvl); valid {
				mm := MakeMap(paths, sep, jsonVal)
				if mm == nil || len(mm) == 0 {
					continue
				}
				jsonstr := MakeJSON(mm)
				jsonfmt := fmtJSON(jsonstr, 2)
				mustWriteFile(fSf("%s%d.json", path, lvl), []byte(jsonfmt))
			} else {
				break
			}
		}
	} else {
		paths := GetAllFullPaths(obj, sep)
		mm := MakeMap(paths, sep, jsonVal)
		jsonstr := MakeJSON(mm)
		jsonfmt := fmtJSON(jsonstr, 2)
		mustWriteFile(fSf("%s0.json", path), []byte(jsonfmt))

		if extContent {
			// extend jsonstr, such as xml->json '#content', "30" => { "#content": "30" }
			jsonext := sReplaceAll(jsonstr, fSf(`"%s"`, jsonVal), fSf(`{"#content": "%s"}`, jsonVal))
			jsonextfmt := fmtJSON(jsonext, 2)
			mustWriteFile(fSf("%s1.json", path), []byte(jsonextfmt))
		}
	}
}

// YieldJSONBySIFList :
func YieldJSONBySIFList(cfgPath string) {

	ICfg := NewCfg(cfgPath)
	failOnErrWhen(ICfg == nil, "%v: LIST, %s", n3err.CFG_INIT_ERR, cfgPath)

	l2j := ICfg.(*List2JSON)
	failOnErrWhen(l2j.Sep == "", "%v: LIST-[Sep]", n3err.CFG_INIT_ERR)

	InitCfgBuf(*l2j, l2j.Sep) // Init Global Maps
	for _, obj := range GetLoadedObjects() {
		YieldJSON4OneCfg(obj, l2j.Sep, l2j.CfgJSONOutDir, l2j.CfgJSONValue, true, false)
	}

	// lsObj := GetLoadedObjects()
	// wg := sync.WaitGroup{}
	// wg.Add(len(lsObj))
	// for _, obj := range lsObj {
	// 	go func(obj, sep, outDir, l2j.CfgJSONValue string) {
	// 		defer wg.Done()
	// 		YieldJSON4OneCfg(obj, sep, outDir, l2j.CfgJSONValue, jqDir)
	// 	}(obj, l2j.Sep, l2j.CfgJSONOutDir, l2j.CfgJSONValue)
	// }
	// wg.Wait()
}

// YieldJSONBySIFNum :
func YieldJSONBySIFNum(cfgPath string) {

	ICfg := NewCfg(cfgPath)
	failOnErrWhen(ICfg == nil, "%v: NUMERIC, %s", n3err.CFG_INIT_ERR, cfgPath)

	n2j := ICfg.(*Num2JSON)
	failOnErrWhen(n2j.Sep == "", "%v: NUMERIC-[Sep]", n3err.CFG_INIT_ERR)

	InitCfgBuf(*n2j, n2j.Sep) // Init Global Maps
	for _, obj := range GetLoadedObjects() {
		YieldJSON4OneCfg(obj, n2j.Sep, n2j.CfgJSONOutDir, n2j.CfgJSONValue, false, true)
	}
}

// YieldJSONBySIFBool :
func YieldJSONBySIFBool(cfgPath string) {

	ICfg := NewCfg(cfgPath)
	failOnErrWhen(ICfg == nil, "%v: BOOLEAN, %s", n3err.CFG_INIT_ERR, cfgPath)

	b2j := ICfg.(*Bool2JSON)
	failOnErrWhen(b2j.Sep == "", "%v: BOOLEAN-[Sep]", n3err.CFG_INIT_ERR)

	InitCfgBuf(*b2j, b2j.Sep) // Init Global Maps
	for _, obj := range GetLoadedObjects() {
		YieldJSON4OneCfg(obj, b2j.Sep, b2j.CfgJSONOutDir, b2j.CfgJSONValue, false, true)
	}
}

// YieldJSONBySIF :
func YieldJSONBySIF(listCfg, numCfg, boolCfg string) {
	YieldJSONBySIFList(listCfg)
	YieldJSONBySIFNum(numCfg)
	YieldJSONBySIFBool(boolCfg)
}

func main() {
	YieldJSONBySIF(os.Args[2], os.Args[3], os.Args[4])
	fPln("JSON Config files are created")
}
