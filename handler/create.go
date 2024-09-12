package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"orderapi/helpers"
)

func Create(w http.ResponseWriter, r *http.Request) {

	// Log the incoming request
	log.Println("Received a request to create an order")

	// Check Method
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON Body into Order struct
	var order helpers.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Get values from JSON
	customer := order.Customer
	product := ""
	price := 0.0

	// Loop through product map and set values
	for p, pr := range order.Products {
		product = p
		price = pr
		break
	}

	// Create Order
	newOrder := helpers.Order{}
	err = newOrder.Create(customer, product, price)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	helpers.Orders = append(helpers.Orders, newOrder)

	// Pretty print the new order
	orderJSON, _ := json.MarshalIndent(newOrder, "", "    ")
	log.Printf("New Order Created:\n%s\n", string(orderJSON))

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newOrder)
}
