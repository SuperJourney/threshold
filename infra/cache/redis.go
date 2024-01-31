package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(redisOption *redis.Options) (*RedisCache, error) {
	client := redis.NewClient(redisOption)

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisCache{
		client: client,
	}, nil
}

func (c *RedisCache) Exist(ctx context.Context, key []byte) (bool, error) {
	result, err := c.client.Exists(string(key)).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func (c *RedisCache) Get(ctx context.Context, key []byte) ([]byte, error) {
	result, err := c.client.Get(string(key)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // 返回nil表示缓存中没有该键值对
		}
		return nil, err
	}
	return result, nil
}

func (c *RedisCache) Set(ctx context.Context, key, value []byte, expired int64) error {
	err := c.client.Set(string(key), value, 0).Err()
	if err != nil {
		return err
	}
	if expired > 0 {
		err = c.client.Expire(string(key), time.Duration(expired)*time.Second).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *RedisCache) Incr(ctx context.Context, key []byte) error {
	err := c.client.Incr(string(key)).Err()
	if err != nil {
		return err
	}
	return nil
}
