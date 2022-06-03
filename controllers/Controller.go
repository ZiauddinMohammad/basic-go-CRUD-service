package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ziauddinmohammad/basic-go-CRUD-service/models"
	"github.com/ziauddinmohammad/basic-go-CRUD-service/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["ID"], 0, 0)
	if err != nil {
		log.Println("error while parsing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, book := range models.Books {
		if book.ID == ID {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	result := models.BookResponse{Message: "ID not found"}
	x, err := json.Marshal(result)
	if err != nil {
		log.Println("error while Marshalling")
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write(x)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println("error while decoding")
	}
	book.ID = utils.NextID()
	models.Books = append(models.Books, book)
	result := models.BookResponse{Data: &book, Message: "New book added"}
	x, err := json.Marshal(result)
	if err != nil {
		log.Println("error while Marshalling")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(x)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["ID"], 0, 0)
	if err != nil {
		log.Println("error while parsing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedbook models.Book
	err1 := json.NewDecoder(r.Body).Decode(&updatedbook)
	if err1 != nil {
		log.Println("error while decoding")
	}
	if updatedbook.ID == 0 {
		updatedbook.ID = ID
	}
	for index, book := range models.Books {
		if book.ID == ID {
			models.Books = append(models.Books[:index], models.Books[index+1:]...)
			models.Books = append(models.Books, updatedbook)
			result := models.BookResponse{Data: &updatedbook, Message: "Updated"}
			x, err := json.Marshal(result)
			if err != nil {
				log.Println("error while Marshalling")
			}
			w.WriteHeader(http.StatusOK)
			w.Write(x)
			return
		}
	}
	result := models.BookResponse{Message: "ID not found"}
	x, err := json.Marshal(result)
	if err != nil {
		log.Println("error while Marshalling")
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write(x)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["ID"], 0, 0)
	if err != nil {
		log.Println("error while parsing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for index, book := range models.Books {
		if book.ID == ID {
			models.Books = append(models.Books[:index], models.Books[index+1:]...)
			result := models.BookResponse{Data: &book, Message: "Deleted"}
			x, err := json.Marshal(result)
			if err != nil {
				log.Println("error while Marshalling")
			}
			w.WriteHeader(http.StatusOK)
			w.Write(x)
			return
		}
	}
	result := models.BookResponse{Message: "ID not found"}
	x, err := json.Marshal(result)
	if err != nil {
		log.Println("error while Marshalling")
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write(x)
}
