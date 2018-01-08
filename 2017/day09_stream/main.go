package main

import (
	"fmt"

	"github.com/adventofcode/2017/helpers"
)

// What is the total score for all groups in your input?
// {} is a group
// <> is garbage
// ! deletes the char right after it

func removeAndString(inputBytes []byte) []byte {
	output := make([]byte, 0, len(inputBytes))
	i := 0
	for i <= len(inputBytes)-1 {
		b := inputBytes[i]
		if b == '!' {
			// skip this and next char
			i = i + 2
			continue
		}
		i++
		output = append(output, b)
	}
	return output
}

func score(inputBytes []byte) int {
	fmt.Println("score ", string(inputBytes))
	score := 0
	var stack []byte
	for _, b := range inputBytes {
		fmt.Println("stack: ", string(stack), score)
		if b == '{' {
			// add to stack
			stack = append(stack, b)
		}
		if b == '}' {
			// add length of stack to score
			// remove 1 from stack
			// pop off last item in stack
			if len(stack) > 0 {
				score += len(stack)
				stack = append(stack[:len(stack)-1])
			} else {
				score++
			}
		}
	}
	return score
}

func endOfJunk(input []byte, start int) int {
	for i := start + 1; i < len(input); i++ {
		if input[i] == '>' {
			fmt.Println("found at ", i, string(input[i]))
			return i
		}
	}
	return 0
}

func removeGarbage(inputBytes []byte) ([]byte, int) {
	score := 0
	output := make([]byte, 0, len(inputBytes))
	fmt.Println("Remove garbage: len ", len(inputBytes), string(inputBytes))
	i := 0
	for i < len(inputBytes) {
		b := inputBytes[i]
		fmt.Println("output", string(output), i, string(b))
		if b == '{' || b == '}' {
			output = append(output, b)
		}
		if b == '<' {
			fmt.Println("< found")
			// start of junk, need to find end of junk
			end := endOfJunk(inputBytes, i)
			fmt.Println("end ", end)
			// if end == 0 {
			// 	// cut off the rest of the input
			// 	// add
			// 	return output, score
			// }
			score = score + end - i - 1

			i = end + 1
			fmt.Println("i is now: ", i)
			continue
		}
		i++
	}
	fmt.Println("returning output: ", string(output))
	return output, score
}

func main() {
	// part 2: count up the garbage

	fileName := "../day09_stream/input"
	inputBytes := helpers.Input(fileName)
	inputBytes = removeAndString(inputBytes)
	inputBytes, junkScore := removeGarbage(inputBytes)
	fmt.Println(string(inputBytes))
	fmt.Println(score(inputBytes))
	fmt.Println("junk score: ", junkScore)

}
