package ping

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wborbajr/mobapi/handler/ping"
)

func SetupRoutes(app *fiber.App) {

	//Middleware
	api := app.Group("/api", logger.New())

	// alive
	v1 := api.Group("/v1/ping")
	v1.Get("/", ping.Ping)

}
