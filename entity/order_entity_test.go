package entity

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOrder_JSONSerialization(t *testing.T) {
	// Create a sample order
	expectedOrder := Order{
		ID:         NewUniqueID(),
		OrderDate:  "2024-01-20",
		CustomerID: NewUniqueID(),
		OrderItems: []*OrderItem{
			{
				ID:        NewUniqueID(),
				ProductID: NewUniqueID(),
				Quantity:  2,
				Price:     29.99,
			},
		},
	}

	// Serialize the order to JSON
	jsonData, err := json.Marshal(expectedOrder)
	if err != nil {
		t.Fatalf("Error marshaling Order to JSON: %v", err)
	}

	// Deserialize the JSON back to an Order
	var actualOrder Order
	err = json.Unmarshal(jsonData, &actualOrder)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON to Order: %v", err)
	}

	// Check if the original and deserialized orders are equal using DeepEqual
	if !reflect.DeepEqual(expectedOrder, actualOrder) {
		t.Errorf("Expected %+v, got %+v", expectedOrder, actualOrder)
	}
}
