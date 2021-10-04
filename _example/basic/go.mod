module main

go 1.17

replace github.com/goware/firewall => ../../

require (
	github.com/go-chi/chi/v5 v5.0.4
	github.com/goware/firewall v0.0.0
)

require github.com/libp2p/go-cidranger v1.1.0 // indirect
