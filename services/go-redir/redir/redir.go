package redir

import (
	"github.com/gofiber/fiber/v2"
)

// Response is a http response JSON structure
type Response struct {
	Status           string   `json:"status"`
	Msg              string   `json:"msg"`
	AllowedProtocols []string `json:"allowedProtocols,omitempty"`
}

// Redir redirects user to the specified URL
func Redir(c *fiber.Ctx) error {
	// target validation
	target := c.Query("target")
	if len(target) == 0 {
		return c.JSON(Response{
			Status: "ERROR",
			Msg:    "No target URL was specified. Add ?target=some.url&protocol=https",
		})
	}

	// protocol validation
	protocol := c.Query("protocol")
	if len(protocol) == 0 {
		protocol = "http"
	}

	// test that the protocol is allowed
	allowedProtocols := []string{"http", "https", "ftp"}
	for _, ap := range allowedProtocols {
		if protocol == ap {
			return c.Redirect(protocol + "://" + target)
		}
	}

	// non-allowed protocol was used
	return c.JSON(Response{
		Status:           "ERROR",
		Msg:              "Only protocols in the \"allowedProtocols\" field are allowed.",
		AllowedProtocols: allowedProtocols,
	})
}
