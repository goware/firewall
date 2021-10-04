package firewall

import (
	"net"

	"github.com/libp2p/go-cidranger"
)

// BlockList inherits from cidranger.Ranger
// credits github.com/libp2p/go-cidranger
type BlockList struct {
	cidranger.Ranger
}

// NewIPBlockList returns a new BlockList with inserted CIDR Ranges
func NewIPBlockList(IPBlocks []string) (*BlockList, error) {
	bl := &BlockList{cidranger.NewPCTrieRanger()}
	err := bl.appendIPBlocks(IPBlocks)
	if err != nil {
		return nil, err
	}

	return bl, nil
}

// AppendIPBlocks Appends more CIDR Ranges to the BlockList Struct
func (bl *BlockList) AppendIPBlocks(IPBlocks []string) error {
	return bl.appendIPBlocks(IPBlocks)
}

func (bl *BlockList) appendIPBlocks(IPBlocks []string) error {
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
