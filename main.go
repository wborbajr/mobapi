package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wborbajr/mobapi/database"
	"github.com/wborbajr/mobapi/router/book"
	"github.com/wborbajr/mobapi/router/ping"
)

func main() {

	// Connect to database
	if err := database.InitDatabase(); err != nil {
		log.Fatal("failed to connect database: ", err)
	}

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
		TimeZone:   "America/Sao_Paulo",
		Output:     os.Stdout,
	}))

	book.SetupRoutes(app)
	ping.SetupRoutes(app)

	log.Printf("[MobAPI] up and running: http://127.0.0.1:%s", "9090")
	log.Fatal(app.Listen(":9090"))

	defer database.DBConn.Close()
}
