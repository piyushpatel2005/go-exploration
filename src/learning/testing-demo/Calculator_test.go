package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestAdd2(t *testing.T) {
	got := Add(-1, 2)
	want := 1
	if got != want {
		t.Errorf("got %d want %d", got, want)
		//t.Skip("Skipping test")
	}
}

// Table driven tests
func TestDivide(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{a: 4, b: 2, expected: 2},
		{a: 2, b: 4, expected: 0},
		{a: 10, b: 2, expected: 5},
	}

	for _, tc := range tests {
		got := Divide(tc.a, tc.b)
		if got != tc.expected {
			t.Errorf("Divide(%d, %d): got %d, want %d", tc.a, tc.b, got, tc.expected)
		}
	}
}
