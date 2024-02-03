package main

import (
	middlewares "api/middlewares/rate_limiter"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Use(middlewares.RateLimiterMiddleware)

	app.GET("/", func(e echo.Context) error {
		return e.String(200, "Welcome!")
	})
	app.Logger.Fatal(app.Start(":3001"))
}
