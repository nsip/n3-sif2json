#!/bin/bash

set -e

r=`tput setaf 1`
g=`tput setaf 2`
y=`tput setaf 3`
w=`tput sgr0`

# sudo password
sudopwd="cppcli"

ORIGINALPATH=`pwd`

####
# if [ $# -ne 1 ]; then
#     echo "Input SIF(txt) Spec File Path"
#     exit -1
# fi
# if [ ! -f "$1" ]; then
#     echo "SIF txt Spec does not exist"
#     exit -1
# fi
# fabs=`realpath "$1"`
# cd ./SpecProcess && ./build.sh "$fabs" && cd $ORIGINALPATH
# cd ./2JSON/SpecCfgMaker/ && ./build.sh && cd $ORIGINALPATH

if [ $# -lt 1 ]; then
    echo "${y}WARN:${w} No SIF(txt) Spec File Input"
fi

FILES="$@"
for f in $FILES; do 
    if [ -f $f ]; then
        fabs=`realpath "$f"`
        cd ./SpecProcess && ./build.sh "$fabs" && cd $ORIGINALPATH
        

        cd ./2JSON/SpecCfgMaker/ && ./build.sh && cd $ORIGINALPATH
    else
        echo "${r}$f Spec (txt) does not exist${w}"
        exit -1
    fi
done

####

WORKPATH="./Preprocess"

# generate config.go for [Server] [2JSON] [2SIF]
echo $sudopwd | sudo -S env "PATH=$PATH" go test -v -timeout 1s -count=1 $WORKPATH/CfgReg -run TestRegCfg -args `whoami` "server" "cvt2json" "cvt2sif"

# Trim Server config.toml for [goclient]
go test -v -timeout 1s -count=1 $WORKPATH/CfgGen -run TestMkCltCfg -args "Path" "Service" "Route" "Server" "Access"
echo "${g}goclient Config.toml Generated${w}"

# generate config.go fo [goclient]
echo $sudopwd | sudo -S env "PATH=$PATH" go test -v -timeout 1s -count=1 $WORKPATH/CfgReg -run TestRegCfg -args `whoami` "goclient"

####

cd ./Server && ./build.sh && cd $ORIGINALPATH && echo "${g}Server Built${w}"
