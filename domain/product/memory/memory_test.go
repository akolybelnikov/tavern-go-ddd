package memory

import (
	"github.com/akolybelnikov/goddd/aggregate"
	"github.com/akolybelnikov/goddd/domain/product"
	"github.com/google/uuid"
	"testing"
)

func TestMemoryProductRepository(t *testing.T) {
	repo := New()
	product, err := aggregate.NewProduct("Beer", "Good for your health", 1.99)
	if err != nil {
		t.Error(err)
	}

	err = repo.Add(product)
	if err != nil {
		t.Error(err)
	}

	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product in repo, got %d", len(repo.products))
	}
}

func TestMemoryProductRepository_Get(t *testing.T) {
	repo := New()
	testProduct, err := aggregate.NewProduct("Beer", "Good for your health", 1.99)
	if err != nil {
		t.Error(err)
	}

	err = repo.Add(testProduct)
	if err != nil {
		t.Error(err)
	}

	if len(repo.products) != 1 {
		t.Errorf("Expected 1 testProduct in repo, got %d", len(repo.products))
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get testProduct by id",
			id:          testProduct.GetId(),
			expectedErr: nil,
		},
		{
			name:        "Get non-existing testProduct by id",
			id:          uuid.New(),
			expectedErr: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err = repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemoryProductRepository_Delete(t *testing.T) {
	repo := New()
	testProduct, err := aggregate.NewProduct("Beer", "Good for your health", 1.99)
	if err != nil {
		t.Error(err)
	}

	err = repo.Add(testProduct)
	if err != nil {
		t.Error(err)
	}

	if len(repo.products) != 1 {
		t.Errorf("Expected 1 testProduct in repo, got %d", len(repo.products))
	}

	err = repo.Delete(testProduct.GetId())
	if err != nil {
		t.Error(err)
	}

	if len(repo.products) != 0 {
		t.Errorf("Expected 0 products in repo, got %d", len(repo.products))
	}
}
