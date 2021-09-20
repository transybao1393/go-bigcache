package services

type IBigCache interface {
	AddToCache(cacheName string, cacheValue string)
	// DataValidation(validateData string)
}
