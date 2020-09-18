package main_test

import (
	"testing"

	"github.com/gofiber/fiber"
)

func TestServerUp(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *Ctx) error {
		return c.Next()
	})

}
