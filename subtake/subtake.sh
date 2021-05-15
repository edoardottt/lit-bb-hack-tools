#!/bin/bash

RED='\033[0;31m'
NC='\033[0m' # No Color

while read line
do
    output=$( dig "$line" | grep CNAME )
    if [ ! -z "$output" ]
    then echo -e "${RED}$line${NC}"; dig "$line" | grep CNAME; echo ""
    echo "---------------------------"
    fi
done
