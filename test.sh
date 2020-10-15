#!/bin/bash
set -e

R=`tput setaf 1`
G=`tput setaf 2`
Y=`tput setaf 3`
B=`tput setaf 4`
W=`tput sgr0`

printf "\n"

ip="192.168.31.168:1324/"    ### 
base=$ip"n3-sif2json/v0.4.10/"    ###

title='SIF2JSON all API Paths'
url=$ip
scode=`curl --write-out "%{http_code}" --silent --output /dev/null $url`
if [ $scode -ne 200 ]; then
    echo "${R}${title}${W}"
    exit 1
else
    echo "${G}${title}${W}"
fi
echo "curl $url"
curl -i $url
printf "\n"

# exit 0

sv=3.4.7

SIFFiles=./data/examples/$sv/*
for f in $SIFFiles
do    
    title='2JSON Test @ '$f
    url=$base"2json?sv=$sv"    
    file="@"$f
    scode=`curl -X POST $url -d $file -w "%{http_code}" -s -o /dev/null`
    if [ $scode -ne 200 ]; then
        echo "${R}${title}${W}"
        exit 1
    else
        echo "${G}${title}${W}"
    fi

    jsonname=`basename $f .xml`.json
    outdir=./data/output/$sv/json/
    mkdir -p $outdir
    outfile=$outdir"$jsonname"
    echo "curl -X POST $url -d $file"
    curl -X POST $url -d $file > $outfile
    cat $outfile
    printf "\n"
done

# exit 0

JSONFiles=./data/output/$sv/json/*
for f in $JSONFiles
do  
    title='2SIF Test @ '$f
    url=$base"2sif?sv=$sv"
    file="@"$f
    scode=`curl -X POST $url -d $file -w "%{http_code}" -s -o /dev/null`
    if [ $scode -ne 200 ]; then
        echo "${R}${title}${W}"
        exit 1
    else
        echo "${G}${title}${W}"
    fi

    sifname=`basename $f .json`.xml
    outdir=./data/output/$sv/sif/
    mkdir -p $outdir
    outfile=$outdir"$sifname"
    echo "curl -X POST $url -d $file"
    curl -X POST $url -d $file > $outfile
    cat $outfile
    printf "\n"
done

echo "${G}All Done${W}"