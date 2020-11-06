package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jansvabik/envoy-proxy/tree/master/services/go-redir/redir"
)

func main() {
	app := fiber.New()
	app.Get("/", redir.Redir)
	app.Listen(":80")
}
