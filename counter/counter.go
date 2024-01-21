package counter

import (
	"api/redis_client"
	"strconv"
)

func Count(ip string) string {
	rs := redis_client.GetRedisService()
	counter := rs.Client.Incr(rs.Ctx, ip)
	return strconv.FormatInt(counter.Val(), 10)
}
