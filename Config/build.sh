#!/bin/bash

set -e

# sudo password
sudopwd="password"

# generate config.go for [Server]
echo $sudopwd | sudo -S env "PATH=$PATH" go test -v -timeout 1s -count=1 ./ -run TestRegCfg -args `whoami`