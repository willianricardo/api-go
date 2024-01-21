package entity

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func TestOrderItem_JSONSerialization(t *testing.T) {
	// Create a sample order item
	expectedOrderItem := OrderItem{
		ID:        NewUniqueID(),
		ProductID: NewUniqueID(),
		Quantity:  2,
		Price:     29.99,
	}

	// Serialize the order item to JSON
	jsonData, err := json.Marshal(expectedOrderItem)
	if err != nil {
		t.Fatalf("Error marshaling OrderItem to JSON: %v", err)
	}

	// Deserialize the JSON back to an OrderItem
	var actualOrderItem OrderItem
	err = json.Unmarshal(jsonData, &actualOrderItem)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON to OrderItem: %v", err)
	}

	// Check if the original and deserialized order items are equal using DeepEqual
	if !reflect.DeepEqual(expectedOrderItem, actualOrderItem) {
		t.Errorf("Expected %+v, got %+v", expectedOrderItem, actualOrderItem)
	}
}

func TestOrderItem_NewOrderItem(t *testing.T) {
	testCases := []struct {
		TestName  string
		ID        string
		ProductID string
		Quantity  int
		Price     float64
		Expected  *OrderItem
		Err       error
	}{
		{
			TestName:  "Valid input",
			ID:        "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			ProductID: "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			Quantity:  1,
			Price:     20.0,
			Expected: &OrderItem{
				ID:        UniqueIDFromValue("2fc427d0-dd52-4a88-91a1-a387ccb36668"),
				ProductID: UniqueIDFromValue("2fc427d0-dd52-4a88-91a1-a387ccb36668"),
				Quantity:  1,
				Price:     20.0,
			},
			Err: nil,
		},
		{
			TestName:  "Zero quantity",
			ID:        "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			ProductID: "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			Quantity:  0,
			Price:     20.0,
			Expected:  nil,
			Err:       errors.New("quantity must be greater than 0"),
		},
		{
			TestName:  "Negative quantity",
			ID:        "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			ProductID: "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			Quantity:  -5,
			Price:     20.0,
			Expected:  nil,
			Err:       errors.New("quantity must be greater than 0"),
		},
		{
			TestName:  "Negative price",
			ID:        "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			ProductID: "2fc427d0-dd52-4a88-91a1-a387ccb36668",
			Quantity:  1,
			Price:     -5.0,
			Expected:  nil,
			Err:       errors.New("price must be greater than or equal to 0"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			orderItem, err := NewOrderItem(tc.ID, tc.ProductID, tc.Quantity, tc.Price)

			if err != nil && (tc.Err == nil || err.Error() != tc.Err.Error()) {
				t.Errorf("Expected error: %v, but got: %v", tc.Err, err)
			}

			if tc.Err == nil {
				if orderItem.ID == "" {
					t.Errorf("Expected non-empty order item ID, got empty")
				}

				if orderItem.ProductID == "" {
					t.Errorf("Expected non-empty order item product ID, got empty")
				}

				if orderItem.Quantity != tc.Expected.Quantity {
					t.Errorf("Expected order item quantity to be %d, got %d", tc.Expected.Quantity, orderItem.Quantity)
				}

				if orderItem.Price != tc.Expected.Price {
					t.Errorf("Expected order item price to be %f, got %f", tc.Expected.Price, orderItem.Price)
				}
			}
		})
	}
}
