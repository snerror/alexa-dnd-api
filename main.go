package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	log.Println("Server is up and running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
