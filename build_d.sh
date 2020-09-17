#!/bin/bash
rm -f ./go.sum
go get -u ./...

oripath=`pwd`

cd ./SIFSpec && ./build_d.sh && cd $oripath && echo "SIF Spec Ready"
cd ./Config && ./build_d.sh && cd $oripath && echo "Config Prepared"
cd ./Server && ./build_d.sh && cd $oripath && echo "Server Built"
