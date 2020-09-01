#!/bin/bash

set -e
shopt -s extglob

ORIPATH=`pwd`

cd ./Server/ && ./clean.sh && cd $ORIPATH && echo "Server clean"
cd ./2JSON/ && rm -f *auto*.go *.json && cd $ORIPATH && echo "2JSON clean"
cd ./2SIF/ && rm -f *auto*.go *.xml && cd $ORIPATH && echo "2SIF clean"
cd ./SIFSpec/ && ./clean.sh && cd $ORIPATH && echo "SIFSpec clean"

rm -rf ./data/json/ ./data/sif/
rm -f ./*.log ./*.json ./*.xml

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f
for f in $(find ./ -name '*.log' -or -name '*.doc'); do rm $f; done