package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/adventofcode/2017/helpers"
)

func main() {
	fileName := "../day02_corruptionChecksum/input"
	fmt.Println("reading ", fileName)
	inputBytes := helpers.Input(fileName)

	checksum(inputBytes)
}

func checksum(inputBytes []byte) int {
	var result int
	var lowest int
	var highest int
	var numStr string

	for i := 0; i < len(inputBytes); i++ {
		// convert to string
		str := string(inputBytes[i])
		switch str {
		case "	":
			// fmt.Println("tab, now can compare ", numStr)
			highest, lowest = readyToCompare(numStr, highest, lowest)
			numStr = ""
		case "\n":
			fmt.Println("newline! Need to compare last number. Setting highest and lowest back to zero and summing things up.")
			highest, lowest = readyToCompare(numStr, highest, lowest)

			numStr = ""
			result += highest - lowest
			highest = 0
			lowest = 0
			fmt.Println("result so far: ", result)
		default:
			// fmt.Println("character! Adding to str", str)
			numStr += str
		}
	}

	highest, lowest = readyToCompare(numStr, highest, lowest)
	numStr = ""
	result += highest - lowest

	fmt.Println("result: ", result)
	return result
}

func readyToCompare(numStr string, highest, lowest int) (int, int) {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}
	return compare(num, highest, lowest)
}

func compare(current, highest, lowest int) (int, int) {
	fmt.Println("comparing ", current, highest, lowest)
	if highest == 0 {
		highest = current
	}
	if lowest == 0 {
		lowest = current
	}
	if current > highest {
		highest = current
	}
	if current < lowest {
		lowest = current
	}

	// fmt.Println("comparing complete", highest, lowest)
	return highest, lowest
}
