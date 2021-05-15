#!/bin/bash

RED='\033[0;31m' # Red
NC='\033[0m' # No Color

# Read line on stdin
while read line
do
    output=$( dig "$line" | grep CNAME )
    # if output is not empty
    if [ ! -z "$output" ]
    # print the CNAME record from dig
    then echo -e "${RED}$line${NC}"; dig "$line" | grep CNAME; echo ""
    echo "---------------------------"
    fi
done
