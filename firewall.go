package firewall

import (
	"net"
	"net/http"
	"time"
)

func Firewall(allowList *AllowIPTree, blockList *BlockList, fwBlockOverride func(r *http.Request) bool) func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// getting ip address without port
			ipAddr := ipAddrWithoutPort(r)
			// converts ip addr string to net.IP
			ip := net.ParseIP(ipAddr)
			if inBlockList, _ := blockList.Contains(ip); inBlockList {
				// check if ip is not in allowList and is not being
				// overriden by fwBlockOverride
				if !allowList.Search(ip) && !fwBlockOverride(r) {
					// ip is blocked
					time.Sleep(29 * time.Second)
					w.WriteHeader(403)
					return
				}
			}
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	return f
}
