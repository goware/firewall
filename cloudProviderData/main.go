package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type azureIn struct {
	Values []values `json:"values"`
}

type values struct {
	Properties properties `json:"properties"`
}

type properties struct {
	AddressPrefixes []string `json:"addressPrefixes"`
}

type azureOut struct {
	Prefixes []prefix `json:"prefixes"`
}

type prefix struct {
	IPPrefix string `json:"ip_prefix"`
}

func main() {
	var (
		azureData azureIn
		out       azureOut
	)

	file, err := os.Open("azure-original-ranges.json")
	if err != nil {
		log.Print(err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print(err.Error())
	}
	err = json.Unmarshal(fileBytes, &azureData)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, value := range azureData.Values {
		for _, address := range value.Properties.AddressPrefixes {
			out.Prefixes = append(out.Prefixes, prefix{address})
		}
	}
	outfile, _ := json.MarshalIndent(out, "", " ")
	_ = ioutil.WriteFile("azure-ip-ranges.json", outfile, 0644)
}
