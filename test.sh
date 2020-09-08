#!/bin/bash
set -e

R=`tput setaf 1`
G=`tput setaf 2`
Y=`tput setaf 3`
B=`tput setaf 4`
W=`tput sgr0`

printf "\n"

ip="192.168.31.168:1324/"
base=$ip"n3-sif2json/v0.4.2/"

title='SIF2JSON all API Paths'
url=$ip
scode=`curl --write-out "%{http_code}" --silent --output /dev/null $url`
if [ $scode -ne 200 ]; then
    echo "${Y}${title}${W}"
    exit 1
else
    echo "${G}${title}${W}"
fi
curl -i $url
printf "\n"

sv=3.4.7

SIFFiles=./data/examples/$sv/*
for f in $SIFFiles
do    
    title='SIF to JSON Test @ '$f
    url=$base"sif2json?sv=$sv"    
    file="@"$f
    scode=`curl -X POST $url -d $file -w "%{http_code}" -s -o /dev/null`
    if [ $scode -ne 200 ]; then
        echo "${Y}${title}${W}"
        exit 1
    else
        echo "${G}${title}${W}"
    fi

    jsonname=`basename $f .xml`.json
    outdir=./data/output/$sv/json/
    mkdir -p $outdir
    outfile=$outdir"$jsonname"
    curl -X POST $url -d $file > $outfile
    cat $outfile
    printf "\n"
done

JSONFiles=./data/output/$sv/json/*
for f in $JSONFiles
do  
    title='JSON to SIF Test @ '$f
    url=$base"json2sif?sv=$sv"
    file="@"$f
    scode=`curl -X POST $url -d $file -w "%{http_code}" -s -o /dev/null`
    if [ $scode -ne 200 ]; then
        echo "${Y}${title}${W}"
        exit 1
    else
        echo "${G}${title}${W}"
    fi

    sifname=`basename $f .json`.xml
    outdir=./data/output/$sv/sif/
    mkdir -p $outdir
    outfile=$outdir"$sifname"
    curl -X POST $url -d $file > $outfile
    cat $outfile
    printf "\n"
done

echo "${G}All Done${W}"