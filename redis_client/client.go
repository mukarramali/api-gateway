package redis_client

import (
	"github.com/redis/go-redis/v9"
)

func GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}
