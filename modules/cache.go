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
	server := conf.Config.MemcachedServer
	log.Println("memcached:", server)

	AppCache = &Cache{client: memcache.New(server)}
}

func (this *Cache) Get(key string) []byte {
	item, err := this.client.Get(key)
	if conf.IsDev() {
		utils.HandleError(err)
	}

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

func (this *Cache) Delete(key string) {
	this.client.Delete(key)

	if conf.IsDev() {
		log.Println("cache delete:", key)
	}
}

func (this *Cache) Increment(key string, delta uint64) uint64 {
	newValue, err := this.client.Increment(key, delta)
	utils.HandleError(err)

	if conf.IsDev() {
		log.Println("cache increment:", key, delta)
	}

	return newValue
}
