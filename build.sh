#!/bin/bash

set -e

red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
reset=`tput sgr0`

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
    echo "${yellow}WARN:${reset} No SIF(txt) Spec File Input"
fi

FILES="$@"
for f in $FILES; do 
    if [ -f $f ]; then
        fabs=`realpath "$f"`
        cd ./SpecProcess && ./build.sh "$fabs" && cd $ORIGINALPATH
        cd ./2JSON/SpecCfgMaker/ && ./build.sh && cd $ORIGINALPATH
    else
        echo "${red}$f Spec (txt) does not exist${reset}"
        exit -1
    fi
done

####

WORKPATH="./Preprocess"

# sudo password
sudopwd="password"

# generate config.go for [Server] [2JSON] [2SIF]
echo $sudopwd | sudo -S env "PATH=$PATH" go test -v -timeout 1s -count=1 $WORKPATH/CfgReg -run TestRegCfg -args `whoami` "server" "cvt2json" "cvt2sif"

# Trim Server config.toml for [goclient]
go test -v -timeout 1s -count=1 $WORKPATH/CfgGen -run TestMkCltCfg -args "Path" "Service" "Route" "Server" "Access"
echo "${green}goclient Config.toml Generated${reset}"

# generate config.go fo [goclient]
echo $sudopwd | sudo -S env "PATH=$PATH" go test -v -timeout 1s -count=1 $WORKPATH/CfgReg -run TestRegCfg -args `whoami` "goclient"

####

cd ./Server && ./build.sh && cd $ORIGINALPATH && echo "${green}Server Built${reset}"
