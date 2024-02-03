package middlewares

import (
	"api/redis_client"
	"api/utils"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const RedisKey = "user-requests:ip"
const Threshold = 30                   // requests per minute
const WindowInMilliseconds = 30 * 1000 // 1 min

func For(ip string) error {
	user := RedisKey + ip

	current_time := time.Now().UnixMilli()
	window_offset := current_time - WindowInMilliseconds

	go addRequestFrom(user)
	go resetThresholdFor(user)

	req_in_window := count(user, window_offset, current_time)

	if req_in_window > Threshold {
		err := fmt.Errorf("%d: Threshold crossed for ip %s", 400, ip)
		return errors.New(err.Error())
	}

	return nil
}

func count(user string, from int64, to int64) int64 {
	rs := redis_client.GetRedisService()
	count, _ := rs.Client.ZCount(rs.Ctx, user, utils.ToStr(from), utils.ToStr(to)).Result()
	return count
}

func resetThresholdFor(user string) {
	rs := redis_client.GetRedisService()

	current_time := time.Now().UnixMilli()
	window_offset := current_time - WindowInMilliseconds

	rs.Client.ZRemRangeByScore(rs.Ctx, user, "-inf", utils.ToStr(window_offset))
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
