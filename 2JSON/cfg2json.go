package cvt2json

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"sort"

	pp "../preprocess"
	"github.com/peterbourgon/mergemap"
)

// GetAllObjects :
func GetAllObjects() []string {
	return append([]string{}, lsObjects...)
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

// initListAttrPaths :
func initListAttrPaths(objListCfg interface{}, name, sep string) {
	// nameType := reflect.TypeOf(objListCfg).Name()
	value := reflect.ValueOf(objListCfg)
	nField := value.NumField()
	for i := 0; i < nField; i++ {
		// path := name + sep + fSp(value.Field(i).Interface())
		path := fSp(value.Field(i).Interface())
		mObjLAttrs[name] = append(mObjLAttrs[name], path)

		if n := sCount(path, sep) + 1; mObjMLenOfLAttr[name] < n {
			mObjMLenOfLAttr[name] = n
		}
	}
	sort.SliceStable(mObjLAttrs[name], func(i, j int) bool {
		return sCount(mObjLAttrs[name][i], sep) < sCount(mObjLAttrs[name][j], sep)
	})
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
	if lvl > mObjMLenOfLAttr[obj] {
		return nil, false
	}
	for _, la := range mObjLAttrs[obj] {
		if lvl == sCount(la, sep)+1 {
			LAs = append(LAs, obj+sep+la)
		}
	}
	return LAs, true
}

// -------------------------------------------- //

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
	if jsonbytes, e := json.Marshal(m); e == nil {
		return string(jsonbytes)
	}
	panic("MakeJSON Fatal")
}

// YieldJSONListAttr4OneCfg :
func YieldJSONListAttr4OneCfg(obj, sep, outdir, jsonval string) {
	if outdir[len(outdir)-1] != '/' {
		outdir += "/"
	}
	path := outdir + obj + "/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	for lvl := 1; lvl < 100; lvl++ {
		if LAs, valid := GetLAttrs(obj, sep, lvl); valid {
			mm := MakeMap(LAs, sep, jsonval)
			if mm == nil || len(mm) == 0 {
				continue
			}
			// jsonstr := MakeJSON(mm)
			jsonstr := pp.FmtJSONStr(MakeJSON(mm), "../preprocess/utils") // format jsonstr
			ioutil.WriteFile(fSf("%s%d.json", path, lvl), []byte(jsonstr), 0666)
		} else {
			break
		}
	}
}

// YieldJSONListAttrCfg :
func YieldJSONListAttrCfg(cfgPath, outDir, jsonval string) {
	if cfg := NewCfg(cfgPath); cfg != nil {
		cfg := cfg.(*List)
		InitAllListAttrPaths(*cfg, cfg.Sep) // Init Global Maps
		for _, obj := range GetAllObjects() {
			YieldJSONListAttr4OneCfg(obj, cfg.Sep, outDir, jsonval)
		}
	}
}
