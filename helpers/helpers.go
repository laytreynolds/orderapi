package helpers

import (
	"fmt"
	"log"
)

var Orders []Order

// Data Structures

type Order struct {
	ID       int64              `json:"id"`
	Customer string             `json:"customer"`
	Products map[string]float64 `json:"products"`
}

var NextID int64 = 1 // Start with 1 or any initial value

// CreateOrder function to create a new order
func Create(customer string, products map[string]float64) Order {
	order := Order{
		ID:       NextID,
		Customer: customer,
		Products: products,
	}
	NextID++ // Increment to prepare for the next order
	return order
}

// AppendOrder function to add an order to the Orders slice
func AppendOrder(order Order) {
	Orders = append(Orders, order)
}

// GetOrders function to retrieve all orders
func GetOrders() []Order {
	return Orders
}

func Logger(format string, args ...interface{}) {

	message := fmt.Sprintf(format, args...)

	log.Print(message) // Log to log file (or wherever your log is set up)
}
