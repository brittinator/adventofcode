package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"../input"
)

func main() {
	numbers := input.ReadNumberInput("05")
	// numbers := []int{1002, 4, 3, 4, 33}
	// numbers := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
	// 	1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
	// 	999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	// // numbers := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	intcode(numbers)
}

func intcode(programs []int) {
	var index int
	for programs[index] != 99 {
		var progression int

		instruction := programs[index]

		fmt.Printf(
			"index %v instruction %v following ... %v %v %v ",
			index, instruction, programs[index+1], programs[index+2], programs[index+3])
		code := instructionCode(strconv.Itoa(instruction))
		modes := getModes(instruction)
		fmt.Println("code & modes ", code, modes)
		switch code {
		case 1, 2:
			p1, p2 := getParams(modes, index, programs)
			// fmt.Println("params ", p1, p2)
			idx3 := programs[index+3]

			if code == 1 {
				fmt.Printf("%v + %v at index %v\n", p1, p2, idx3)
				// add
				val := p1 + p2
				programs[idx3] = val
			} else if code == 2 {
				fmt.Printf("%v * %v at index %v\n", p1, p2, idx3)
				// multiply
				val := p1 * p2
				programs[idx3] = val
			} else {
				log.Fatal("not a code 1 or 2")
			}

			progression = 4
		case 3:
			input := getInput()
			param := programs[index+1]
			fmt.Printf("placing %v at index %v\n", input, param)
			programs[param] = input
			// only a single param, so only need to move 2 spaces
			progression = 2
		case 4:
			param := programs[index+1]
			if modes[0] == 0 {
				fmt.Println("\nprinting ", programs[param])
			} else {
				fmt.Println("\nprinting ", param)
			}
			// only a single param, so only need to move 2 spaces
			progression = 2
		case 5:
			p1, p2 := getParams(modes, index, programs)

			if p1 != 0 {
				fmt.Printf("is %v != 0? YES", p1)
				progression = p2 - index
			} else {
				fmt.Printf("is %v != 0? NO", p1)
				progression = 3
			}
		case 6:
			p1, p2 := getParams(modes, index, programs)

			if p1 == 0 {
				fmt.Printf("is %v = 0? YES", p1)
				progression = p2 - index
			} else {
				fmt.Printf("is %v = 0? NO", p1)
				progression = 3
			}
		case 7:
			p1, p2 := getParams(modes, index, programs)
			idx3 := programs[index+3]
			fmt.Printf("is %v < %v? ", p1, p2)

			if p1 < p2 {
				programs[idx3] = 1
			} else {
				programs[idx3] = 0
			}
			progression = 4
		case 8:
			p1, p2 := getParams(modes, index, programs)
			idx3 := programs[index+3]
			fmt.Printf("is %v == %v? will set at index 3 (idx: %v)\n", p1, p2, idx3)
			if p1 == p2 {
				programs[idx3] = 1
			} else {
				programs[idx3] = 0
			}
			progression = 4
		default:
			log.Fatalf("code unrecognized %v", code)
		}
		fmt.Println(index, progression)
		index = index + progression
	}
}

func getInput() int {
	// return 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter input")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	fmt.Println("text ", text)
	n, _ := strconv.Atoi(text)

	return n
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
