package server

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func NewCache(addr, password string, db int) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &Cache{
		client: rdb,
	}
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) error {
	return c.client.Set(context.Background(), key, value, duration).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return c.client.Get(context.Background(), key).Result()
}
