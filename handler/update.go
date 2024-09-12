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
	}

	// Search the orders array for orders matching the id
	for _, order := range helpers.Orders {
		if order.ID == id {
			// Parse JSON Body into Order structs
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

			if order.Products == nil {
				order.Products = make(map[string]float64)
			}
			order.Customer = customer
			order.Products[product] = price
		}

		// Pretty print the new order
		orderJSON, _ := json.MarshalIndent(order, "", "    ")
		log.Printf("Order %d Updated:\n%s\n", id, string(orderJSON))

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
	}

}
