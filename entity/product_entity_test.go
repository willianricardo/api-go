package entity

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestProduct_JSONSerialization(t *testing.T) {
	// Create a sample product
	productID := NewUniqueID()
	expectedProduct := Product{
		ID:    productID,
		Name:  "Sample Product",
		Price: 29.99,
	}

	// Serialize the product to JSON
	jsonData, err := json.Marshal(expectedProduct)
	if err != nil {
		t.Fatalf("Error marshaling Product to JSON: %v", err)
	}

	// Deserialize the JSON back to a Product
	var actualProduct Product
	err = json.Unmarshal(jsonData, &actualProduct)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON to Product: %v", err)
	}

	// Check if the original and deserialized products are equal
	if expectedProduct != actualProduct {
		t.Errorf("Expected %+v, got %+v", expectedProduct, actualProduct)
	}
}

func TestProduct_NewProduct(t *testing.T) {
	testCases := []struct {
		TestName string
		ID       string
		Name     string
		Price    float64
		Expected *Product
		Err      error
	}{
		{
			TestName: "Valid input",
			ID:       "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			Name:     "Test Product",
			Price:    20.0,
			Expected: &Product{
				ID:    UniqueIDFromValue("2fc427d0-dd52-4a88-91a1-a387ccb36668"),
				Name:  "Test Product",
				Price: 20.0,
			},
			Err: nil,
		},
		{
			TestName: "Empty name",
			ID:       "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			Name:     "",
			Price:    20.0,
			Expected: nil,
			Err:      errors.New("name is required"),
		},
		{
			TestName: "Negative price",
			ID:       "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			Name:     "Test Product",
			Price:    -5.0,
			Expected: nil,
			Err:      errors.New("price must be greater than or equal to 0"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			product, err := NewProduct(tc.ID, tc.Name, tc.Price)

			if err != nil && (tc.Err == nil || err.Error() != tc.Err.Error()) {
				t.Errorf("Expected error: %v, but got: %v", tc.Err, err)
			}

			if tc.Err == nil {
				if product.ID == "" {
					t.Errorf("Expected non-empty product ID, got empty")
				}

				if product.Name != tc.Expected.Name {
					t.Errorf("Expected product name to be %s, got %s", tc.Expected.Name, product.Name)
				}

				if product.Price != tc.Expected.Price {
					t.Errorf("Expected product price to be %f, got %f", tc.Expected.Price, product.Price)
				}
			}
		})
	}
}
