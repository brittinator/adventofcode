package main

import (
	"testing"

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

// func TestIntcodeStreaming(t *testing.T) {
// 	// input := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
// 	// 	27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
// 	// func intcodeStreaming(programs []int, incoming <-chan int, outgoing chan<- int) {
// 	input := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}

// 	in := make(chan int, 2)
// 	out := make(chan int, 1)

// 	in <- 1
// 	in <- 432

// 	go intcodeStreaming(input, in, out)

// 	fmt.Println("hello")

// 	var value int
// 	for value = range out {
// 		fmt.Println("value ", value)

// 	}

// 	assert.Equal(t, value, 4321)
// }

func TestIntcodeStreamingRunner(t *testing.T) {
	nums := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
		27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
	intcodeRunnerStreaming(nums)
}
