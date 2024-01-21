package main

import (
	"api/rate_limiter"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var time_taken int64

func main() {
	app := echo.New()

	app.GET("/", func(e echo.Context) error {
		user := e.RealIP()
		request_time := time.Now().UnixMicro()

		err := rate_limiter.For(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusTooManyRequests, "You have reached threshold")
		}

		time_taken = time.Now().UnixMicro() - request_time

		return e.String(200, "Time taken by redis "+strconv.FormatInt(time_taken, 10)+"!")
	})
	app.Logger.Fatal(app.Start(":3001"))
}
