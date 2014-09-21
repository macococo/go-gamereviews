package modules

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/macococo/go-gamereviews/conf"
	"github.com/macococo/go-gamereviews/utils"
	"log"
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
	utils.HandleError(err)

	if item == nil {
		if conf.IsDev() {
			log.Println("cache not hit:", key)
		}
		return nil
	}

	if conf.IsDev() {
		log.Println("cache hit:", key)
	}

	return item.Value
}

func (this *Cache) Put(key string, value []byte) {
	item := memcache.Item{Key: key, Value: value}
	this.client.Set(&item)

	if conf.IsDev() {
		log.Println("cache put:", key)
	}

}
