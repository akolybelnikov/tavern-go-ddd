package product

import (
	"testing"
)

func TestProduct_NewProduct(t *testing.T) {
	type TestCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}

	testCases := []TestCase{
		{
			test:        "should return error is name is empty",
			name:        "",
			expectedErr: ErrMissingValues,
		},
		{
			test:        "valid values",
			name:        "test",
			description: "test",
			price:       1.0,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewProduct(tc.name, tc.description, tc.price)
			if err != tc.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}
