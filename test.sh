#!/bin/bash
set -e

R=`tput setaf 1`
G=`tput setaf 2`
Y=`tput setaf 3`
B=`tput setaf 4`
W=`tput sgr0`

printf "\n"

ip="192.168.31.168:1324/"
base=$ip"n3-sif2json/v0.3.5/"


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


title='SIF to JSON Test'
url=$base"sif2json"
file="@./data/examples347/Activity_0.xml"
scode=`curl -X POST $url -d $file -w "%{http_code}" -s -o /dev/null`
if [ $scode -ne 200 ]; then
    echo "${Y}${title}${W}"
else
    echo "${G}${title}${W}"
fi
out=Activity_0.json
curl -X POST $url -d $file > $out
cat $out
printf "\n"


title='JSON to SIF Test'
url=$base"json2sif"
file="@./Activity_0.json"
scode=`curl -X POST $url -d $file -w "%{http_code}" -s -o /dev/null`
if [ $scode -ne 200 ]; then
    echo "${Y}${title}${W}"
else
    echo "${G}${title}${W}"
fi
curl -X POST $url -d $file
printf "\n"