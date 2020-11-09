package calc

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Add makes an addition of specified numbers
func Add(c *fiber.Ctx) error {
	// nums length validation
	nums := c.Query("nums")
	if len(nums) == 0 {
		return c.JSON(Response{
			Status: "ERROR",
			Msg:    "No number to addition was specified. Add ?nums=1,2,3,4,5,-10",
			Data:   false,
		})
	}

	// calculating the sum and testing each number
	list := strings.Split(nums, ",")
	r := 0
	for i, v := range list {
		n, err := strconv.Atoi(v)
		if err != nil {
			return c.JSON(Response{
				Status: "ERROR",
				Msg:    "Param No. " + strconv.Itoa(i) + " is not a number.",
				Data:   false,
			})
		}
		r += n
	}

	// ok response
	return c.JSON(Response{
		Status: "OK",
		Msg:    "The sum was successfully calculated.",
		Data:   r,
	})
}
