package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goTutorial/Models"
	"log"
	"net/http"
	"strconv"
)

// Init Authors var as a slice Author struct
var Authors []Models.Author

// Get all Authors
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := dbConnect()
	json.NewEncoder(w).Encode(db.Find(&Authors))
}
func dbConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		//defer db.Close()
		fmt.Println(err)
	}
	return db

}


// Get single Author
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Gets params
	// Loop through Authors and find one with the id from the params
	for _, item := range Authors {
		id, err1 := strconv.ParseUint(r.FormValue("id"), 10, 32)
		if err1 != nil {
			fmt.Println(err1)
		}
		if uint(id) == item.ID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Models.Author{})
}

// Add new Author
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var Author Models.Author
	err := decoder.Decode(&Author)
	if err != nil {
		panic(err)
	}
	log.Println("new Author :" , Author)


	db := dbConnect()
	db.NewRecord(Author)
	db.Create(&Author)
	_ = json.NewDecoder(r.Body).Decode(&Author)
	Authors = append(Authors, Author)
	json.NewEncoder(w).Encode(Author)
}

// Update Author
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Authors {
		id, err1 := strconv.ParseUint(params["id"], 10, 32)
		if err1 != nil {
			fmt.Println(err1)
		}
		if item.ID == uint(id) {
			Authors = append(Authors[:index], Authors[index+1:]...)
			var Author Models.Author
			_ = json.NewDecoder(r.Body).Decode(&Author)
			Author.ID = uint(id)
			Authors = append(Authors, Author)
			json.NewEncoder(w).Encode(Author)
			return
		}
	}
}

// Delete Author
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Authors {
		id, err1 := strconv.ParseUint(params["id"], 10, 32)
		if err1 != nil {
			fmt.Println(err1)
		}
		if item.ID == uint(id) {
			Authors = append(Authors[:index], Authors[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Authors)
}
