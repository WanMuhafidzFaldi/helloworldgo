package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
} 
func Show(w http.ResponseWriter, r *http.Request) {
	var response Response
	response.Status = 0
	response.Message = "Hello World"
	json.NewEncoder(w).Encode(response)
}

func main() {

	router := mux.NewRouter() // create routes
	router.HandleFunc("/", Show).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":"+"", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
