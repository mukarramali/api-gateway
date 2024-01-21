package main

import (
	"api/counter"
	"api/redis_client"
	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	client := redis_client.GetClient()
	ctx := context.Background()

	app.GET("/", func(e echo.Context) error {
		counter := counter.Count(client, ctx, e.RealIP())
		return e.String(200, "Counter: "+counter)
	})
	app.Logger.Fatal(app.Start(":3001"))
}
