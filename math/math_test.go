package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 3)
	if result != 4 {
		t.Errorf("Add(1, 3) = %d; want 4", result)
	}
}
