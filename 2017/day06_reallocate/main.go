package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/adventofcode/2017/helpers"
)

func convertToArray(input []byte) []int {
	output := make([]int, 0, len(input))

	var numStr string
	for _, b := range input {
		str := string(b)
		switch str {
		case "	":
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
	// need to add last num to array
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}
	output = append(output, num)

	return output
}

type bank struct {
	bnk    []int
	cycles int
	combos map[string]int
}

func (b *bank) findMax() (int, int) {
	var max int
	var index int
	for i := 0; i <= len(b.bnk)-1; i++ {
		if b.bnk[i] > max {
			max = b.bnk[i]
			index = i
		}
	}

	return max, index
}

func (b *bank) redistribute(amt, index int) {
	// fmt.Println("amoutn at start ", amt)
	if amt == 0 {
		return
	}

	for {
		for i := index + 1; i <= len(b.bnk)-1; i++ {
			if amt <= 0 {
				return
			}
			b.bnk[i]++
			amt--
			// fmt.Println("redistributing rest", b.bnk)
			// fmt.Println("amoutn left ", amt)
		}

		index = -1
	}
}

func (b *bank) seen() bool {
	var currentCombo string
	for _, n := range b.bnk {
		currentCombo += strconv.Itoa(n)
	}

	if _, seen := b.combos[currentCombo]; seen {
		return seen
	}
	b.combos[currentCombo] = 1
	return false
}

func main() {
	fileName := "../day06_reallocate/input"
	inputBytes := helpers.Input(fileName)
	input := convertToArray(inputBytes)
	// fmt.Println("input >>>", input)

	bank := bank{bnk: input, cycles: 0, combos: make(map[string]int)}
	for {
		// fmt.Println("number of cycles :'", bank.cycles)
		max, index := bank.findMax()
		// fmt.Println("max and index are :", bank.bnk, max, index)
		// remove values at index
		input[index] = 0

		bank.redistribute(max, index)
		// fmt.Println("after redistributed: ", bank.bnk)
		// add a cycle
		bank.cycles++
		// add to seen bucket
		if bank.seen() {
			fmt.Printf("combo %v seen before. \n Took %v cycles.\n", bank.bnk, bank.cycles)
			bank.cycles = 0
			bank.combos = map[string]int{}
			var toFind string
			for _, n := range bank.bnk {
				toFind += strconv.Itoa(n)
			}
			bank.combos[toFind] = 1
			fmt.Println("resetting state and finding it again")
			// return
			// sloppy nested for loops but it works b/c returning after found
			for {
				// fmt.Println("number of cycles :'", bank.cycles)
				max, index := bank.findMax()
				// fmt.Println("max and index are :", bank.bnk, max, index)
				// remove values at index
				input[index] = 0

				bank.redistribute(max, index)
				// fmt.Println("after redistributed: ", bank.bnk)
				// add a cycle
				bank.cycles++
				// add to seen bucket
				if bank.seen() {
					fmt.Printf("combo %v seen before. \n Took %v cycles to see it a 3rd time.\n", bank.bnk, bank.cycles)
					return
				}
			}
		}
	}
}
