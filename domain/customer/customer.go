// Package customer holds aggregates that combine many entities into a full object
package customer

import (
	"errors"
	"github.com/akolybelnikov/tavern-go-ddd"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("customer aggregate has to have a valid person")
)

// Customer is an aggregate that combines all entities needed to represent a customer
type Customer struct {
	// person is the root entity of the customer with person.ID being the main identifier for this aggregate
	person       *tavern.Person
	products     []*tavern.Item
	transactions []tavern.Transaction
}

// NewCustomer is a factory to validate and create a new Customer aggregate
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &tavern.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:       person,
		products:     make([]*tavern.Item, 0),
		transactions: make([]tavern.Transaction, 0),
	}, nil
}

// GetID returns the customer's root entity ID
func (c Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID sets the root ID
func (c Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.ID = id
}

// SetName changes customer's name
func (c Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.Name = name
}

// GetName returns the customer's root entity Name
func (c Customer) GetName() string {
	return c.person.Name
}
