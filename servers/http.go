package servers

import (
	"log"
	"net/http"
	"orderapi/handler"
)

func RunHTTPServer() {
	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/get", handler.Get)
	http.HandleFunc("/update", handler.Update)
	http.HandleFunc("/delete", handler.Delete)
	http.HandleFunc("/getall", handler.GetAll)

	log.Printf("Listening on %s...", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
