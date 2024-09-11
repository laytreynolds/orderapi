package helpers

import "sync"

var Orders []Order

// Data Structures
type Order struct {
	ID       int64              `json:"id"`
	Customer string             `json:"customer"`
	Products map[string]float64 `json:"products"`
}

type AutoInc struct {
	sync.Mutex
	id int64
}

func (a *AutoInc) ID() (id int64) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++

	return
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
