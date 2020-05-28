package client

import "testing"

func TestDO(t *testing.T) {
	str, err := DO(
		"./config.toml",
		"SIF2JSON",
		Args{
			File:      "../../data/examples/Activity_0.xml",
			Ver:       "3.4.5",
			WholeDump: true,
			ToNATS:    false,
		})
	fPln(str)
	fPln(err)
}
