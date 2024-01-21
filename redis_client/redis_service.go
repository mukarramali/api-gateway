package redis_client

import (
	"context"
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
	})
	return instance
}
