package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Response is a http response JSON structure
type Response struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

// GetCurrentTime gets the current time and creates an http response
func GetCurrentTime(c *fiber.Ctx) error {
	return c.JSON(Response{
		Status: "OK",
		Msg:    "Here you are, current time.",
		Data:   time.Now(),
	})
}

// GetCurrentTimeFormatted gets the current time, formats it and creates an http response
func GetCurrentTimeFormatted(c *fiber.Ctx) error {
	formats := map[string]string{
		"cs-CZ": "1. 2. 2006 15:04:05 (-07:00)",
		"en-GB": "01/02/06 15:04:05 (-07:00)",
		"en-US": "02/01/06 03:04:05PM (-07:00)",
		"fr-FR": "02/01/2006 15:04:05 (-07:00)",
	}

	// test that the country format exists
	format := c.Params("country")
	if _, ok := formats[format]; !ok {
		return c.JSON(Response{
			Status: "ERROR",
			Msg:    "We don't know this country/language format.",
		})
	}

	return c.JSON(Response{
		Status: "OK",
		Msg:    "Here you are, formatted current time.",
		Data:   time.Now().Format(formats[format]),
	})
}
