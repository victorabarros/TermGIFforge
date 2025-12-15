package id

import (
	"testing"
)

func TestNewUUUIDAsString(t *testing.T) {
	id := NewUUUIDAsString("test")
	if id == "" {
		t.Fatalf("Failed to generate UUID")
	}
	t.Logf("ID: %s", id)

	id = NewUUUIDAsString("test")
	if id == "" {
		t.Fatalf("Failed to generate UUID")
	}
	t.Logf("ID: %s", id)
}
