#!/bin/bash

set -e
ORIGINALPATH=`pwd`
VERSION="v0.1.0"

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
    echo "WARN: No SIF(txt) Spec File Input"
fi
FILES="$@"
for f in $FILES; do 
    if [ -f $f ]; then
        fabs=`realpath "$f"`
        cd ./SpecProcess && ./build.sh "$fabs" && cd $ORIGINALPATH
        cd ./2JSON/SpecCfgMaker/ && ./build.sh && cd $ORIGINALPATH
    else
        echo "$fabs Spec (txt) does not exist"
        exit -1
    fi
done
####

cd Server && ./build.sh && cd $ORIGINALPATH && echo "Server Built"
cd Client && ./build.sh && cd $ORIGINALPATH && echo "Client Built"
