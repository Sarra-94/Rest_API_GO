package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goTutorial/Models"
	"log"
	"net/http"
	"strconv"
)

// Init books var as a slice Book struct
var books []Models.Book

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := dbConnect()
	json.NewEncoder(w).Encode(db.Find(&books))
}



// Get single book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		id, err1 := strconv.ParseUint(r.FormValue("id"), 10, 32)
		if err1 != nil {
			fmt.Println(err1)
		}
		if uint(id) == item.ID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Models.Book{})
}

// Add new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var book Models.Book
	err := decoder.Decode(&book)
	if err != nil {
		panic(err)
	}
	log.Println("new book :", book)

	db := dbConnect()
	db.NewRecord(book)
	db.Create(&book)
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		id, err1 := strconv.ParseUint(params["id"], 10, 32)
		if err1 != nil {
			fmt.Println(err1)
		}
		if item.ID == uint(id) {
			books = append(books[:index], books[index+1:]...)
			var book Models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = uint(id)
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		id, err1 := strconv.ParseUint(params["id"], 10, 32)
		if err1 != nil {
			fmt.Println(err1)
		}
		if item.ID == uint(id) {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
