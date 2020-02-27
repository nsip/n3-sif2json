package main

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

	l2j := NewCfg("./List2JSON.toml").(*List2JSON)
	InitCfgBuf(*l2j, l2j.Sep) // Init Global Maps
	fPln(GetLoadedObjects())

	obj := "PurchaseOrder"
	YieldJSON4OneCfg(obj, l2j.Sep, "./data", "[]", l2j.JQDir, true, false)
	fPln(GetAllFullPaths(obj, l2j.Sep))

	if _, ok := mObjMaxLenOfPath[obj]; !ok {
		fPln("Not Init")
	}
	fPln(mObjMaxLenOfPath[obj])
}

func TestMakeJSON(t *testing.T) {
	YieldJSONBySIF("./List2JSON.toml", "./Num2JSON.toml", "./Bool2JSON.toml")
}
