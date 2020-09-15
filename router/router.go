package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handlerbook "github.com/wborbajr/mobapi/handler"
)

func SetupRoutes(app *fiber.App) {
	//Middleware
	api := app.Group("/api", logger.New())

	// Routes to V1
	v1 := api.Group("/v1/book")
	v1.Get("/", handlerbook.GetBooks)
	v1.Get("/:id", handlerbook.GetBook)
	v1.Delete("/:id", handlerbook.DeleteBook)
	v1.Post("/", handlerbook.NewBook)
	v1.Put("/:id", handlerbook.UpdateBook)

}
