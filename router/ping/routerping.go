package ping

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wborbajr/mobapi/handler/ping"
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

	// alive
	v1 := api.Group("/v1/ping")
	v1.Get("/", ping.Ping)

}
