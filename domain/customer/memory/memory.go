// Package memory is an in-memory implementation of the customer repository
package memory

import (
	"fmt"
	"github.com/akolybelnikov/tavern-go-ddd/domain/customer"
	"github.com/google/uuid"
	"sync"
)

// CustomerMemoryRepository fulfills the customer CustomerMemoryRepository interface
type CustomerMemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *CustomerMemoryRepository {
	return &CustomerMemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

// Get finds a customer by ID
func (r *CustomerMemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if c, ok := r.customers[id]; ok {
		return c, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}

// Add will add a new customer to the repository
func (r *CustomerMemoryRepository) Add(c customer.Customer) error {
	if r.customers == nil {
		r.Lock()
		r.customers = make(map[uuid.UUID]customer.Customer)
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
func (r *CustomerMemoryRepository) Update(c customer.Customer) error {
	// Make sure the customer is in the repository
	if _, ok := r.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer doesn't exist: %w", customer.ErrUpdateCustomer)
	}
	r.Lock()
	r.customers[c.GetID()] = c
	r.Unlock()
	return nil
}
