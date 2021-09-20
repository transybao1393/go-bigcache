package bigcacheService

import (
	singleton "github.com/transybao1393/go-bigcache/singleton"
)

//- global variable
var (
	cache = singleton.GetBigCacheInstance()
)

type BigCacheStringStore interface {
	CreateOne() bool
	CreateMany(keys []string) bool
	FindAll() map[string]string
	FindOne(key string) map[string]string
	UpdateOne(key string) map[string]string
	DeleteAll() bool
	DeleteOne(key string)
}

/*
*
*	MAIN FUNCTIONS
*
 */
func CreateOne(key string, name string) (bool, string) {
	error := cache.Append(key, []byte(name))
	if error != nil {
		return false, "Cannot create one"
		// log.Fatal("Cannot create one with error", error)
	}
	return true, "Success"
}
func CreateMany() {}
func FindAll() {
	cache.Iterator()
}
func FindOne(key string) string {
	entry, err := cache.Get(key)
	if err != nil {
		return ""
	}
	return string(entry)
}

func UpdateOne(key string, updatedString string) (bool, string) {
	err := cache.Set(key, []byte(updatedString))
	if err != nil {
		return false, err.Error()
	}
	return true, "Success"
}

func DeleteAll(key string) (bool, string) {
	//- iterate to all and delete every single key
	err := cache.Reset()
	if err != nil {
		return false, err.Error()
	}
	return true, "Success"
}

func DeleteOne(key string) (bool, string) {
	err := cache.Delete(key)
	if err != nil {
		return false, err.Error()
	}
	return true, "Success"
}
