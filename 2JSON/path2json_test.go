package cvt2json

import (
	"testing"
)

func TestOthers(t *testing.T) {

	m1 := MakeOneMap("PurchaseOrder~PurchasingItems~PurchasingItem~ExpenseAccounts", "~", "[]")
	m2 := MakeOneMap("PurchaseOrder~PurchasingItems~PurchasingItem1", "~", "[]")
	m3 := MakeOneMap("PurchaseOrder~PurchasingItems~PurchasingItem~ExpenseAccounts1", "~", "[]")
	m4 := MakeOneMap("PurchaseOrder~PurchasingItems1", "~", "[]")
	mm := MergeMaps(m1, m2, m3, m4)
	fPln(mm)
	fPln()

	// fPln(GetAllObjects())
	cfg := NewCfg("./config/Path2JSON.toml").(*Path2JSON)
	InitAllListAttrPaths(*cfg, cfg.Sep) // Init Global Maps
	obj := "PurchaseOrder"
	// YieldJSONListAttr4OneCfg(obj, cfg.Sep, "./data", "[]", cfg.JQDir)
	fPln(GetAllLAttrs(obj, cfg.Sep))

	if _, ok := mObjMaxLenOfLAttr[obj]; !ok {
		fPln("Not Init")
	}
	fPln(mObjMaxLenOfLAttr[obj])
}

func TestMakeJSON(t *testing.T) {
	YieldJSONListAttrCfg("./config/Path2JSON.toml", "../ListAttr", "[]")
}
