package main

import (
	"log"
	"net/http"
	"orderapi/handler"
)

var port = ":8080"

func main() {

	// Create an order
	http.HandleFunc("/create", handler.Create)

	// Get Order
	http.HandleFunc("/get", handler.Get)

	log.Printf("Listening on %s...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
