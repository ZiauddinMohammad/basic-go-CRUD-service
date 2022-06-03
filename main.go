package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ziauddinmohammad/basic-go-CRUD-service/routes"
)

func main() {

	// Create a router
	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	log.Println("Server started at port 8080 ...")
	//Start the server
	err := http.ListenAndServe(":8080", router)
	//EXit in case of error during server start
	if err != nil {
		log.Fatal(err)
	}
}
