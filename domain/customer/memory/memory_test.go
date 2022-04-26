package memory

import (
	"github.com/akolybelnikov/tavern-go-ddd/domain/customer"
	"github.com/google/uuid"
	"testing"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := customer.NewCustomer("Pavel")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()

	repo := CustomerMemoryRepository{customers: map[uuid.UUID]customer.Customer{id: cust}}
	testCases := []testCase{
		{
			name:        "No Customer by ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "Customer by ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Customer",
			cust:        "Jen",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := CustomerMemoryRepository{customers: map[uuid.UUID]customer.Customer{}}

			c, err := customer.NewCustomer(tc.cust)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(c)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(c.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != c.GetID() {
				t.Errorf("Expected %v, got %v", c.GetID(), found.GetID())
			}
		})
	}
}
