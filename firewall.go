package firewall

import (
	"github.com/libp2p/go-cidranger"
	"net"
	"net/http"
	"time"
)

// BlockList inherits from cidranger.Ranger
//credits github.com/libp2p/go-cidranger
type BlockList struct {
	cidranger.Ranger
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

// NewIPBlockList returns a new BlockList with inserted CIDR Ranges
func NewIPBlockList(IPBlocks []string) (*BlockList, error) {
	bl := BlockList{cidranger.NewPCTrieRanger()}
	err := bl.appendIPBlocks(IPBlocks)
	if err != nil {
		return nil, err
	}

	return &bl, nil
}

// AppendIPBlocks Appends more CIDR Ranges to the BlockList Struct
func (bl *BlockList) AppendIPBlocks(IPBlocks []string) error {
	return bl.appendIPBlocks(IPBlocks)
}

func Firewall(allowList *AllowIPTree, blockList *BlockList, fwBlockOverride func(r *http.Request) bool) func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// getting ip address without port
			ipAddr := ipAddrWithoutPort(r)
			// converts ip addr string to net.IP
			ip := net.ParseIP(ipAddr)
			if inBlockList, _ := blockList.Contains(ip); inBlockList {
				if allowList.Search(ip) {
					// ip in allowList -> Serve request
					goto SERVE
				}
				if fwBlockOverride(r) {
					// request is override through BlockOverride func -> serve request
					goto SERVE
				}
				// ip is blocked
				time.Sleep(29 * time.Second)
				w.WriteHeader(403)
				return
			}
		SERVE:
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	return f
}
