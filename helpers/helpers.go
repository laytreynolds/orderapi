package helpers

var Orders []Order

// Data Structure
type Order struct {
	Customer string             `json:"customer"`
	Products map[string]float64 `json:"products"`
}

// Create Helper
func (o *Order) Create(c string, product string, price float64) error {
	if o.Products == nil {
		o.Products = make(map[string]float64)
	}
	o.Customer = c
	o.Products[product] = price
	return nil
}
