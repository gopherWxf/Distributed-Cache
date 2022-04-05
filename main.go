package main

import (
	"Distributed-cache/cache"
	"Distributed-cache/http"
	"Distributed-cache/tcp"
)

func main() {
	inmemoryCache := cache.New("inmemory")
	go tcp.New(inmemoryCache).Listen()
	http.New(inmemoryCache).Listen()
}
