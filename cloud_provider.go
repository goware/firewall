package firewall

import (
	_ "embed"
	"encoding/json"
	"github.com/gocarina/gocsv"
)

var cloudProviderBlockList []string = nil

// CloudProviderBlockList returns a slice of IP Ranges of aws, azure and gcp
func CloudProviderBlockList() (blockList []string) {
	if cloudProviderBlockList != nil {
		// return memoized list
		return cloudProviderBlockList
	}

	var (
		awsRanges   awsRanges
		azureRanges azureRanges
		gcpRanges   gcpRanges
	)

	// aws
	err := json.Unmarshal(awsFile, &awsRanges)
	if err != nil {
		panic(err)
	}

	for _, prefix := range awsRanges.Prefixes {
		if prefix.IPPrefix != "" {
			blockList = append(blockList, prefix.IPPrefix)
		} else if prefix.IPv6Prefix != "" {
			blockList = append(blockList, prefix.IPv6Prefix)
		}
	}

	// azure
	json.Unmarshal(azureFile, &azureRanges)
	for _, prefix := range azureRanges.Prefixes {
		blockList = append(blockList, prefix.IPPrefix)
	}

	// gcp
	json.Unmarshal(gcpFile, &gcpRanges)
	for _, prefix := range gcpRanges.Prefixes {
		if prefix.IPPrefix != "" {
			blockList = append(blockList, prefix.IPPrefix)
		} else if prefix.IPv6Prefix != "" {
			blockList = append(blockList, prefix.IPv6Prefix)
		}
	}

	// linode
	var linodeRanges []*csvIPPrefix
	gocsv.UnmarshalBytes(linodeFile, &linodeRanges)
	for _, prefix := range linodeRanges {
		if prefix.IPPrefix != "" {
			blockList = append(blockList, prefix.IPPrefix)
		}
	}

	// digital ocean
	var digitalOceanRanges []*csvIPPrefix
	gocsv.UnmarshalBytes(digitalOceanFile, &digitalOceanRanges)
	for _, prefix := range digitalOceanRanges {
		if prefix.IPPrefix != "" {
			blockList = append(blockList, prefix.IPPrefix)
		}
	}

	// memoize
	cloudProviderBlockList = blockList

	return blockList
}

var (
	//go:embed cloud-provider-data/aws-ip-ranges.json
	awsFile []byte
	//go:embed cloud-provider-data/azure-ip-ranges.json
	azureFile []byte
	//go:embed cloud-provider-data/gcp-ip-ranges.json
	gcpFile []byte
	//go:embed cloud-provider-data/linode-ranges.csv
	linodeFile []byte
	//go:embed cloud-provider-data/digital-ocean-ranges.csv
	digitalOceanFile []byte
)

type awsRanges struct {
	Prefixes []awsPrefix `json:"prefixes"`
}

type awsPrefix struct {
	IPPrefix   string `json:"ip_prefix"`
	IPv6Prefix string `json:"ipv6_prefix"`
}

type azureRanges struct {
	Prefixes []azurePrefix `json:"prefixes"`
}

type azurePrefix struct {
	IPPrefix string `json:"ip_prefix"`
}

type gcpRanges struct {
	Prefixes []gcpPrefix `json:"prefixes"`
}

type gcpPrefix struct {
	IPPrefix   string `json:"ipv4Prefix"`
	IPv6Prefix string `json:"ipv6Prefix"`
}

type csvIPPrefix struct {
	IPPrefix string `csv:"prefix"`
}
