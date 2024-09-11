package handler

import (
	"log"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to update", r.Method)

	// Check Method
	if r.Method != http.MethodPut {

	}
}
