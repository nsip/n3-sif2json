#!/bin/bash
rm -f ./go.sum
go get -u ./...

R=`tput setaf 1`
G=`tput setaf 2`
Y=`tput setaf 3`
W=`tput sgr0`

oripath=`pwd`

cd ./SIFSpec && ./build.sh && cd $oripath && echo "${G}SIF Spec Ready${W}"
cd ./Config && ./build.sh && cd $oripath && echo "${G}Config Prepared${W}"
cd ./Server && ./build.sh && cd $oripath && echo "${G}Server Built${W}"
