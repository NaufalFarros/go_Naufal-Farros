package Models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
}
