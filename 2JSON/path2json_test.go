package cvt2json

import (
	"testing"
)

func TestOthers(t *testing.T) {
	// fPln(GetAllObjects())
	// obj := "PurchaseOrder"
	// YieldJSON4OneCfg(obj, cfg.Sep, "./data", "[]")
	// LAs := GetAllLAttrs(obj)
	// fPln(GetLsAttr(obj, cfg.Sep, 1))
	// MakeJSON(obj, cfg.Sep, "[]")
	// m1 := MakeOneMap("PurchaseOrder~PurchasingItems~PurchasingItem~ExpenseAccounts", cfg.Sep, "[]")
	// m2 := MakeOneMap("PurchaseOrder~PurchasingItems~PurchasingItem1", cfg.Sep, "[]")
	// m3 := MakeOneMap("PurchaseOrder~PurchasingItems~PurchasingItem~ExpenseAccounts1", cfg.Sep, "[]")
	// m4 := MakeOneMap("PurchaseOrder~PurchasingItems1", cfg.Sep, "[]")
	// mm := MergeMaps(m1, m2, m3, m4)
	// if _, ok := mObjMLenOfLAttr[""]; !ok {
	// 	fPln("Not Init")
	// }
	// fPln(mObjMLenOfLAttr[obj])
}

func TestMakeJSON(t *testing.T) {
	YieldJSONListAttrCfg("./config/Path2JSON.toml", "../ListAttr", "[]")
}
