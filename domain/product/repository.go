package product

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("the product wasn't found")
	ErrProductAlreadyExists = errors.New("the product already exists")
)

type Repository interface {
	GetAll() ([]Product, error)
	GetByID(uuid.UUID) (Product, error)
	Add(Product) error
	Update(Product) error
	Delete(uuid.UUID) error
}
