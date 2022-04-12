package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	books = getInitBooks()

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books", updateBook).Methods("PUT")
	router.HandleFunc("/api/books", deleteBook).Methods("DELETE")

	SERVER_PORT := goDotEnvVariable("PORT")

	println("Server will run on port " + SERVER_PORT)
	log.Fatal(http.ListenAndServe(":"+SERVER_PORT, router))
}
