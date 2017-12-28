package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "input"
	fmt.Println("reading ", fileName)
	file, err := os.Open(fileName)
	// read the whole file at once
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	checksum(input)
}

func checksum(inputBytes []byte) int {
	var result int
	var input string
	var lowest int
	var highest int
	var numStr string

	for i := 0; i < len(inputBytes); i++ {
		str := string(inputBytes[i])
		if str == "	" {
			fmt.Println("tab! ", numStr, len(numStr))
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				fmt.Println("error!")
				return -1
			}
			highest, lowest = compare(num, highest, lowest)
			// resets
			numStr = ""
			num = 0
		}
		fmt.Println(str)
		if str == "\n" {
			fmt.Println("newline!")
			diff := highest - lowest
			result += diff
			// resets
			lowest = 0
			highest = 0
			fmt.Println("result is now ", result)
		}
		numStr += str
		fmt.Println(numStr)
	}
	fmt.Println(input)
	// walk thru each line
	// find max diff, sum
	// add to tally
	// repeat
	fmt.Println("result: ", result)
	return result
}

func compare(current, highest, lowest int) (int, int) {
	if highest == 0 {
		highest = current
	}
	if lowest == 0 {
		lowest = current
	}
	fmt.Println("comparing ", current, highest, lowest)
	if current > highest {
		highest = current
	}
	if current < lowest {
		lowest = current
	}
	fmt.Println("comparing complete", highest, lowest)
	return highest, lowest
}
