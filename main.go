package main

import (
	"Distributed-cache/cache"
	"Distributed-cache/http"
	"Distributed-cache/tcp"
	"flag"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)

	mode := flag.String("mode", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is ", *mode)
	cache := cache.New(*mode)
	//inmemoryCache := cache.New("inmemory")
	go tcp.New(cache).Listen()
	http.New(cache).Listen()
}
