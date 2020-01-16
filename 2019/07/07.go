package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"../input"
)

func main() {
	input := input.ReadNumberInput("07")
	intcodeRunner(input)
}

func intcodeRunnerStreaming(programs []int) int {
	// combo := []int{9, 8, 7, 6, 5}
	inA := make(chan int, 10) // from E -> A
	inB := make(chan int, 10) // from A -> B
	inC := make(chan int, 10) // from B -> C
	inD := make(chan int, 10) // from C -> D
	inE := make(chan int, 10) // from D -> E

	var wg sync.WaitGroup
	// send the first 2 values
	inA <- 9
	inA <- 0
	wg.Add(1)
	// give each a copy of programs!!!!!!!!!
	pgm := make([]int, len(programs))
	for i, p := range programs {
		pgm[i] = p
	}
	go intcodeStreaming(pgm, inA, inB, &wg, "A")
	inB <- 8
	wg.Add(1)
	pgm = make([]int, len(programs))
	for i, p := range programs {
		pgm[i] = p
	}
	go intcodeStreaming(pgm, inB, inC, &wg, "B")
	inC <- 7
	wg.Add(1)
	pgm = make([]int, len(programs))
	for i, p := range programs {
		pgm[i] = p
	}
	go intcodeStreaming(pgm, inC, inD, &wg, "C")
	inD <- 6
	wg.Add(1)
	pgm = make([]int, len(programs))
	for i, p := range programs {
		pgm[i] = p
	}
	go intcodeStreaming(pgm, inD, inE, &wg, "D")
	inE <- 5
	wg.Add(1)
	pgm = make([]int, len(programs))
	for i, p := range programs {
		pgm[i] = p
	}
	go intcodeStreaming(pgm, inE, inA, &wg, "E")

	wg.Wait()
	return -1
}

func intcodeRunner(programs []int) int {
	var maxValue int
	combos := generateCombos()

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

func intcodeStreaming(programs []int, incoming <-chan int, outgoing chan<- int, wg *sync.WaitGroup, letter string) {
	var index int
	for programs[index] != 99 {
		l := log.New(os.Stdout, fmt.Sprintf("%v:%v: ", letter, index), 0)
		// progression is how far to move the indexes forward.
		var progression int
		instruction := programs[index]

		// l.Printf(
		// 	"index %v instruction %v following ... %v %v %v ",
		// 	index, instruction, programs[index+1], programs[index+2], programs[index+3])
		code := instructionCode(strconv.Itoa(instruction))
		modes := getModes(instruction)
		// fmt.Println("code & modes ", code, modes)
		switch code {
		case 1, 2:
			p1, p2 := getParams(modes, index, programs)
			idx3 := programs[index+3]
			if code == 1 { // add
				// fmt.Printf("%v + %v at index %v\n", p1, p2, idx3)
				val := p1 + p2
				programs[idx3] = val
			} else { // multiply
				// fmt.Printf("%v * %v at index %v\n", p1, p2, idx3)
				val := p1 * p2
				programs[idx3] = val
			}
			progression = 4

		case 3:
			// fmt.Println("get Input ...")
			input := <-incoming
			l.Println("input is", input)
			param := programs[index+1]
			// fmt.Printf("placing %v at index %v\n", input, param)
			programs[param] = input
			progression = 2
		case 4:
			param := programs[index+1]
			if modes[0] == 0 {
				param = programs[param]
			}
			// l.Println("outgoing ...", param)
			outgoing <- param
			l.Println("completed outgoing of ", param)
			progression = 2
		case 5:
			p1, p2 := getParams(modes, index, programs)
			progression = 3
			if p1 != 0 {
				// fmt.Printf("is %v != 0? YES", p1)
				progression = p2 - index
			}
		case 6:
			p1, p2 := getParams(modes, index, programs)
			progression = 3
			if p1 == 0 {
				// fmt.Printf("is %v = 0? YES", p1)
				progression = p2 - index
			}
		case 7:
			p1, p2 := getParams(modes, index, programs)
			idx3 := programs[index+3]
			// fmt.Printf("is %v < %v? ", p1, p2)
			programs[idx3] = 0
			if p1 < p2 {
				programs[idx3] = 1
			}
			progression = 4
		case 8:
			p1, p2 := getParams(modes, index, programs)
			idx3 := programs[index+3]
			// fmt.Printf("is %v == %v? will set at index 3 (idx: %v)\n", p1, p2, idx3)
			programs[idx3] = 0
			if p1 == p2 {
				programs[idx3] = 1
			}
			progression = 4
		default:
			log.Fatalf("code unrecognized %v", code)
		}
		// fmt.Println(index, progression)
		index = index + progression
		// l.Println("new index: ", index)
	}
	l := log.New(os.Stdout, fmt.Sprintf("%v: ", letter), 0)

	l.Printf("found 99! @ index %v, closing & done(wg)", index)
	close(outgoing)
	wg.Done()
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
