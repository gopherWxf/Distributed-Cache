package cache

import (
	"github.com/cockroachdb/pebble"
	"log"
)

type pebbleCache struct {
	db *pebble.DB
}

func (p *pebbleCache) Set(s string, bytes []byte) error {
	err := p.db.Set([]byte(s), bytes, pebble.Sync)
	return err
}

func (p *pebbleCache) Get(s string) ([]byte, error) {
	value, closer, err := p.db.Get([]byte(s))
	if err == nil {
		defer closer.Close()
	}
	return value, err
}

func (p *pebbleCache) Del(s string) error {
	err := p.db.Delete([]byte(s), pebble.Sync)
	return err
}

func (p *pebbleCache) GetStat() Stat {
	tables, err := p.db.SSTables(pebble.WithProperties())
	if err != nil {
		return Stat{}
	}
	var count, keySize, valueSize int64
	for sstIdx := range tables {
		for _, sstPro := range tables[sstIdx] {
			count += int64(sstPro.Properties.NumEntries)
			keySize += int64(sstPro.Properties.RawKeySize)
			valueSize += int64(sstPro.Properties.RawValueSize)
		}
	}
	return Stat{
		Count:     count,
		KeySize:   keySize,
		ValueSize: valueSize,
	}
}

func newPebbleCache() *pebbleCache {
	db, err := pebble.Open("pebbleDB", &pebble.Options{})
	if err != nil {
		log.Fatal(err)
	}
	return &pebbleCache{db: db}
}
