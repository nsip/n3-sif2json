 #!/bin/bash

set -e

r=`tput setaf 1`
g=`tput setaf 2`
y=`tput setaf 3`
w=`tput sgr0`

# sudo password
sudopwd="cppcli"

workpath="./preprocess"

# generate config.go for [Server]
echo $sudopwd | sudo -S env "PATH=$PATH" go test -v -timeout 1s -count=1 $workpath/cfgreg -run TestRegCfg -args `whoami` "server"

# Trim Server config.toml for [goclient]
go test -v -timeout 1s -count=1 $workpath/cfggen -run TestMkCltCfg -args "Path" "Service" "Route" "Server" "Access"
echo "${g}goclient Config.toml Generated${w}"

# generate config.go fo [goclient]
echo $sudopwd | sudo -S env "PATH=$PATH" go test -v -timeout 1s -count=1 $workpath/cfgreg -run TestRegCfg -args `whoami` "goclient"

######################

rm -rf ./build

GOARCH=amd64
LDFLAGS="-s -w"
OUT=server

OUTPATH=./build/linux64/
mkdir -p $OUTPATH
CGO_ENABLED=0 GOOS="linux" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
mv $OUT $OUTPATH
cp ./config_rel.toml $OUTPATH'config.toml'

OUTPATH=./build/win64/
mkdir -p $OUTPATH
CGO_ENABLED=0 GOOS="windows" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT.exe
mv $OUT.exe $OUTPATH
cp ./config_rel.toml $OUTPATH'config.toml'

OUTPATH=./build/mac/
mkdir -p $OUTPATH
CGO_ENABLED=0 GOOS="darwin" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
mv $OUT $OUTPATH
cp ./config_rel.toml $OUTPATH'config.toml'

# GOARCH=arm
# OUTPATH=./build/linuxarm/
# mkdir -p $OUTPATH
# CGO_ENABLED=0 GOOS="linux" GOARCH="$GOARCH" GOARM=7 go build -ldflags="$LDFLAGS" -o $OUT
# mv $OUT $OUTPATH
# cp ./config_rel.toml $OUTPATH
