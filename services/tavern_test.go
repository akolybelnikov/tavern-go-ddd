package services

import (
	"github.com/akolybelnikov/goddd/domain/customer"
	"github.com/google/uuid"
	"testing"
)

func TestTavern(t *testing.T) {
	products := initProducts(t)

	s, err := NewOrderService(WithMemoryCustomerRepository(), WithMemoryProductRepository(products))
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(s))
	if err != nil {
		t.Error(err)
	}

	customer, err := customer.NewCustomer("Dmitry")
	if err != nil {
		t.Error(err)
	}

	err = s.customers.Add(customer)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetId(),
	}
	err = tavern.Order(customer.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}

func TestTavern_Mongo(t *testing.T) {
	products := initProducts(t)
	conn := "mongodb://localhost:27017"

	s, err := NewOrderService(WithMongoCustomerRepository(conn), WithMemoryProductRepository(products))
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(s))
	if err != nil {
		t.Error(err)
	}

	customer, err := customer.NewCustomer("Dmitry")
	if err != nil {
		t.Error(err)
	}

	err = s.customers.Add(customer)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetId(),
	}
	err = tavern.Order(customer.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
