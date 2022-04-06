package cache

import "log"

type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
}

func New(mode string) Cache {
	var cache Cache
	if mode == "inmemory" {
		cache = newInMemoryCache()
	}
	if mode == "pebble" {
		cache = newPebbleCache()
	}
	if cache == nil {
		panic("unknown cache type " + mode)
	}
	log.Println(mode, "ready to serve")
	return cache
}
