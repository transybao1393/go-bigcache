package bigcacheService

type BigCacheStringStore interface {
	CreateOne() bool
	CreateMany(keys []string) bool
	FindAll() map[string]string
	FindOne(key string) map[string]string
	UpdateOne(key string) map[string]string
	DeleteAll() bool
	DeleteOne(key string)
}
