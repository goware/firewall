package firewall

import (
	"net"

	"github.com/libp2p/go-cidranger"
)

// IPList inherits from cidranger.Ranger
// credits github.com/libp2p/go-cidranger
type IPList struct {
	cidranger.Ranger
}

// NewIPList returns a new IPList with inserted CIDR Ranges
func NewIPList(IPBlocks []string) (*IPList, error) {
	bl := &IPList{cidranger.NewPCTrieRanger()}
	err := bl.appendIPBlocks(IPBlocks)
	if err != nil {
		return nil, err
	}

	return bl, nil
}

// AppendIPBlocks Appends more CIDR Ranges to the IPList Struct
func (bl *IPList) AppendIPBlocks(IPBlocks []string) error {
	return bl.appendIPBlocks(IPBlocks)
}

func (bl *IPList) appendIPBlocks(IPBlocks []string) error {
	var (
		err   error
		IPNet *net.IPNet
	)
	for _, IPBlock := range IPBlocks {
		if IPBlock != "" {
			_, IPNet, err = net.ParseCIDR(IPBlock)
			if err != nil {
				return err
			}
			bl.Insert(cidranger.NewBasicRangerEntry(*IPNet))
		}
	}
	return nil
}
