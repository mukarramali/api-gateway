package middlewares

import (
	"api/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var time_taken int64

func RateLimiterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		user := e.RealIP()
		request_time := time.Now().UnixMicro()

		err := LimitWithRedisSet(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusTooManyRequests, "You have reached threshold")
		}

		time_taken = time.Now().UnixMicro() - request_time
		fmt.Println("Time taken:" + utils.ToStr(time_taken))

		return next(e)
	}
}
