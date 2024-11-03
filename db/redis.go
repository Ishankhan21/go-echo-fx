package db

import (
	"github.com/redis/go-redis/v9"
)

func GetRedisconnection() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}
