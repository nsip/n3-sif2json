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

// ***
func TestMakeJSON(t *testing.T) {
	YieldJSONBySIF(
		"../3.4.6/toml/List2JSON.toml",
		"../3.4.6/toml/Num2JSON.toml",
		"../3.4.6/toml/Bool2JSON.toml",
		"3.4.6",
	)
	YieldJSONBySIF(
		"../3.4.7/toml/List2JSON.toml",
		"../3.4.7/toml/Num2JSON.toml",
		"../3.4.7/toml/Bool2JSON.toml",
		"3.4.7",
	)
}

// ***
func TestBinariseRes(t *testing.T) {
	printFileBytes("sif346", "TXT346", "../3.4.6/res_txt.go", false, "../3.4.6.txt")
	printFileBytes("sif347", "TXT347", "../3.4.7/res_txt.go", false, "../3.4.7.txt")
	createDirBytes("sif346", "JSON346", "../3.4.6/json", "../3.4.6/res_json.go", false, "json", "346")
	createDirBytes("sif347", "JSON347", "../3.4.7/json", "../3.4.7/res_json.go", false, "json", "347")
}
