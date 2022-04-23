// Package memory is an in-memory implementation of the customer repository
package memory

import (
	"fmt"
	"github.com/akolybelnikov/goddd/aggregate"
	"github.com/akolybelnikov/goddd/domain/customer"
	"github.com/google/uuid"
	"sync"
)

// CustomerRepository fulfills the customer CustomerRepository interface
type CustomerRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *CustomerRepository {
	return &CustomerRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Get finds a customer by ID
func (r *CustomerRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if c, ok := r.customers[id]; ok {
		return c, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

// Add will add a new customer to the repository
func (r *CustomerRepository) Add(c aggregate.Customer) error {
	if r.customers == nil {
		r.Lock()
		r.customers = make(map[uuid.UUID]aggregate.Customer)
		r.Unlock()
	}
	// Make sure this customer doesn't exist yet
	if _, ok := r.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	r.Lock()
	r.customers[c.GetID()] = c
	r.Unlock()
	return nil
}

// Update will replace an existing customer information with the new customer information
func (r *CustomerRepository) Update(c aggregate.Customer) error {
	// Make sure the customer is in the repository
	if _, ok := r.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer doesn't exist: %w", customer.ErrUpdateCustomer)
	}
	r.Lock()
	r.customers[c.GetID()] = c
	r.Unlock()
	return nil
}
