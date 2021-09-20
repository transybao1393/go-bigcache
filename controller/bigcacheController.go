package controller

import (
	"fmt"
	"log"

	"github.com/transybao1393/go-bigcache/singleton"
)

var (
	bigCacheInstance = singleton.GetBigCacheInstance()
)

type bigCacheControllerStruct struct{}

func BigCacheControllerInit() IBigCacheController {
	return &bigCacheControllerStruct{}
}

func (bc *bigCacheControllerStruct) GetCacheData(key string) string {
	//- init
	cache := singleton.GetBigCacheInstance()

	entry, err := cache.Get(key)
	if err != nil {
		//- error
		// log.Fatal(err)
		// panic(err)
	}
	//- else
	fmt.Println("cache data", string(entry))
	return string(entry)
}

func (bc *bigCacheControllerStruct) SetCacheData(name string, value string) {
	//- add data to cache
	addToCache(name, value)
	// s.AddToCache(name, value)
}

//- service function
func addToCache(cacheName string, cacheValue string) {
	//- singleton

	//- check if exist => update existing one
	//- else => create new
	bigCacheInstance.Append(cacheName, []byte(cacheValue))

	entry, err := bigCacheInstance.Get(cacheName)
	if err != nil {
		//- error
		log.Fatal(err)
	}
	//- else
	fmt.Println("cache data", string(entry))
}
