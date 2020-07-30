#!/bin/bash

set -e
ORIGINALPATH=`pwd`

red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
reset=`tput sgr0`

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

cd Server && ./build.sh && cd $ORIGINALPATH && echo "${green}Server Built${reset}"
cd Client && ./build.sh && cd $ORIGINALPATH && echo "${green}Client Built${reset}"
