package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestUniqueID_NewUniqueID(t *testing.T) {
	id := NewUniqueID()

	// Check if the generated ID is a valid UUID
	_, err := uuid.Parse(string(id))
	if err != nil {
		t.Errorf("NewUniqueID() generated an invalid UUID: %v", err)
	}
}

func TestUniqueID_UniqueIDFromValue(t *testing.T) {
	validUUID := "550e8400-e29b-41d4-a716-446655440000"
	invalidValue := "invalid-uuid-value"

	// Test with a valid UUID
	id := UniqueIDFromValue(validUUID)
	if string(id) != validUUID {
		t.Errorf("UniqueIDFromValue(%s) expected %s, got %s", validUUID, validUUID, string(id))
	}

	// Test with an invalid UUID
	id = UniqueIDFromValue(invalidValue)
	if string(id) != "" {
		t.Errorf("UniqueIDFromValue(%s) expected empty UniqueID, got %s", invalidValue, string(id))
	}
}

func TestUniqueID_ValueFromUniqueID(t *testing.T) {
	originalUUID := "550e8400-e29b-41d4-a716-446655440000"
	id := UniqueID(originalUUID)

	// Test conversion from UniqueID to string
	result := ValueFromUniqueID(id)
	if result != originalUUID {
		t.Errorf("ValueFromUniqueID(%s) expected %s, got %s", originalUUID, originalUUID, result)
	}
}
