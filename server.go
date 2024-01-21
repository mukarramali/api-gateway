package main

import (
	"api/counter"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/", func(e echo.Context) error {
		counter := counter.Count(e.RealIP())
		return e.String(200, "Counter: "+counter)
	})
	app.Logger.Fatal(app.Start(":3001"))
}
