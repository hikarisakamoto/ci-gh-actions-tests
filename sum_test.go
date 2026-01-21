package main

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(8, 2)

	if sum != 10 {
		t.Errorf("Expected %d and it was %d", 10, sum)
	}
}
