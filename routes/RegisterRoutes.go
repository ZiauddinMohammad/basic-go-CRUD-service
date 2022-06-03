package routes

import (
	"github.com/gorilla/mux"
	"github.com/ziauddinmohammad/basic-go-CRUD-service/controllers"
)

func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")          // Gets all books
	router.HandleFunc("/book/{ID}", controllers.GetBookByID).Methods("GET")   // Gets book by id
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")        // Creates a book
	router.HandleFunc("/book/{ID}", controllers.UpdateBook).Methods("PUT")    // Updates a book
	router.HandleFunc("/book/{ID}", controllers.DeleteBook).Methods("DELETE") // Deletes a book

}
