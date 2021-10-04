package firewall

import (
	_ "embed"
	"encoding/json"
)

// CloudProviderBlockList returns a slice of IP Ranges of aws, azure and gcp
func CloudProviderBlockList() (blockList []string) {
	var (
		awsRanges   awsRanges
		azureRanges azureRanges
		gcpRanges   gcpRanges
	)

	// aws
	json.Unmarshal(awsFile, &awsRanges)

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
	return blockList
}

var (
	//go:embed cloud-provider-data/aws-ip-ranges.json
	awsFile []byte
	//go:embed cloud-provider-data/azure-ip-ranges.json
	azureFile []byte
	//go:embed cloud-provider-data/gcp-ip-ranges.json
	gcpFile []byte
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
