// Package services holds all the services that connect repositories into a business flow
package services

import (
	"github.com/akolybelnikov/goddd/domain/customer"
	"github.com/google/uuid"
)

// OrderConfiguration is an alias for a function that will take in a pointer to an instance of OrderService and modify it
type OrderConfiguration func(s *OrderService) error

// OrderService is an implementation of the OrderService
type OrderService struct {
	customers customer.Repository
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

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) error {
	_, err := o.customers.Get(customerID)
	if err != nil {
		return err
	}

	return nil
}
