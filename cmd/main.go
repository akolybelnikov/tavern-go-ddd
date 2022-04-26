// Package main runs the tavern and performs an Order
package main

import (
	"github.com/akolybelnikov/tavern-go-ddd/domain/product"
	"github.com/akolybelnikov/tavern-go-ddd/services/order"
	"github.com/akolybelnikov/tavern-go-ddd/services/tavern"
	"github.com/google/uuid"
)

func main() {

	products := productInventory()
	// Create Order Service to use in tavern
	os, err := order.NewOrderService(
		//order.WithMongoCustomerRepository("mongodb://localhost:27017"),
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}
	// Create tavern service
	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(os))
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetId(),
	}
	// Execute Order
	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		panic(err)
	}
	peanuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	products := []product.Product{
		beer, peanuts, wine,
	}
	return products
}
