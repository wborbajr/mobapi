package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	ID        int       `json:"id,primary_key"`
	Title     string    `json:"name"`
	Author    string    `json:"author"`
	Rating    string    `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// 	DeletedAt time.Time `json:"deleted_at"`
}
