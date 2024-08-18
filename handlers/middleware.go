package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func WithAuthenticatedUser(c *fiber.Ctx) error {
	fmt.Println("this is getting called")
	return c.Next()
}
