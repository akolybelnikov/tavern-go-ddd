package services

import (
	"github.com/akolybelnikov/goddd/aggregate"
	"github.com/google/uuid"
	"testing"
)

func initProducts(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "healthy beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peanuts, err := aggregate.NewProduct("Peanuts", "Healthy snack", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := aggregate.NewProduct("Wine", "Healthy drink", 2.99)
	if err != nil {
		t.Error(err)
	}
	products := []aggregate.Product{
		beer, peanuts, wine,
	}

	return products
}

func TestOrder_NewOrderService(t *testing.T) {
	products := initProducts(t)

	s, err := NewOrderService(WithMemoryCustomerRepository(), WithMemoryProductRepository(products))
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("Pavel")
	if err != nil {
		t.Error(err)
	}

	err = s.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetId(),
	}

	_, err = s.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
