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

// initListAttrPaths :
func initListAttrPaths(objListCfg interface{}, name, sep string) {
	// nameType := reflect.TypeOf(objListCfg).Name()
	value := reflect.ValueOf(objListCfg)
	nField := value.NumField()

	// for [ListAttrs] version,
	// [nField] should be 1 as all paths have been wrapped into [ListAttrs] Array
	for i := 0; i < nField; i++ {

		// [L1], [L2], [L3] ... version
		// # path := name + sep + fSp(value.Field(i).Interface())
		// path := fSp(value.Field(i).Interface())
		// mObjLAttrs[name] = append(mObjLAttrs[name], path)
		// if n := sCount(path, sep) + 1; mObjMaxLenOfLAttr[name] < n {
		// 	mObjMaxLenOfLAttr[name] = n
		// }

		// [ListAttrs] version
		lsPath := fSp(value.Field(i).Interface())
		lsPath = lsPath[1 : len(lsPath)-1]
		mObjLAttrs[name] = append(mObjLAttrs[name], sSplit(lsPath, " ")...)
		for _, path := range mObjLAttrs[name] {
			if n := sCount(path, sep) + 1; mObjMaxLenOfLAttr[name] < n {
				mObjMaxLenOfLAttr[name] = n
			}
		}
	}
	sort.SliceStable(mObjLAttrs[name], func(i, j int) bool {
		return sCount(mObjLAttrs[name][i], sep) < sCount(mObjLAttrs[name][j], sep)
	})
}

// InitAllListAttrPaths :
func InitAllListAttrPaths(cfg interface{}, sep string) {
	value := reflect.ValueOf(cfg)
	nField, valType := value.NumField(), value.Type()
	for i := 0; i < nField; i++ {
		fVal, fValTyp := value.Field(i), valType.Field(i)
		// nameType := reflect.TypeOf(fVal.Interface()).Name()
		// fPln(nameType)
		if fVal.Kind() == reflect.Struct {
			initListAttrPaths(fVal.Interface(), fValTyp.Name, sep)
			lsObjects = append(lsObjects, fValTyp.Name)
		}
	}
}

// GetAllObjects :
func GetAllObjects() []string {
	return append([]string{}, lsObjects...)
}

// GetAllLAttrs :
func GetAllLAttrs(obj, sep string) (LAs []string) {
	for _, la := range mObjLAttrs[obj] {
		// fPln(la)
		LAs = append(LAs, obj+sep+la)
	}
	return
}

// GetLAttrs :
func GetLAttrs(obj, sep string, lvl int) (LAs []string, valid bool) {
	if lvl > mObjMaxLenOfLAttr[obj] {
		return nil, false
	}
	for _, la := range mObjLAttrs[obj] {
		if lvl == sCount(la, sep)+1 {
			LAs = append(LAs, obj+sep+la)
		}
	}
	return LAs, true
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

// YieldJSONListAttr4OneCfg :
func YieldJSONListAttr4OneCfg(obj, sep, outDir, jsonVal, jqDir string) {
	if outDir[len(outDir)-1] != '/' {
		outDir += "/"
	}
	path := outDir + obj + "/"

	// delete all obsolete json files when new config-json files are coming
	cmn.FailOnErr("%v", os.RemoveAll(path))
	fPf("%s is removed\n", path)
	cmn.FailOnErr("%v", os.MkdirAll(path, os.ModePerm))
	fPf("%s is created\n", path)

	for lvl := 1; lvl < 100; lvl++ {
		if LAs, valid := GetLAttrs(obj, sep, lvl); valid {
			mm := MakeMap(LAs, sep, jsonVal)
			if mm == nil || len(mm) == 0 {
				continue
			}
			// jsonstr := MakeJSON(mm)
			jsonstr := pp.FmtJSONStr(MakeJSON(mm), jqDir) // format jsonstr ( !! pp syscall doesn't work properly for parallel )
			ioutil.WriteFile(fSf("%s%d.json", path, lvl), []byte(jsonstr), 0666)
		} else {
			break
		}
	}
}

// YieldJSONListAttrCfg :
func YieldJSONListAttrCfg(cfgPath, jsonVal string) {
	ICfg := NewCfg(cfgPath)
	cmn.FailOnCondition(ICfg == nil, "%v", fEf("ListAttribute Configuration File Couldn't Be Loaded"))
	cfg := ICfg.(*cfg2json)
	cmn.FailOnCondition(cfg.Sep == "", "%v", fEf("Config-[Sep] loaded error"))
	cmn.FailOnCondition(cfg.JQDir == "", "%v", fEf("Config-[JQDir] loaded error"))
	InitAllListAttrPaths(*cfg, cfg.Sep) // Init Global Maps

	for _, obj := range GetAllObjects() {
		YieldJSONListAttr4OneCfg(obj, cfg.Sep, cfg.CfgJSONOutDir, jsonVal, cfg.JQDir)
	}

	// lsObj := GetAllObjects()
	// wg := sync.WaitGroup{}
	// wg.Add(len(lsObj))
	// for _, obj := range lsObj {
	// 	go func(obj, sep, outDir, jsonVal, jqDir string) {
	// 		defer wg.Done()
	// 		YieldJSONListAttr4OneCfg(obj, sep, outDir, jsonVal, jqDir)
	// 	}(obj, cfg.Sep, cfg.CfgJSONOutDir, jsonVal, cfg.JQDir)
	// }
	// wg.Wait()
}
