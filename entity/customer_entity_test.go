package entity

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestCustomer_JSONSerialization(t *testing.T) {
	// Create a sample customer
	expectedCustomer := Customer{
		ID:   NewUniqueID(),
		Name: "John Doe",
	}

	// Serialize the customer to JSON
	jsonData, err := json.Marshal(expectedCustomer)
	if err != nil {
		t.Fatalf("Error marshaling Customer to JSON: %v", err)
	}

	// Deserialize the JSON back to a Customer
	var actualCustomer Customer
	err = json.Unmarshal(jsonData, &actualCustomer)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON to Customer: %v", err)
	}

	// Check if the original and deserialized customers are equal
	if expectedCustomer != actualCustomer {
		t.Errorf("Expected %+v, got %+v", expectedCustomer, actualCustomer)
	}
}

func TestCustomer_NewCustomer(t *testing.T) {
	testCases := []struct {
		TestName string
		ID       string
		Name     string
		Expected *Customer
		Err      error
	}{
		{
			TestName: "Valid input",
			ID:       "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			Name:     "Test Customer",
			Expected: &Customer{
				ID:   UniqueIDFromValue("2fc427d0-dd52-4a88-91a1-a387ccb36668"),
				Name: "Test Customer",
			},
			Err: nil,
		},
		{
			TestName: "Empty name",
			ID:       "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			Name:     "",
			Expected: nil,
			Err:      errors.New("name is required"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			customer, err := NewCustomer(tc.ID, tc.Name)

			if err != nil && (tc.Err == nil || err.Error() != tc.Err.Error()) {
				t.Errorf("Expected error: %v, but got: %v", tc.Err, err)
			}

			if tc.Err == nil {
				if customer.ID == "" {
					t.Errorf("Expected non-empty customer ID, got empty")
				}

				if customer.Name != tc.Expected.Name {
					t.Errorf("Expected customer name to be %s, got %s", tc.Expected.Name, customer.Name)
				}
			}
		})
	}
}
