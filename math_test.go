package main

import "testing"

func TestAdd(t *testing.T) {
	total := sum(2, 3)
	if total != 5 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 5)
	}
}

func TestAdd2(t *testing.T) {
	total := sum(-10, 10)
	if total != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 0)
	}
}