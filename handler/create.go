package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"orderapi/helpers"
)

func Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received a %s request to create an order", r.Method)

	// Check Method
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON Body into Order struct
	var order helpers.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Validate order data
	if order.Customer == "" || len(order.Products) == 0 {
		http.Error(w, "Customer and products are required", http.StatusBadRequest)
		return
	}

	// Create Order
	newOrder := helpers.Create(order.Customer, order.Products)

	// Append the newly created order to the orders slice
	helpers.Orders = append(helpers.Orders, newOrder)

	// Pretty print the new order
	orderJSON, _ := json.MarshalIndent(newOrder, "", "    ")
	log.Printf("New Order Created:\n%s\n", string(orderJSON))

	// Set the response header and status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Send the new order as a response
	json.NewEncoder(w).Encode(newOrder)
}
