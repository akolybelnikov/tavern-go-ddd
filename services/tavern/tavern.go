package tavern

import (
	"github.com/akolybelnikov/tavern-go-ddd/services/order"
	"github.com/google/uuid"
	"log"
)

type Configuration func(t *Tavern) error

type Tavern struct {
	// Order service is used to handle orders
	OrderService *order.Service
	// Billing service is used to handle billing
	BillingService interface{}
}

// NewTavern takes a variable amount of TavernConfigurations and builds a Tavern
func NewTavern(cfgs ...Configuration) (*Tavern, error) {
	t := &Tavern{}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

// WithOrderService applies a given OrderService to the Tavern
func WithOrderService(s *order.Service) Configuration {
	return func(t *Tavern) error {
		t.OrderService = s
		return nil
	}
}

// Order performs an order for a customer
func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf("Bill the customer %0.0f", price)
	// Bill the customer
	//err = t.BillingService.Bill(customer, price)
	return nil
}
