package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-kafka-redis/cache"
	"time"
)

type Cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) cache.Cache {
	return &Cache{
		client: client,
	}
}

func (c *Cache) GetKey(key string) (interface{}, error) {
	ctx := context.Background()
	return c.client.Get(ctx, key).Result()
}

func (c *Cache) SetKey(key string, value interface{}, ttl int) error {
	ctx := context.Background()
	return c.client.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
}

func (c *Cache) PutKey(key string, value interface{}) error {
	ctx := context.Background()
	return c.client.Set(ctx, key, value, 0).Err()
}

func (c *Cache) DelKey(key string) error {
	ctx := context.Background()
	return c.client.Del(ctx, key).Err()
}
