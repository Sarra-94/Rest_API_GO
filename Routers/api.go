package Routers

import "github.com/gorilla/mux"
import 	"goTutorial/Controllers"

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	// Route handles & endpoints

	// book routers
	r.HandleFunc("/books", Controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", Controllers.GetBook).Methods("GET")
	r.HandleFunc("/books", Controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", Controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", Controllers.DeleteBook).Methods("DELETE")

   // author routers
	r.HandleFunc("/authors", Controllers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", Controllers.GetAuthor).Methods("GET")
	r.HandleFunc("/authors", Controllers.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors/{id}", Controllers.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", Controllers.DeleteAuthor).Methods("DELETE")
	return r
}