package main

import (
	"log"
	"net/http"
	"orderapi/handler"
)

var (
	port = ":8080"
)

func main() {

	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/get", handler.Get)
	http.HandleFunc("/update", handler.Update)
	http.HandleFunc("/delete", handler.Delete)

	log.Printf("Listening on %s...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
