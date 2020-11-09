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
				Msg:    "Param No. " + strconv.Itoa(i+1) + " is not an integer.",
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

// Subtract makes an integer substraction of specified numbers
func Subtract(c *fiber.Ctx) error {
	// nums length validation
	nums := c.Query("nums")
	if len(nums) == 0 {
		return c.JSON(Response{
			Status: "ERROR",
			Msg:    "No number to subtraction was specified. Add ?nums=1,2,3,4,5,-10",
			Data:   false,
		})
	}

	// calculating the subtraction and testing each number
	list := strings.Split(nums, ",")
	r, err := strconv.Atoi(list[0])
	if err != nil {
		return c.JSON(Response{
			Status: "ERROR",
			Msg:    "Param No. 1 is not an integer.",
			Data:   false,
		})
	}
	for i, v := range list[1:] {
		n, err := strconv.Atoi(v)
		if err != nil {
			return c.JSON(Response{
				Status: "ERROR",
				Msg:    "Param No. " + strconv.Itoa(i+2) + " is not an integer.",
				Data:   false,
			})
		}
		r -= n
	}

	// ok response
	return c.JSON(Response{
		Status: "OK",
		Msg:    "The subtraction was successfully calculated.",
		Data:   r,
	})
}

// Multiply multiplies the specified numbers
func Multiply(c *fiber.Ctx) error {
	// nums length validation
	nums := c.Query("nums")
	if len(nums) == 0 {
		return c.JSON(Response{
			Status: "ERROR",
			Msg:    "No number to multiply was specified. Add ?nums=1,2,3,4,5,-10",
			Data:   false,
		})
	}

	// calculating the multiplication and testing each number
	list := strings.Split(nums, ",")
	r := 1
	for i, v := range list {
		n, err := strconv.Atoi(v)
		if err != nil {
			return c.JSON(Response{
				Status: "ERROR",
				Msg:    "Param No. " + strconv.Itoa(i+1) + " is not an integer.",
				Data:   false,
			})
		}
		r *= n
	}

	// ok response
	return c.JSON(Response{
		Status: "OK",
		Msg:    "The multiplication was successfully calculated.",
		Data:   r,
	})
}

// Divide makes an integer division of specified numbers
func Divide(c *fiber.Ctx) error {
	// nums length validation
	nums := c.Query("nums")
	if len(nums) == 0 {
		return c.JSON(Response{
			Status: "ERROR",
			Msg:    "No number to division was specified. Add ?nums=12,4,-1",
			Data:   false,
		})
	}

	// calculating the division and testing each number
	list := strings.Split(nums, ",")
	r, err := strconv.Atoi(list[0])
	if err != nil {
		return c.JSON(Response{
			Status: "ERROR",
			Msg:    "Param No. 1 is not an integer.",
			Data:   false,
		})
	}
	for i, v := range list[1:] {
		n, err := strconv.Atoi(v)
		if err != nil {
			return c.JSON(Response{
				Status: "ERROR",
				Msg:    "Param No. " + strconv.Itoa(i+2) + " is not an integer.",
				Data:   false,
			})
		}
		r /= n
	}

	// ok response
	return c.JSON(Response{
		Status: "OK",
		Msg:    "The division was successfully calculated.",
		Data:   r,
	})
}
