package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/adventofcode/2017/helpers"
)

// GOT THIS IN THE FIRST TRY WHOO HOO!!!!!!!

type maze struct {
	mmap   []int
	cursor int
	steps  int
}

func convertToSlice(input []byte) []int {
	output := make([]int, 0, len(input))

	var numStr string
	for _, b := range input {
		// convert to string
		str := string(b)
		switch str {
		case "\n":
			// end of num
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			output = append(output, num)
			numStr = ""
		default:
			numStr += str
		}
	}
	// last num
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}
	output = append(output, num)

	return output
}

// move current num number of spaces, return true if reached end of maze
func (m *maze) move() bool {
	// fmt.Println("map is ", m.mmap)
	// fmt.Println("cursor is ", m.cursor)
	// find cursor and grab old value
	cur := m.cursor // 0
	// check if cursor will still be in map
	if m.cursor >= len(m.mmap) {
		return true
	}
	// move cursor to next value
	m.cursor = cur + m.mmap[cur] // 0 + 0
	// increment old cursor value +1
	m.mmap[cur]++
	// increment step
	m.steps++
	// fmt.Println("end map is ", m.mmap)
	// fmt.Println("end cursor is ", m.cursor)
	return false
}

func main() {
	fileName := "../day05_maze/input"
	fmt.Println("reading ", fileName)
	inputBytes := helpers.Input(fileName)

	input := convertToSlice(inputBytes)
	// fmt.Println("input ", input)
	maze := maze{input, 0, 0}

	for {
		if maze.move() {
			fmt.Printf("maze completed in %v steps\n", maze.steps)
			return
		}
	}
}
