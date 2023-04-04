package main

import (
	"log"
	"net/http"
)

// função principal
func main() {
	router := Router()
	log.Fatal(http.ListenAndServe(":8000", router))
}