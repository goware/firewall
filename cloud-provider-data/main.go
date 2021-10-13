package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type azureOriginalRanges struct {
	Values []azureValues `json:"values"`
}

type azureValues struct {
	Properties azureProperties `json:"properties"`
}

type azureProperties struct {
	AddressPrefixes []string `json:"addressPrefixes"`
}

type azureModifiedRanges struct {
	Prefixes []azurePrefix `json:"prefixes"`
}

type azurePrefix struct {
	IPPrefix string `json:"ip_prefix"`
}

func main() {
	var (
		azureData azureOriginalRanges
		out       azureModifiedRanges
	)

	// TODO: change, why is this hardcoded..?
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
			out.Prefixes = append(out.Prefixes, azurePrefix{address})
		}
	}
	outfile, _ := json.MarshalIndent(out, "", " ")
	_ = ioutil.WriteFile("azure-ip-ranges.json", outfile, 0644)
}
