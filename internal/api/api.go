package api

import (
	"fmt"
	"football-api/internal/database"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func ConnectRedis() (*RedisClient, error) {
	rc, err := database.New("localhost:6379")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &RedisClient{Client: rc.Client}, err
}

type Api struct {
	Rdb *RedisClient
}

func NewApi(redisClient *redis.Client) *Api {
	return &Api{
		Rdb: &RedisClient{Client: redisClient},
	}
}
