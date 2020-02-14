#!/bin/bash

shopt -s extglob

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

cd ./2SIF
./clean.sh
cd -

rm -rf ./2JSON/SIFCfg/*
rm -f ./2JSON/config/Bool2JSON.toml ./2JSON/config/List2JSON.toml ./2JSON/config/Num2JSON.toml
rm -rf ./data/*.json ./data/*_out.xml
