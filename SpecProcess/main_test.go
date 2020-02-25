package main

import "testing"

func TestMain(t *testing.T) {
	main()
}

func TestGenTomlAndStruct(t *testing.T) {
	GenTomlAndStruct(
		"../SIFSpec/out.txt",
		"../2JSON/config/base-go/config",
		"../2JSON/config/base-toml/List2JSON",
		"../2JSON/config/base-toml/Num2JSON",
		"../2JSON/config/base-toml/Bool2JSON",
		"../2JSON/config/",
	)
}
