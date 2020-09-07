package main

import (
	"testing"

	"github.com/cdutwhu/n3-util/n3cfg"
)

// ***
func TestGenTomlAndStruct(t *testing.T) {
	GenTomlAndGoSrc("../3.4.6.txt", "../3.4.6/")
	GenTomlAndGoSrc("../3.4.7.txt", "../3.4.7/")
}

// *** echo password | sudo -S env "PATH=$PATH" go test -v -count=1 ./ -run TestRegister
func TestRegister(t *testing.T) {
	user := "qmiao"
	prj := n3cfg.PrjName()

	toml346 := "../3.4.6/toml/"
	n3cfg.Register(user, toml346+"List2JSON.toml", prj, "sif346list")
	n3cfg.Register(user, toml346+"Bool2JSON.toml", prj, "sif346bool")
	n3cfg.Register(user, toml346+"Num2JSON.toml", prj, "sif346num")

	toml347 := "../3.4.7/toml/"
	n3cfg.Register(user, toml347+"List2JSON.toml", prj, "sif347list")
	n3cfg.Register(user, toml347+"Bool2JSON.toml", prj, "sif347bool")
	n3cfg.Register(user, toml347+"Num2JSON.toml", prj, "sif347num")
}
