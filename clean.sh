#!/bin/bash

shopt -s extglob

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f

cd ./Server
./clean.sh
cd -

cd ./Client
./clean.sh
cd -

cd ./data
./clean.sh
cd -

cd ./2JSON
./clean.sh
cd -

cd ./2JSON/SpecCfgMaker
./clean.sh
cd -

cd ./2SIF
./clean.sh
cd -

rm -rf ./2JSON/SpecCfg/*
rm -rf ./data/*.json ./data/*_out.xml
