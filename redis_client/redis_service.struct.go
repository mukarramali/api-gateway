package redis_client

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Client *redis.Client
	Ctx    context.Context
}
