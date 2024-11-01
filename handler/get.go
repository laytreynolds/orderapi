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
	customer := r.URL.Query().Get("customer")

	// If the "customer" query parameter is not provided
	if customer == "" {
		http.Error(w, "Missing 'customer' query parameter", http.StatusBadRequest)
		log.Println("Missing 'Customer' Parameter")
		return
	}

	// Search for orders matching the customer name
	matchingOrders := []helpers.Order{}
	for _, order := range helpers.Orders {
		if order.Customer == customer {
			matchingOrders = append(matchingOrders, order)
		}
	}

	if len(matchingOrders) == 0 {
		http.Error(w, "No orders found for the customer", http.StatusNotFound)
		log.Printf("No orders found for %s\n", customer)
		return
	}

	// Return the found orders
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(matchingOrders); err != nil {
		log.Println("Failed to encode response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	orderJSON, _ := json.MarshalIndent(matchingOrders, "", "    ")
	log.Printf("Found Orders for %s:\n%s\n", customer, string(orderJSON))
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request
	log.Println("Received a request to retrieve all orders")

	// Check Method
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Printf("HTTP Method %s Not Allowed", r.Method)
		return
	}

	// Set headers
	w.Header().Set("Content-Type", "application/json")

	// Check if orders is empty
	if len(helpers.Orders) == 0 {
		http.Error(w, "No orders found", http.StatusNotFound)
		log.Printf("No orders found")
		return
	}

	// Parse orders
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(helpers.Orders); err != nil {
		log.Println("Failed to encode response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	count := len(helpers.Orders)

	orderJSON, _ := json.MarshalIndent(helpers.Orders, "", "    ")
	log.Printf("Found %d Orders:\n%s\n", count, string(orderJSON))
}
