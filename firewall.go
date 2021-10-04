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
