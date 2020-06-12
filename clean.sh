#!/bin/bash

set -e
shopt -s extglob

ORIGINALPATH=`pwd`

cd ./Server && ./clean.sh && cd $ORIGINALPATH && echo "Server clean"
cd ./Client && ./clean.sh && cd $ORIGINALPATH && echo "Client clean"
cd ./data && ./clean.sh && cd $ORIGINALPATH && echo "data clean"
cd ./2JSON && ./clean.sh && cd $ORIGINALPATH && echo "2JSON clean"
cd ./2JSON/SpecCfgMaker && ./clean.sh && cd $ORIGINALPATH && echo "SpecCfgMaker clean"
cd ./2SIF && ./clean.sh && cd $ORIGINALPATH && echo "2SIF clean"

rm -rf ./2JSON/SpecCfg/*
rm -rf ./data/json/ ./data/sif/
rm -f ./*.log

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f
for f in $(find ./ -name '*.log' -or -name '*.doc'); do rm $f; done