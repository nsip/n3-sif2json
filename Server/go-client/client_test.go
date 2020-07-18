package client

import (
	"testing"
)

func TestDO(t *testing.T) {

	str, err := DO(
		"./config.toml",
		"HELP",
		nil,
	)
	fPln(str)
	fPln(err)

	// bytes, err := ioutil.ReadFile("../../data/examples/NAPTest_0.xml")
	// failOnErr("%v", err)
	// str, err := DO(
	// 	"./config.toml",
	// 	"SIF2JSON",
	// 	&Args{
	// 		Data:   bytes,
	// 		Ver:    "3.4.7",
	// 		ToNATS: false,
	// 	},
	// )
	// fPln(str)
	// fPln(err)
	// ioutil.WriteFile("./out.json", []byte(str), 0666)
}
