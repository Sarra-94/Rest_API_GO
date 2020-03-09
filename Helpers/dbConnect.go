package Helpers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"goTutorial/Models"
)

func DbConnect() *gorm.DB {
db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
if err != nil {
//defer db.Close()
fmt.Println(err)
}
return db

}
func Migration() {
	db := DbConnect()
	db.AutoMigrate(&Models.Author{})
	db.AutoMigrate(&Models.Book{})
}