package Models

import "github.com/jinzhu/gorm"

// Book struct (Model)
type Book struct {
	gorm.Model
	//ID   uint `gorm:"primary_key;column:ID" json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
}


