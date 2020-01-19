package cvt2json

import (
	"testing"
)

func TestMakeMap(t *testing.T) {
	m1 := MakeOneMap("PurchaseOrder~PurchasingItems~PurchasingItem~ExpenseAccounts", "~", "[]")
	m2 := MakeOneMap("PurchaseOrder~PurchasingItems~PurchasingItem1", "~", "[]")
	m3 := MakeOneMap("PurchaseOrder~PurchasingItems~PurchasingItem~ExpenseAccounts1", "~", "[]")
	m4 := MakeOneMap("PurchaseOrder~PurchasingItems1", "~", "[]")
	mm := MergeMaps(m1, m2, m3, m4)
	fPln(mm)
}

func TestOthers(t *testing.T) {

	cfg := NewCfg("./config/list2json.toml").(*list2json)
	InitCfgBuf(*cfg, cfg.Sep) // Init Global Maps
	fPln(GetLoadedObjects())

	obj := "PurchaseOrder"
	YieldJSON4OneCfg(obj, cfg.Sep, "./data", "[]", cfg.JQDir, true, false)
	fPln(GetAllFullPaths(obj, cfg.Sep))

	if _, ok := mObjMaxLenOfPath[obj]; !ok {
		fPln("Not Init")
	}
	fPln(mObjMaxLenOfPath[obj])
}

func TestMakeJSON(t *testing.T) {
	YieldCfgJSON4LIST("./config/list2json.toml")
	YieldCfgJSON4NUM("./config/num2json.toml")
	YieldCfgJSON4BOOL("./config/bool2json.toml")
}
