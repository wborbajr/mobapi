package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wborbajr/mobapi/model"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() error {
	var err error

	DBConn, err = gorm.Open("sqlite3", "./infra/books.sqlite")
	if err != nil {
		return err
	}

	fmt.Println("Connection opened to database...")
	DBConn.AutoMigrate(&model.Book{})
	fmt.Println("Database migrated...")

	return nil
}
