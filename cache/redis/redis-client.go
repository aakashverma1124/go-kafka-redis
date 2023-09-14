package redis

import (
	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func GetClient(address string) *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "",
			DB:       0,
		})
	}
	return client
}
