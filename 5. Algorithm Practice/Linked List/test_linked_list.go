package main

import "testing"

func TestSum(t *testing.T) {
	total := 0
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
