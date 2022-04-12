package main

import (
	"encoding/json"
	"net/http"
)

// slice is variable length array
// init books var as slice Book struct
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
}

func createBook(w http.ResponseWriter, r *http.Request) {
}

func updateBook(w http.ResponseWriter, r *http.Request) {
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
}
