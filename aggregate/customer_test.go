package aggregate

import "testing"

func TestCustomer_NewCustomer(t *testing.T) {
	type TestCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []TestCase{
		{
			test:        "Empty Name validation",
			name:        "",
			expectedErr: ErrInvalidPerson,
		},
		{
			test:        "Valid Name",
			name:        "John Smith Jr.",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
