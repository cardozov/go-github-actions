package main

import "testing"

func TestAdd(t *testing.T) {
	total := sumTwoValues(2, 3)
	if total != 5 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 5)
	}
}