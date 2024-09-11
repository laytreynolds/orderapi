package handler

import (
	"log"
	"net/http"
	"orderapi/helpers"
)

func Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to update", r.Method)

	// Check Method
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Printf("HTTP Method %s Not Allowed", r.Method)
		return
	}

	// Extract the query parameter "customer"
	queryValues := r.URL.Query()
	customer := queryValues.Get("customer")

	// Search the orders array for orders matching the customer name
	matchingOrders := []helpers.Order{}
	for _, order := range helpers.Orders {
		if order.Customer == customer {

		}
	}

}
