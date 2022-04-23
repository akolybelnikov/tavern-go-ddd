package product

import (
	"errors"
	"github.com/akolybelnikov/goddd/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("the product wasn't found")
	ErrProductAlreadyExists = errors.New("the product already exists")
)

type Repository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(uuid.UUID) (aggregate.Product, error)
	Add(aggregate.Product) error
	Update(aggregate.Product) error
	Delete(uuid.UUID) error
}
