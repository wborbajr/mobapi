package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wborbajr/mobapi/handler/book"
)

func SetupRoutes(app *fiber.App) {

	// app.Get("/api/v1/ping/", ping.Ping)

	// app.Get("/api/v1/book/", book.GetBooks)
	// app.Get("/api/v1/book/:id", book.GetBook)
	// app.Delete("/api/v1/book/:id", book.DeleteBook)
	// app.Post("/api/v1/book/", book.NewBook)
	// app.Put("/api/v1/book/:id", book.UpdateBook)

	//Middleware
	api := app.Group("/api", logger.New())

	// Routes to V1
	v1 := api.Group("/v1/book")
	v1.Get("/", book.GetBooks)
	v1.Get("/:id", book.GetBook)
	v1.Delete("/:id", book.DeleteBook)
	v1.Post("/", book.NewBook)
	v1.Put("/:id", book.UpdateBook)

}
