package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(redisClient *redis.Client) *RedisClient {
	return &RedisClient{Client: redisClient}
}

func New(addr string) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisClient{Client: rdb}, nil

}
