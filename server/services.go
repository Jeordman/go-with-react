package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// get env variable for passed string
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// return starting value for api
func getInitBooks() []Book {
	println("Initializing books...")

	return append(books,
		Book{
			ID:    "1",
			Isbn:  "448743",
			Title: "test book 1",
			Author: &Author{
				Firstname: "John",
				Lastname:  "Doe",
			},
		},
		Book{
			ID:    "2",
			Isbn:  "454545",
			Title: "test book 2",
			Author: &Author{
				Firstname: "Steve",
				Lastname:  "Smith",
			},
		},
	)
}
