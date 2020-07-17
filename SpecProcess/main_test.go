package main

import "testing"

func TestMain(t *testing.T) {
	main()
}

func TestGenTomlAndStruct(t *testing.T) {
	GenTomlAndGoSrc(
		"../SIFSpec/3.4.7.txt",
		"../2JSON/SpecCfgMaker/base-go/spec",
		"../2JSON/SpecCfgMaker/base-toml/List2JSON",
		"../2JSON/SpecCfgMaker/base-toml/Num2JSON",
		"../2JSON/SpecCfgMaker/base-toml/Bool2JSON",
		"../2JSON/SpecCfgMaker/",
	)
}
