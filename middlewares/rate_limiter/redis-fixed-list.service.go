package middlewares

import (
	"api/redis_client"
	"api/utils"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisFixedListLimiter struct{}

/*
Use redis' sorted sets to find the number of records from offset time to now
*/
func (s *RedisFixedListLimiter) For(ip string) error {
	user := RedisKey + ip

	current_time := time.Now().UnixMilli()
	window_offset := current_time - WindowInMilliseconds

	go s.addRequestFrom(user)

	req_in_window := s.count(user, window_offset, current_time)

	if req_in_window > Threshold {
		err := fmt.Errorf("%d: Threshold crossed for ip %s", 400, ip)
		return errors.New(err.Error())
	}

	return nil
}

func (s *RedisFixedListLimiter) count(user string, from int64, to int64) int64 {
	rs := redis_client.GetRedisService()
	count, _ := rs.Client.ZCount(rs.Ctx, user, utils.ToStr(from), utils.ToStr(to)).Result()
	return count
}

func (s *RedisFixedListLimiter) addRequestFrom(user string) {
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
