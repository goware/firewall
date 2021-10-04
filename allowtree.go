package firewall

import (
	"bytes"
	"fmt"
	"net"
)

// AllowIPTree is a data structure for storing ip addresses and searching them efficiently
type AllowIPTree struct {
	leftNode  *AllowIPTree
	rightNode *AllowIPTree

	entry net.IP
}
// compareEntries compares two ip addresses and inserts them
func compareEntries(entry1 net.IP, entry2 net.IP) int {
	return bytes.Compare([]byte(entry1.String()), []byte(entry2.String()))
}

// insert inserts a net.IP object inside the AllowIPTree struct
func (a *AllowIPTree) insert(entry net.IP) error {
	if entry == nil {
		return fmt.Errorf("ip cannot be nil")
	}
	if compareEntries(entry, a.entry) > 0 {
		// move right
		if a.rightNode == nil {
			a.rightNode = &AllowIPTree{entry: entry}
		} else {
			a.rightNode.insert(entry)
		}
	} else {
		// move left
		if a.leftNode == nil {
			a.leftNode = &AllowIPTree{entry: entry}
		} else {
			a.leftNode.insert(entry)
		}
	}
	return nil
}

// Search searches for an IP object inside the AllowIPTree struct
func (a *AllowIPTree) Search(IP net.IP) (inTree bool) {
	if compareEntries(IP, a.entry) > 0 {
		// move right
		if a.rightNode == nil {
			return false
		} else if a.rightNode.entry.To4().String() == IP.To4().String() {
			return true
		} else {
			a.rightNode.Search(IP)
		}
	} else {
		// move left
		if a.leftNode == nil {
			return false
		} else if a.leftNode.entry.To4().String() == IP.To4().String() {
			return true
		} else {
			a.leftNode.Search(IP)
		}
	}
	return
}

func (a *AllowIPTree) parseAllowIPAddress(IPAddresses []string) error {
	var (
		err error
		ip  net.IP
	)
	for _, IPAddr := range IPAddresses {
		ip = net.ParseIP(IPAddr)
		if ip == nil {
			// throw error
			return fmt.Errorf("ip cannot be nil")
		}
		if a.entry == nil {
			a.entry = ip
		} else {
			err = a.insert(ip)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// AppendIPList appends a slice of ip strings to the AllowIPTree struct
func (a *AllowIPTree) AppendIPList(IPs []string) error {
	return a.parseAllowIPAddress(IPs)
}

// NewIPAllowList returns a new AllowIPTree struct with already inserted
// slice of IPs.
func NewIPAllowList(IPs []string) (*AllowIPTree, error) {
	var allowList AllowIPTree
	err := allowList.AppendIPList(IPs)
	if err != nil {
		return nil, err
	}
	return &allowList, nil
}
