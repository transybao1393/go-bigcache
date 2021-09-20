package controller

type IBigCacheController interface {
	GetCacheData(key string) string
	SetCacheData(name string, value string)
}
