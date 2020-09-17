package book

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/wborbajr/mobapi/database"
	"github.com/wborbajr/mobapi/model"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []model.Book
	db.Find(&books)

	return c.JSON(books)
}
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	db.Find(&book, id)
	return c.JSON(book)
}
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No Book Found with ID")
	}
	db.Delete(&book)
	return c.SendString("Book Successfully deleted")
}
func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&book)

	return c.JSON(book)
}
func UpdateBook(c *fiber.Ctx) error {
	return c.JSON("Update a book")
}
