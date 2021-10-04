# firewall

Middleware for Blocking IP ranges by inserting CIDR Blocks and searching IPs through those blocks.

## Features

 - Easy to use
 - Efficient and Fast
 - Convenient Default option Blocks Major Cloud Providers


## Usage

See the full [Example](_example/basic/main.go)
```go
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/goware/firewall"
)

func main() {
	// Create New Router
	r := chi.NewRouter()
	// Create Block list
	// firewall.CloudProviderBlockList() returns a list of string of ip ranges of 
	// gcp, aws, azure
	blockList, _ := firewall.NewIPBlockList(firewall.CloudProviderBlockList())
	// Add more IP range Blocks to the list
	blockList.AppendIPBlocks([]string{"127.0.0.0/1", "::1/128"})
	// Create an allowList
	// if an ip range is in the blocklist ranges, but is inside allowlist
	// then the request is served
	// This is usefull to unblock your own hosted services
	allowList, _ := firewall.NewIPAllowList([]string{"192.168.0.1"})
	// fwBlockOverride is a function that is called if 
	// an ip is inside the blocklist, and is not in allowlist
	// this function returns a bool
	// if its true, then the client is approved and served
	fwBlockOverride := func(r *http.Request) bool {
		if r.Header.Get("internal") == "true" {
			return true
		}
		return false
	}
	// Add firewall
	r.Use(firewall.Firewall(allowList, blockList, fwBlockOverride))
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	// Serve
	http.ListenAndServe(":3000", r)
}

```

> Friendly Tip we get ip address of clients by parsing the list of X-FORWARDED-FOR header, so that we can avoid proxy
> addresses, to learn more visit: [CloudFlare Real IP](https://support.cloudflare.com/hc/en-us/articles/206776727-Understanding-the-True-Client-IP-Header)
> Also Read: [Blog]( https://husobee.github.io/golang/ip-address/2015/12/17/remote-ip-go.html)

## Credits

 - [go-cidranger](https://github.com/libp2p/go-cidranger)
   This middleware is based on this implementation of storing ip ranges in a data structre
   It makes it very efficient to store ip ranges and check if an ip is in one of those ranges
 
## LICENSE
[MIT](LICENSE)