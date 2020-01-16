package cvt2json

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"sort"

	cmn "github.com/cdutwhu/json-util/common"
	pp "github.com/cdutwhu/json-util/preprocess"
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
	segs := sSplitRev(path, sep)
	for i, seg := range segs {
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
	cmn.FailOnErr("MakeJSON Fatal: %v", e)
	return string(jsonbytes)
}

// ----------------------------------------------- //

// YieldJSON4OneCfg :
func YieldJSON4OneCfg(obj, sep, outDir, jsonVal, jqDir string, levelized bool) {
	if outDir[len(outDir)-1] != '/' {
		outDir += "/"
	}
	path := outDir + obj + "/"

	// delete all obsolete json files when new config-json files are coming
	cmn.FailOnErr("%v", os.RemoveAll(path))
	fPf("%s is removed\n", path)
	cmn.FailOnErr("%v", os.MkdirAll(path, os.ModePerm))
	fPf("%s is created\n", path)

	if levelized {
		for lvl := 1; lvl < 100; lvl++ {
			if paths, valid := GetLvlFullPaths(obj, sep, lvl); valid {
				mm := MakeMap(paths, sep, jsonVal)
				if mm == nil || len(mm) == 0 {
					continue
				}
				// jsonstr := MakeJSON(mm)
				jsonstr := pp.FmtJSONStr(MakeJSON(mm), jqDir) // format jsonstr ( Only single thread use this line )
				ioutil.WriteFile(fSf("%s%d.json", path, lvl), []byte(jsonstr), 0666)
			} else {
				break
			}
		}
	} else {
		paths := GetAllFullPaths(obj, sep)
		mm := MakeMap(paths, sep, jsonVal)
		// jsonstr := MakeJSON(mm)
		jsonstr := pp.FmtJSONStr(MakeJSON(mm), jqDir) // format jsonstr ( Only single thread use this line )
		ioutil.WriteFile(fSf("%s%s.json", path, obj), []byte(jsonstr), 0666)
	}
}

// YieldCfgJSON4LIST :
func YieldCfgJSON4LIST(cfgPath, jsonVal string) {

	ICfg := NewCfg(cfgPath)
	cmn.FailOnCondition(ICfg == nil, "%v", fEf("LIST Configuration File Couldn't Be Loaded"))

	cfg := ICfg.(*list2json)
	cmn.FailOnCondition(cfg.Sep == "", "%v", fEf("Config-[Sep] loaded error"))
	cmn.FailOnCondition(cfg.JQDir == "", "%v", fEf("Config-[JQDir] loaded error"))

	InitCfgBuf(*cfg, cfg.Sep) // Init Global Maps
	for _, obj := range GetLoadedObjects() {
		YieldJSON4OneCfg(obj, cfg.Sep, cfg.CfgJSONOutDir, jsonVal, cfg.JQDir, true)
	}

	// lsObj := GetLoadedObjects()
	// wg := sync.WaitGroup{}
	// wg.Add(len(lsObj))
	// for _, obj := range lsObj {
	// 	go func(obj, sep, outDir, jsonVal, jqDir string) {
	// 		defer wg.Done()
	// 		YieldJSON4OneCfg(obj, sep, outDir, jsonVal, jqDir)
	// 	}(obj, cfg.Sep, cfg.CfgJSONOutDir, jsonVal, cfg.JQDir)
	// }
	// wg.Wait()
}

// YieldCfgJSON4NUM :
func YieldCfgJSON4NUM(cfgPath, jsonVal string) {

	ICfg := NewCfg(cfgPath)
	cmn.FailOnCondition(ICfg == nil, "%v", fEf("NUMERIC Configuration File Couldn't Be Loaded"))

	cfg := ICfg.(*num2json)
	cmn.FailOnCondition(cfg.Sep == "", "%v", fEf("Config-[Sep] loaded error"))
	cmn.FailOnCondition(cfg.JQDir == "", "%v", fEf("Config-[JQDir] loaded error"))

	InitCfgBuf(*cfg, cfg.Sep) // Init Global Maps
	for _, obj := range GetLoadedObjects() {
		YieldJSON4OneCfg(obj, cfg.Sep, cfg.CfgJSONOutDir, jsonVal, cfg.JQDir, false)
	}
}

// YieldCfgJSON4BOOL :
func YieldCfgJSON4BOOL(cfgPath, jsonVal string) {

	ICfg := NewCfg(cfgPath)
	cmn.FailOnCondition(ICfg == nil, "%v", fEf("BOOLEAN Configuration File Couldn't Be Loaded"))

	cfg := ICfg.(*bool2json)
	cmn.FailOnCondition(cfg.Sep == "", "%v", fEf("Config-[Sep] loaded error"))
	cmn.FailOnCondition(cfg.JQDir == "", "%v", fEf("Config-[JQDir] loaded error"))

	InitCfgBuf(*cfg, cfg.Sep) // Init Global Maps
	for _, obj := range GetLoadedObjects() {
		YieldJSON4OneCfg(obj, cfg.Sep, cfg.CfgJSONOutDir, jsonVal, cfg.JQDir, false)
	}
}
