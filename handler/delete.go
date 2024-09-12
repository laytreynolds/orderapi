package handler

import (
	"log"
	"net/http"
	"orderapi/helpers"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to delete an order", r.Method)

	// Check Method
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Printf("HTTP Method %s Not Allowed", r.Method)
		return
	}

	// Extract the query parameter "id"
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid ID:", err)
		return
	}

	// Search for the order with the specified ID
	for i, order := range helpers.Orders {
		if order.ID == id {
			// Remove the order by appending the slices
			helpers.Orders = append(helpers.Orders[:i], helpers.Orders[i+1:]...) // Remove the order

			// Log and respond
			log.Printf("Order %d deleted successfully\n", id)
			w.WriteHeader(http.StatusNoContent) // Successful deletion, no content to return
			return
		}
	}

	// If the order with the given ID was not found
	http.Error(w, "Order not found", http.StatusNotFound)
	log.Printf("Order %d not found\n", id)
}
