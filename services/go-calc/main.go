package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jansvabik/envoy-proxy/tree/master/services/go-calc/calc"
)

func main() {
	router := fiber.New()
	router.Get("/", calc.HealthCheck)
	router.Get("/add", calc.Add)
	router.Get("/subtract", calc.Subtract)
	router.Get("/multiply", calc.Multiply)
	router.Get("/divide", calc.Divide)
	router.Listen(":80")
}
