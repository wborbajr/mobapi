package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating string `json:"rating"`
}
