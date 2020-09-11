#!/bin/bash

set -e

r=`tput setaf 1`
g=`tput setaf 2`
y=`tput setaf 3`
w=`tput sgr0`

CGO_ENABLED=0 go run ./1_txt2toml/main.go
CGO_ENABLED=0 go run ./2_toml2json/config.go ./2_toml2json/main.go
