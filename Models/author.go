package Models

import "github.com/jinzhu/gorm"

// Author struct
type Author struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Books []Book `gorm:"foreignkey:BookRefer"`
}

