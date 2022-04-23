// Package services holds all the services that connect repositories into a business flow
package services

import (
	"github.com/akolybelnikov/goddd/aggregate"
	"github.com/akolybelnikov/goddd/domain/customer"
	"github.com/akolybelnikov/goddd/domain/customer/memory"
	"github.com/akolybelnikov/goddd/domain/product"
	memory2 "github.com/akolybelnikov/goddd/domain/product/memory"
	"github.com/google/uuid"
	"log"
)

// OrderConfiguration is an alias for a function that will take in a pointer to an instance of OrderService and modify it
type OrderConfiguration func(s *OrderService) error

// OrderService is an implementation of the OrderService
type OrderService struct {
	customers customer.Repository
	products  product.Repository
}

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	s := &OrderService{}
	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// WithCustomerRepository applies a given customer repository to the OrderService
func WithCustomerRepository(r customer.Repository) OrderConfiguration {
	return func(s *OrderService) error {
		s.customers = r
		return nil
	}
}

// WithMemoryCustomerRepository applies a memory customer repository to the OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	repo := memory.New()
	return WithCustomerRepository(repo)
}

// WithMemoryProductRepository adds a in memory product repo and adds all input products
func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(s *OrderService) error {
		r := memory2.New()

		for _, p := range products {
			err := r.Add(p)
			if err != nil {
				return err
			}
		}

		s.products = r
		return nil
	}
}

// CreateOrder will chain together all repositories to create an order for a customer
// will return the collected price of all Products
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	_, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []aggregate.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	log.Printf("Customer %s has ordered %d products", customerID, len(products))

	return price, nil
}
