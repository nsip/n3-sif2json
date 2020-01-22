package cvt2json

import (
	"testing"

	cfg "github.com/nsip/n3-sif2json/2JSON/config"
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

	l2j := cfg.NewCfg("./config/List2JSON.toml").(*cfg.List2JSON)
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
	YieldJSONBySIF("./config/List2JSON.toml", "./config/Num2JSON.toml", "./config/Bool2JSON.toml")
}
