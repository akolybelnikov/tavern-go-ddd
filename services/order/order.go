// Package order holds all the services that connect repositories into a business flow
package order

import (
	"context"
	"github.com/akolybelnikov/tavern-go-ddd/domain/customer"
	"github.com/akolybelnikov/tavern-go-ddd/domain/customer/memory"
	"github.com/akolybelnikov/tavern-go-ddd/domain/customer/mongo"
	"github.com/akolybelnikov/tavern-go-ddd/domain/product"
	memory2 "github.com/akolybelnikov/tavern-go-ddd/domain/product/memory"
	"github.com/google/uuid"
	"log"
)

// Configuration is an alias for a function that will take in a pointer to an instance of OrderService and modify it
type Configuration func(s *Service) error

// Service is an implementation of the OrderService
type Service struct {
	customers customer.Repository
	products  product.Repository
}

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...Configuration) (*Service, error) {
	s := &Service{}
	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// WithCustomerRepository applies a given customer repository to the OrderService
func WithCustomerRepository(r customer.Repository) Configuration {
	return func(s *Service) error {
		s.customers = r
		return nil
	}
}

// WithMemoryCustomerRepository applies a memory customer repository to the OrderService
func WithMemoryCustomerRepository() Configuration {
	repo := memory.New()
	return WithCustomerRepository(repo)
}

func WithMongoCustomerRepository(connectionString string) Configuration {
	return func(s *Service) error {
		r, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		s.customers = r
		return nil
	}
}

// WithMemoryProductRepository adds a in memory product repo and adds all input products
func WithMemoryProductRepository(products []product.Product) Configuration {
	return func(s *Service) error {
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
func (o *Service) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	_, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []product.Product
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

// AddCustomer will add a new customer and return the customerID
func (o *Service) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}
	// Add to Repo
	err = o.customers.Add(c)
	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
