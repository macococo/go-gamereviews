package utils

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	AppCache *Cache
)

type Cache struct {
	client *memcache.Client
}

func init() {
	AppCache = &Cache{client: memcache.New("127.0.0.1:11211")}
}

func (this *Cache) Get(key string) []byte {
	item, err := this.client.Get(key)
	HandleError(err)

	if item == nil {
		return nil
	}
	return item.Value
}

func (this *Cache) Put(key string, value []byte) {
	item := memcache.Item{Key: key, Value: value}
	this.client.Set(&item)
}
