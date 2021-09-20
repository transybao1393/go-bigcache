package services

import (
	"fmt"
	"log"

	"github.com/transybao1393/go-bigcache/singleton"
)

type BigCacheService struct{}

func InitBigCacheService() IBigCache {
	return &BigCacheService{}
}

func (bcs *BigCacheService) AddToCache(cacheName string, cacheValue string) {
	//- singleton
	cache := singleton.GetBigCacheInstance()

	//- check if exist => update existing one
	//- else => create new
	cache.Append(cacheName, []byte(cacheValue))

	entry, err := cache.Get(cacheName)
	if err != nil {
		//- error
		log.Fatal(err)
	}
	//- else
	fmt.Println("cache data", string(entry))
}
