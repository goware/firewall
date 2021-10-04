package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/goware/firewall"
)

func main() {
	r := chi.NewRouter()
	blockList, _ := firewall.NewIPBlockList(firewall.CloudProviderBlockList())
	blockList.AppendIPBlocks([]string{"127.0.0.0/1", "::1/128"})
	allowList, _ := firewall.NewIPAllowList([]string{"192.168.0.1"})
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
