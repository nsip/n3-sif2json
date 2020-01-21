#!/bin/bash

shopt -s extglob

rm -rf ./2JSON/SIFCfg/*
rm -f ./2JSON/config/Bool2JSON.toml ./2JSON/config/List2JSON.toml ./2JSON/config/Num2JSON.toml
rm -rf ./data/!(Activity.xml)
