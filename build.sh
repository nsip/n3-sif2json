#!/bin/bash

set -e

r=`tput setaf 1`
g=`tput setaf 2`
y=`tput setaf 3`
w=`tput sgr0`

cd ./SIFSpec && ./build.sh && cd - && echo "${g}SIF Spec Ready${w}"
cd ./Config && ./build.sh && cd - && echo "${g}Config Registered${w}"
cd ./Server && ./build.sh && cd - && echo "${g}Server Built${w}"
