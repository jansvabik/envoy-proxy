package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jansvabik/envoy-proxy/tree/master/services/go-time/handlers"
)

func main() {
	app := fiber.New()
	app.Get("/", handlers.GetCurrentTime)
	app.Get("/:country", handlers.GetCurrentTimeFormatted)
	app.Listen(":9001")
}
