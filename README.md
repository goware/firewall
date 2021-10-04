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
   r := chi.NewRouter()
   blockList, err := firewall.NewIPList(firewall.CloudProviderBlockList())
   if err != nil {
      panic(err.Error())
   }
   err = blockList.AppendIPBlocks([]string{"127.0.0.0/1", "::1/128"})
   if err != nil {
      panic(err.Error())
   }
   // make allowList with ip addr in cidr notation,
   // so we can insert ip ranges and ip addr
   // refer https://whatismyipaddress.com/cidr
   allowList, err := firewall.NewIPList([]string{"192.168.0.1/32"})
   if err != nil {
      panic(err.Error())
   }
   fwBlockOverride := func(r *http.Request) bool {
      if r.Header.Get("internal") == "true" {
         return true
      }
      return false
   }
   r.Use(firewall.Firewall(allowList, blockList, fwBlockOverride))
   r.Use(middleware.Logger)
   r.Get("/", func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("welcome"))
   })
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