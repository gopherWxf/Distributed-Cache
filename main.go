package main

import (
	"Distributed-cache/cache"
	"Distributed-cache/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
