package main

import "os"

func main() {
	if len(os.Args) < 4 {
		fPln("You are not allowed to use this tool to create JSON config files unless fully understand what you are doing.\n" +
			"Project author or other admins are advised to do this for creating SIF Specifications JSON config files for this project.\n" +
			"If you still want to do it by yourself, make sure <config.go>, <List2JSON.toml>, <Num2JSON.toml> and <Bool2JSON.toml> are all created.\n" +
			"Then input following arguments orderly:\n" +
			"  1. List2JSON.toml file path\n" +
			"  2. Num2JSON.toml file path\n" +
			"  3. Bool2JSON.toml file path")
		return
	}
	listCfgToml := os.Args[1]
	numCfgToml := os.Args[2]
	boolCfgToml := os.Args[3]
	YieldJSONBySIF(listCfgToml, numCfgToml, boolCfgToml)
	fPln("JSON Config files are created")
}
