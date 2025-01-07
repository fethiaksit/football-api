package api

import (
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

type Api struct {
	Rdb *RedisClient
}

func NewApi(redisClient *redis.Client) *Api {
	return &Api{
		Rdb: &RedisClient{Client: redisClient},
	}
}
