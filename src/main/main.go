package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/wborbajr/mobapi/src/book"
	"github.com/wborbajr/mobapi/src/database"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/book", book.GetBooks)
	v1.Get("/book/:id", book.GetBook)
	v1.Delete("/book/:id", book.DeleteBook)
	v1.Post("/book", book.NewBook)
	v1.Put("/book/:id", book.UpdateBook)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.sqlite")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database...")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated...")
}

func main() {
	initDatabase()

	app := fiber.New(fiber.Config{
		Concurrency:          256 * 1024,
		CompressedFileSuffix: ".fiber.gz",
	})

	// limiting 3 requests per 10 seconds
	app.Use(limiter.New(limiter.Config{
		Duration: 10 * time.Second,
		Max:      14,
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
		Output:     os.Stdout,
	}))

	setupRoutes(app)

	log.Fatal(app.Listen(":9090"))
	println("MobAPI")

	defer database.DBConn.Close()
}
