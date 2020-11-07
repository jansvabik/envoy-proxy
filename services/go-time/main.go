package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jansvabik/envoy-proxy/tree/master/services/go-time/handlers"
)

func main() {
	router := fiber.New()
	router.Get("/", handlers.GetCurrentTime)
	router.Get("/:country", handlers.GetCurrentTimeFormatted)
	router.Listen(":80")
}
