#!/bin/bash
set -e

# all api
echo 'SIF2JSON all API Paths'
curl 192.168.31.168:1324/
echo ''

# SIF to JSON
echo 'SIF to JSON Test'
curl -X POST 192.168.31.168:1324/n3-sif2json/v0.3.4/sif2json -d '@./data/examples347/Activity_0.xml' > Activity_0.json
echo ''

# JSON to SIF
echo 'JSON to SIF Test'
curl -X POST 192.168.31.168:1324/n3-sif2json/v0.3.4/json2sif -d '@./Activity_0.json'
echo ''