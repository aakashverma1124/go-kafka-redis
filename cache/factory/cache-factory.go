package factory

import (
	"go-kafka-redis/cache"
	"go-kafka-redis/cache/redis"
	"go-kafka-redis/constant"
)

type CacheFactory struct{}

func (cacheFactory *CacheFactory) GetCache(address, cacheType string) cache.Cache {
	if cacheType == constant.REDIS_CACHE {
		client := redis.GetClient(address)
		return redis.NewCache(client)
	}
	return nil
}
