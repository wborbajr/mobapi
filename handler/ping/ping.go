package ping

import (
	"github.com/gofiber/fiber/v2"
)

type JSONResponse struct {
	Ping string `json:"ping"`
}

func Ping(c *fiber.Ctx) error {

	jsonResponse := JSONResponse{
		Ping: "pong",
	}

	return c.JSON(jsonResponse)
}
