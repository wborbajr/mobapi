package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wborbajr/mobapi/handler/book"
)

func SetupRoutes(app *fiber.App) {

	//Middleware
	api := app.Group("/api")

	// Routes to V1
	v1 := api.Group("/v1/book")
	v1.Get("/", book.GetBooks)
	v1.Get("/:id", book.GetBook)
	v1.Delete("/:id", book.DeleteBook)
	v1.Post("/", book.NewBook)
	v1.Put("/:id", book.UpdateBook)

}
