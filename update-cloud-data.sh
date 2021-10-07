#!/bin/bash

green=`tput setaf 2`

echo $green

LINKS="https://www.gstatic.com/ipranges/goog.json gcp-ip-ranges.json"

cd ./cloud-provider-data

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

# Linode
echo "Downloading Linode Ranges..."
curl -s -L -o linode-ranges-original.csv "https://geoip.linode.com/"
echo "prefix, country, subdivision, city, zipcode, allocation_size" > linode-ranges.csv
echo "$(sed -e '1,3d' < linode-ranges-original.csv)" >> linode-ranges.csv

# digital ocean
echo "Downloading Digital Ocean Ranges..."
echo "prefix, country, state, city, zipcode" > digital-ocean-ranges.csv
echo "$(curl -L https://digitalocean.com/geo/google.csv)" >> digital-ocean-ranges.csv

echo "done"

echo "Cleaning downloaded files"

rm azure-original-ranges.json
rm linode-ranges-original.csv