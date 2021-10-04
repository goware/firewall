#!/bin/bash

green=`tput setaf 2`

echo $green

LINKS="https://www.gstatic.com/ipranges/goog.json gcp-ip-ranges.json"

cd cloudProviderData

# aws
echo "Downloading AWS Ranges..."
curl -s -L -o "aws-ip-ranges.json" "https://ip-ranges.amazonaws.com/ip-ranges.json"

# gcp
echo "Downloading GCP Ranges..."
curl -s -L -o "gcp-ip-ranges.json" "https://www.gstatic.com/ipranges/goog.json"

# Azure portion

AZURE_FILE_LINK="$(curl https://www.microsoft.com/en-us/download/confirmation.aspx\?id\=56519 | grep -Eoi '<a [^>]+>' | grep -Eo 'href="[^\"]+"' | grep "download.microsoft.com/download/" | grep -m 1 -Eo '(http|https)://[^"]+'
)"

echo "Downloading Azure Ranges..."
curl -s -L -o azure-original-ranges.json $AZURE_FILE_LINK

echo "Converting Azure Ranges to new format..."
go run ./main.go


echo "done"

echo "Cleaning downloaded files"

rm azure-original-ranges.json
