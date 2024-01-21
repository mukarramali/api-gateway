package rate_limiter

import (
	"api/redis_client"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const RedisKey = "user-requests:ip"
const Threshold = 5                    // requests per minute
const WindowInMilliseconds = 10 * 1000 // 1 min

func str(num int64) string {
	return strconv.FormatInt(num, 10)
}

func For(ip string) error {
	user := RedisKey + ip

	current_time := time.Now().UnixMilli()
	window_offset := current_time - WindowInMilliseconds

	addRequestFrom(user)
	resetThresholdFor(user)

	req_in_window := count(user, window_offset, current_time)

	if req_in_window > Threshold {
		err := fmt.Errorf("%d: Threshold crossed for ip %s", 400, ip)
		return errors.New(err.Error())
	}

	return nil
}

func count(user string, from int64, to int64) int64 {
	rs := redis_client.GetRedisService()
	count, _ := rs.Client.ZCount(rs.Ctx, user, str(from), str(to)).Result()
	return count
}

func resetThresholdFor(user string) {
	rs := redis_client.GetRedisService()

	current_time := time.Now().UnixMilli()
	window_offset := current_time - WindowInMilliseconds

	rs.Client.ZRemRangeByScore(rs.Ctx, user, "-inf", str(window_offset))
}

func addRequestFrom(user string) {
	rs := redis_client.GetRedisService()
	current_time := time.Now().UnixMilli()
	rs.Client.ZAdd(
		rs.Ctx,
		user,
		redis.Z{
			Score:  float64(current_time),
			Member: current_time,
		},
	).Result()
}
