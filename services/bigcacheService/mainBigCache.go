package bigcacheService

import (
	"fmt"
	"log"

	singleton "github.com/transybao1393/go-bigcache/singleton"
)

var (
	shards      int
	lifeWindow  int8
	cleanWindow int8
	verbose     bool
)

type BigCacheStruct struct {
	name  string
	value string
}

/*
*
*	MAIN FUNCTIONS
*
 */
func GetCacheData(key string) string {
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

func SetCacheData(name string, value string) {
	//- load env file
	// env, err := LoadConfig("..")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }
	// fmt.Println("config", reflect.TypeOf(env))
	// fmt.Println("type of verbose", reflect.TypeOf(env.Verbose))

	//- data validation

	//- add data to cache
	addToCache(name, value)
}

/*
*
*	SERVICE FUNCTIONS
*
 */
func addToCache(cacheName string, cacheValue string) {

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

func dataValidation() {
	// singleton.GetBigCacheInstance()
}

func showAllKeys() {}

func showByKey(key string) {}
