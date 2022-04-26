package tavern

import (
	"github.com/akolybelnikov/tavern-go-ddd/domain/product"
	"github.com/akolybelnikov/tavern-go-ddd/services/order"
	"github.com/google/uuid"
	"testing"
)

func initProducts(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}

func TestTavern(t *testing.T) {
	products := initProducts(t)

	s, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products))
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(s))
	if err != nil {
		t.Error(err)
	}

	uid, err := s.AddCustomer("Pavel")
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetId(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}
}

func TestTavern_Mongo(t *testing.T) {
	products := initProducts(t)
	conn := "mongodb://localhost:27017"

	s, err := order.NewOrderService(
		order.WithMongoCustomerRepository(conn),
		order.WithMemoryProductRepository(products))
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(s))
	if err != nil {
		t.Error(err)
	}

	uid, err := s.AddCustomer("Pavel")
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetId(),
	}
	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}
}
