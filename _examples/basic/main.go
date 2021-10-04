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
	allowList, err := firewall.NewIPList([]string{"192.168.0.1/32", "::1/32"})
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
