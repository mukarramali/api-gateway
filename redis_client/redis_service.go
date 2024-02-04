package redis_client

import (
	"context"
	"fmt"
	"sync"
)

var instance *RedisService
var once sync.Once

func GetRedisService() *RedisService {
	once.Do(func() {
		instance = &RedisService{
			Client: GetClient(),
			Ctx:    context.Background(),
		}
		_, err := instance.Client.Ping(instance.Ctx).Result()
		if err != nil {
			fmt.Println("Error connecting to Redis:", err)
		} else {
			fmt.Println("Redis connected")
		}
	})

	return instance
}
