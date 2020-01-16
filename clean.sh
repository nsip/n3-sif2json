#!/bin/bash

shopt -s extglob

rm -rf ./2JSON/data/
rm -f ./2JSON/config/bool2json.toml ./2JSON/config/list2json.toml ./2JSON/config/num2json.toml
rm -rf ./data/!(Activity.xml)
