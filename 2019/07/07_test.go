package main

import (
	"testing"

	"../input"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	assert.Equal(t, intcodeRunner(input), 43210)
}

func TestPartTwo(t *testing.T) {
	input := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
		27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}

	// 98765
	assert.Equal(t, intcodeRunner(input), 139629729)

	// input = []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
	// 	-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
	// 	53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}

	// assert.Equal(t, intcodeRunner(input, 5), 18216)

}

func TestIntcodeStreamingRunner(t *testing.T) {
	// combo := []int{9, 8, 7, 6, 5}
	nums := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
		27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
	assert.Equal(t, 139629729, intcodeRunnerStreaming(nums))

	// combo := []int{9, 7, 8, 5, 6}
	nums = []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
		-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
		53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}

	assert.Equal(t, 18216, intcodeRunnerStreaming(nums))

	input := input.ReadNumberInput("07")
	intcodeRunner(input)
	assert.Equal(t, intcodeRunnerStreaming(input), 14365052)

}
