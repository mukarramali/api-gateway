package counter

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func Count(store *redis.Client,
	ctx context.Context,
	ip string) string {
	counter := store.Incr(ctx, ip)
	return strconv.FormatInt(counter.Val(), 10)
}
