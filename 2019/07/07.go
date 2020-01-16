package main

import (
	"fmt"
	"log"
	"strconv"

	"../input"
)

func main() {
	input := input.ReadNumberInput("07")

	// input := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	// input := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
	// 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	// input := []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
	// 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}
	intcodeRunner(input)
}

func intcodeRunner(programs []int) int {
	var maxValue int
	combos := generateCombos()

	// var amp = 0
	// amp = intcode(programs, 0, 0)
	// fmt.Println("amp: ", amp)
	// amp = intcode(programs, 1, amp)
	// amp = intcode(programs, 2, amp)
	// amp = intcode(programs, 3, amp)
	// amp = intcode(programs, 4, amp)
	// fmt.Println("amp: ", amp)

	for _, combo := range combos {
		var amp = 0
		for ampIndex := 0; ampIndex < len(combo); ampIndex++ {
			input1 := combo[ampIndex]
			input2 := amp
			// fmt.Println("trying ", combos[0][ampIndex])
			amp = intcode(programs, input1, input2)
		}

		// fmt.Println("amp: ", amp)
		if amp > maxValue {
			maxValue = amp
		}
	}

	fmt.Println("max ", maxValue)
	return maxValue
}

func intcode(programs []int, input1, input2 int) int {
	var index int
	var hasReceivedSingleInput bool
	var lastOutputtedValue int
	for programs[index] != 99 {
		var progression int

		instruction := programs[index]

		// fmt.Printf(
		// 	"index %v instruction %v following ... %v %v %v ",
		// 	index, instruction, programs[index+1], programs[index+2], programs[index+3])
		code := instructionCode(strconv.Itoa(instruction))
		modes := getModes(instruction)
		// fmt.Println("code & modes ", code, modes)
		switch code {
		case 1, 2:
			p1, p2 := getParams(modes, index, programs)
			// fmt.Println("params ", p1, p2)
			idx3 := programs[index+3]

			if code == 1 {
				// fmt.Printf("%v + %v at index %v\n", p1, p2, idx3)
				// add
				val := p1 + p2
				programs[idx3] = val
			} else if code == 2 {
				// fmt.Printf("%v * %v at index %v\n", p1, p2, idx3)
				// multiply
				val := p1 * p2
				programs[idx3] = val
			} else {
				log.Fatal("not a code 1 or 2")
			}

			progression = 4
		case 3:
			input := input1
			if hasReceivedSingleInput {
				input = input2
			}
			// fmt.Println("case 3! ", hasReceivedSingleInput, input)
			hasReceivedSingleInput = true
			param := programs[index+1]
			// fmt.Printf("placing %v at index %v\n", input, param)
			programs[param] = input
			// only a single param, so only need to move 2 spaces
			progression = 2
		case 4:
			param := programs[index+1]
			if modes[0] == 0 {
				// fmt.Println("\nprinting ", programs[param])
				lastOutputtedValue = programs[param]
			} else {
				// fmt.Println("\nprinting ", param)
				lastOutputtedValue = param
			}
			// only a single param, so only need to move 2 spaces
			progression = 2
		case 5:
			p1, p2 := getParams(modes, index, programs)

			if p1 != 0 {
				// fmt.Printf("is %v != 0? YES", p1)
				progression = p2 - index
			} else {
				// fmt.Printf("is %v != 0? NO", p1)
				progression = 3
			}
		case 6:
			p1, p2 := getParams(modes, index, programs)

			if p1 == 0 {
				// fmt.Printf("is %v = 0? YES", p1)
				progression = p2 - index
			} else {
				// fmt.Printf("is %v = 0? NO", p1)
				progression = 3
			}
		case 7:
			p1, p2 := getParams(modes, index, programs)
			idx3 := programs[index+3]
			// fmt.Printf("is %v < %v? ", p1, p2)

			if p1 < p2 {
				programs[idx3] = 1
			} else {
				programs[idx3] = 0
			}
			progression = 4
		case 8:
			p1, p2 := getParams(modes, index, programs)
			idx3 := programs[index+3]
			// fmt.Printf("is %v == %v? will set at index 3 (idx: %v)\n", p1, p2, idx3)
			if p1 == p2 {
				programs[idx3] = 1
			} else {
				programs[idx3] = 0
			}
			progression = 4
		default:
			log.Fatalf("code unrecognized %v", code)
		}
		// fmt.Println(index, progression)
		index = index + progression
	}

	// fmt.Println("returning ", lastOutputtedValue)
	return lastOutputtedValue
}

func getModes(instruction int) []int {
	if instruction < 100 {
		return []int{0, 0}
	}

	if instruction < 1000 {
		str := strconv.Itoa(instruction)
		num, _ := strconv.Atoi(string(str[0]))
		return []int{num, 0}
	}
	if instruction < 10000 {
		str := strconv.Itoa(instruction)
		num1, _ := strconv.Atoi(string(str[0]))
		num0, _ := strconv.Atoi(string(str[1]))

		return []int{num0, num1}
	}

	str := strconv.Itoa(instruction)
	num1, _ := strconv.Atoi(string(str[1]))
	num0, _ := strconv.Atoi(string(str[2]))

	return []int{num0, num1}
}

func getParams(modes []int, index int, programs []int) (int, int) {
	var p1, p2 int
	if modes[0] == 0 {
		// positional
		idx1 := programs[index+1]
		p1 = programs[idx1]
	} else {
		// immediate
		p1 = programs[index+1]
	}
	if modes[1] == 0 {
		// positional
		idx2 := programs[index+2]
		p2 = programs[idx2]
	} else {
		p2 = programs[index+2]
	}

	return p1, p2
}

func instructionCode(instructions string) int {
	if instructions == "1" {
		return 1
	} else if instructions == "2" {
		return 2
	} else if instructions == "3" {
		return 3
	} else if instructions == "4" {
		return 4
	} else if instructions == "5" {
		return 5
	} else if instructions == "6" {
		return 6
	} else if instructions == "7" {
		return 7
	} else if instructions == "8" {
		return 8
	}
	// task last 2 digits
	codeStr := instructions[len(instructions)-2:]
	code, _ := strconv.Atoi(codeStr)
	if code < 1 || code > 8 {
		log.Fatalf("incorrect code %v instruction %v\n", code, instructions)
	}
	return code
}

func generateCombos() [][]int {
	availablePhases := []bool{true, true, true, true, true}

	// I cheated: I know there will be 120 combos to choose from
	combos := make([][]int, 120)
	for i := range combos {
		// A-E is 5 digits
		combos[i] = make([]int, 5)
	}

	// the index to put the combo in combos
	var i int

	for a := 0; a < len(availablePhases); a++ {
		availablePhases[a] = false
		for b := 0; b < len(availablePhases); b++ {
			if availablePhases[b] == false {
				continue
			}
			availablePhases[b] = false
			for c := 0; c < len(availablePhases); c++ {
				if availablePhases[c] == false {
					continue
				}
				availablePhases[c] = false

				for d := 0; d < len(availablePhases); d++ {
					if availablePhases[d] == false {
						continue
					}
					availablePhases[d] = false

					for e := 0; e < len(availablePhases); e++ {
						if availablePhases[e] == false {
							continue
						}
						combos[i] = []int{a, b, c, d, e}
						i++
					}

					availablePhases[d] = true
				}

				availablePhases[c] = true
			}
			availablePhases[b] = true
		}
		availablePhases[a] = true
	}

	return combos
}
