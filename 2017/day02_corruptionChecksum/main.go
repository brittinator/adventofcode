package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/adventofcode/2017/helpers"
)

func main() {
	fileName := "../day02_corruptionChecksum/input"
	fmt.Println("reading ", fileName)
	inputBytes := helpers.Input(fileName)

	checksum(inputBytes)

	fileName = "../day02_corruptionChecksum/input2"
	fmt.Println("reading ", fileName)
	inputBytes = helpers.Input(fileName)
	divisibleValues(inputBytes)
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
			// fmt.Println("newline! Need to compare last number. Setting highest and lowest back to zero and summing things up.")
			highest, lowest = readyToCompare(numStr, highest, lowest)

			numStr = ""
			result += highest - lowest
			highest = 0
			lowest = 0
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

func divisibleValues(inputBytes []byte) int {
	var result int
	var numStr string
	var nums []int

	for i := 0; i < len(inputBytes); i++ {
		// convert to string
		str := string(inputBytes[i])
		switch str {
		case "	":
			// fmt.Println("tab, now can compare ", numStr)
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)

			numStr = ""
		case "\n":
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
			numStr = ""

			fmt.Println("newline! Finding the evenly divisible numbers")
			res := mod(nums)
			nums = make([]int, 0)

			result += res
			fmt.Println("result so far: ", result)
		default:
			numStr += str
		}
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}
	nums = append(nums, num)
	numStr = ""

	fmt.Println("newline! Finding the evenly divisible numbers")
	res := mod(nums)
	fmt.Println("result of mod is ", res)

	nums = make([]int, 0)

	result += res

	fmt.Println("result: ", result)
	return result

}

func mod(nums []int) int {
	for _, n := range nums {
		for _, nn := range nums {
			if n == nn {
				continue
			}
			if math.Mod(float64(n), float64(nn)) == 0 {
				fmt.Printf("%v / %v = %v", n, nn, n/nn)
				return n / nn
			}
			if math.Mod(float64(nn), float64(n)) == 0 {
				fmt.Printf("%v / %v = %v", nn, n, nn/n)

				return nn / n
			}
		}
	}

	return 0
}
