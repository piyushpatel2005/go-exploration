package testing

import (
	"fmt"
	"testing"
)

func TestCalculateValues(t *testing.T) {
	fmt.Println("hello from test")
	var (
		expected = 5
		a        = 2
		b        = 3
	)
	result := calculateValues(a, b)
	if result != expected {
		t.Errorf("Expected 5, got %d", result)
	}
}
