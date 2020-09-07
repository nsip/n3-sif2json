#!/bin/bash
set -e

R=`tput setaf 1`
G=`tput setaf 2`
Y=`tput setaf 3`
B=`tput setaf 4`
W=`tput sgr0`

printf "\n"

ip="192.168.31.168:1324/"
base=$ip"n3-sif2json/v0.3.7/"

title='SIF2JSON all API Paths'
url=$ip
scode=`curl --write-out "%{http_code}" --silent --output /dev/null $url`
if [ $scode -ne 200 ]; then
    echo "${Y}${title}${W}"
else
    echo "${G}${title}${W}"
fi
curl -i $url
printf "\n"

sv=3.4.6

examples346=./data/examples346/*
for f in $examples346
do
    title='SIF to JSON Test @ '$f
    url=$base"sif2json?sv="$sv
    file="@"$f
    scode=`curl -X POST $url -d $file -w "%{http_code}" -s -o /dev/null`
    if [ $scode -ne 200 ]; then
        echo "${Y}${title}${W}"
        exit 1
    else
        echo "${G}${title}${W}"
    fi

    outdir=./data/output/json346/
    mkdir -p $outdir
    out=$outdir`basename $f .xml`.json
    curl -X POST $url -d $file > $out
    cat $out
    printf "\n"
done


jsonfiles=./data/output/json346/*
for f in $jsonfiles
do
    title='JSON to SIF Test @ '$f
    url=$base"json2sif?sv="$sv
    file="@"$f
    scode=`curl -X POST $url -d $file -w "%{http_code}" -s -o /dev/null`
    if [ $scode -ne 200 ]; then
        echo "${Y}${title}${W}"
        exit 1
    else
        echo "${G}${title}${W}"
    fi

    outdir=./data/output/sif346/
    mkdir -p $outdir
    out=$outdir`basename $f .json`.xml
    curl -X POST $url -d $file > $out
    cat $out
    printf "\n"
done
