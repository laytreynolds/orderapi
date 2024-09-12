package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"orderapi/helpers"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to update", r.Method)

	// Check Method
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Printf("HTTP Method %s Not Allowed", r.Method)
		return
	}

	// Extract the query parameters (ID)
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("Invalid ID:", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Search the orders array for orders matching the id
	for i := range helpers.Orders {
		order := &helpers.Orders[i] // Use a pointer to update the original order

		if order.ID == id {
			// Parse JSON Body into a new Order struct
			var updatedOrder helpers.Order
			if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
				log.Println("Failed to decode request body:", err)
				http.Error(w, "Failed to decode request body", http.StatusBadRequest)
				return
			}

			// Update fields of the original order
			order.Customer = updatedOrder.Customer
			order.Products = updatedOrder.Products // Assuming Products is a map

			// Pretty print the updated order
			orderJSON, _ := json.MarshalIndent(order, "", "    ")
			log.Printf("Order %d Updated:\n%s\n", id, string(orderJSON))

			// Set response header and status
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(order)
			return
		}
	}

	// If the order with the given ID was not found
	http.Error(w, "Order not found", http.StatusNotFound)
}
