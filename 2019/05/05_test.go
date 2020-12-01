package main

import (
	"testing"
)

func Test(t *testing.T) {
	// numbers := input.ReadNumberInput("05")
	numbers := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	intcode(numbers)
}
