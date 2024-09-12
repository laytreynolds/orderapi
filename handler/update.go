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

	// Extract the query parameters
	queryValues := r.URL.Query()
	str := queryValues.Get("id")
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Search the orders array for orders matching the id
	for i, order := range helpers.Orders {
		if order.ID == id {
			// Parse JSON Body into Order struct
			err := json.NewDecoder(r.Body).Decode(&order)
			if err != nil {
				http.Error(w, "Failed to decode request body", http.StatusBadRequest)
				return
			}

			// Update the original order in the slice
			helpers.Orders[i] = order // Save the updated order back to the slice

			// Pretty print the new order
			orderJSON, _ := json.MarshalIndent(order, "", "    ")
			log.Printf("Order %d Updated:\n%s\n", id, string(orderJSON))

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(order)
			return // Ensure to return after handling the update
		}
	}

	// If the order with the given ID was not found
	http.Error(w, "Order not found", http.StatusNotFound)
}
