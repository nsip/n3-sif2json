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
	printFileBytes("sif346", "TXT", "../3.4.6/txt.go", false, "../3.4.6.txt")
	createDirBytes("sif346", "JSON_BOOL", "../3.4.6/json/BOOLEAN/", "../3.4.6/json_bool.go", false, "346", "json", "BOOLEAN")
	createDirBytes("sif346", "JSON_LIST", "../3.4.6/json/LIST/", "../3.4.6/json_list.go", false, "346", "json", "LIST")
	createDirBytes("sif346", "JSON_NUM", "../3.4.6/json/NUMERIC/", "../3.4.6/json_num.go", false, "346", "json", "NUMERIC")

	printFileBytes("sif347", "TXT", "../3.4.7/txt.go", false, "../3.4.7.txt")
	createDirBytes("sif347", "JSON_BOOL", "../3.4.7/json/BOOLEAN/", "../3.4.7/json_bool.go", false, "347", "json", "BOOLEAN")
	createDirBytes("sif347", "JSON_LIST", "../3.4.7/json/LIST/", "../3.4.7/json_list.go", false, "347", "json", "LIST")
	createDirBytes("sif347", "JSON_NUM", "../3.4.7/json/NUMERIC/", "../3.4.7/json_num.go", false, "347", "json", "NUMERIC")
}
