// Package memory is a in memory implementation of the ProductRepository interface.
package memory

import (
	"github.com/akolybelnikov/goddd/domain/product"
	"github.com/google/uuid"
	"sync"
)

type ProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

// New is a factory function to generate a new repository of products
func New() *ProductRepository {
	return &ProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

// GetAll returns all products as a slice. A database implementation can return an error.
func (r *ProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
}

// GetByID searches for a product based on it's ID
func (r *ProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := r.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return product.Product{}, product.ErrProductNotFound
}

// Add will add a new product to the repository
func (r *ProductRepository) Add(newProduct product.Product) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[newProduct.GetId()]; ok {
		return product.ErrProductAlreadyExists
	}

	r.products[newProduct.GetId()] = newProduct
	return nil
}

// Update will change all values for a product based on it's ID
func (r *ProductRepository) Update(updateProduct product.Product) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[updateProduct.GetId()]; !ok {
		return product.ErrProductNotFound
	}

	r.products[updateProduct.GetId()] = updateProduct
	return nil
}

// Delete remove a product from the repository
func (r *ProductRepository) Delete(id uuid.UUID) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(r.products, id)
	return nil
}
