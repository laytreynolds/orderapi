package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"orderapi/helpers"
)

func Get(w http.ResponseWriter, r *http.Request) {

	// Log the incoming request
	log.Println("Received a request to retrieve orders")

	// Check Method
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Printf("HTTP Method %s Not Allowed", r.Method)

		return
	}

	// Extract the query parameter "customer"
	queryValues := r.URL.Query()
	customer := queryValues.Get("customer")

	// If the "customer" query parameter is not provided
	if customer == "" {
		http.Error(w, "Missing 'customer' query parameter", http.StatusBadRequest)
		log.Println("Missing 'Customer' Parameter")

		return
	}

	// Search the orders array for orders matching the customer name
	matchingOrders := []helpers.Order{}
	for _, order := range helpers.Orders {
		if order.Customer == customer {
			matchingOrders = append(matchingOrders, order)
		}
	}

	// Return the search results
	w.Header().Set("Content-Type", "application/json")
	if len(matchingOrders) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No orders found for the customer")
		log.Printf("Not Orders found for %s\n", customer)

	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(matchingOrders)
		orderJSON, _ := json.MarshalIndent(matchingOrders, "", "    ")
		log.Printf("Found Orders for %s:\n%s\n", customer, string(orderJSON))
	}
}
