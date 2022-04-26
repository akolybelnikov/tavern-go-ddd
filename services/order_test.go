package services

import (
	"github.com/akolybelnikov/tavern-go-ddd/domain/customer"
	"github.com/akolybelnikov/tavern-go-ddd/domain/product"
	"github.com/google/uuid"
	"testing"
)

func initProducts(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "healthy beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peanuts, err := product.NewProduct("Peanuts", "Healthy snack", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy drink", 2.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
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

	cust, err := customer.NewCustomer("Pavel")
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
