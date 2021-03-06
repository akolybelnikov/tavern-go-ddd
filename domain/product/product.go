// Package product holds aggregates that combine many entities into a full object
// Product is an aggregate that represents a product
package product

import (
	"errors"
	"github.com/akolybelnikov/tavern-go-ddd"
	"github.com/google/uuid"
)

var (
	ErrMissingValues = errors.New("missing product values")
)

// Product is an aggregate that combines item with a price and a quantity
type Product struct {
	// item is the root entity which is an Item
	item  *tavern.Item
	price float64
	// quantity is the number of products in stock
	quantity int
}

// NewProduct will create a new product
// It will return an error if name or description are empty
func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &tavern.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetId() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *tavern.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
