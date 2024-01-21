package main

import (
	"api/rate_limiter"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/", func(e echo.Context) error {
		user := e.RealIP()
		err := rate_limiter.For(user)
		if err != nil {
			e.Response().Header().Set("X-RateLimit-Limit", "5")
			return echo.NewHTTPError(http.StatusTooManyRequests, "You have reached threshold")
		}
		return e.String(200, "You can use the api!")
	})
	app.Logger.Fatal(app.Start(":3001"))
}
