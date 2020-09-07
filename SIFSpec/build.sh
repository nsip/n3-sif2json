#!/bin/bash

set -e

r=`tput setaf 1`
g=`tput setaf 2`
y=`tput setaf 3`
w=`tput sgr0`

# sudo password
sudopwd="password"

gotestdir="./1_txt2toml/"

# create toml files
go test -v -timeout 1s -count=1 $gotestdir -run TestGenTomlAndStruct

# register toml under n3cfg
echo $sudopwd | sudo -S env "PATH=$PATH" go test -v -count=1 $gotestdir -run TestRegister

gotestdir="./2_toml2json/"

# create json files
go test -v -timeout 1s -count=1 $gotestdir -run TestMakeJSON

# create json files binaries
go test -v -timeout 1s -count=1 $gotestdir -run TestBinariseRes

###################

### Create toml files at building time ###

# if [ $# -ne 1 ]; then
#     echo "Input SIF(txt) Spec File Path"
#     exit -1
# fi

# if [ ! -f $1 ]; then
#     echo "SIF txt Spec does not exist"
#     exit -1
# fi

# SPEC=$1
# MAKER=../2JSON/SpecCfgMaker
# BASE_GO=$MAKER/base-go
# BASE_TOML=$MAKER/base-toml
# CGO_ENABLED=0 go run var.go main.go -- $SPEC $BASE_GO/spec $BASE_TOML/List2JSON $BASE_TOML/Num2JSON $BASE_TOML/Bool2JSON $MAKER/