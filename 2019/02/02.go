package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../inputs/02.txt")
	defer f.Close()
	if err != nil {
		log.Fatal("cannot open file", err)
	}

	rawBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var input []int
	nums := strings.Split(string(rawBytes), ",")
	for _, str := range nums {
		num, _ := strconv.Atoi(str)
		input = append(input, num)
	}

	// To do this, before running the program,
	// replace position 1 with the value 12 and
	// replace position 2 with the value 2.
	input[1] = 12
	input[2] = 2
	originalInput := make([]int, len(input))

	for i, v := range input {
		originalInput[i] = v
	}
	// intcode(input)

	// find which pair produce 19690720
	for v1 := 0; v1 < 100; v1++ {
		for v2 := 0; v2 < 100; v2++ {
			input := make([]int, len(originalInput))
			for i, v := range originalInput {
				input[i] = v
			}
			input[1] = v1
			input[2] = v2

			intcode(input)

			if input[0] == 19690720 {
				log.Fatalf(
					"noun %v verb %v 100 * n + v %v\n",
					v1, v2, 100*v1+v2,
				)
			}
		}
	}
	fmt.Println("NOT FOUND :(")

	// fmt.Println("value at position 0: ", input[0])
}

func intcode(programs []int) {
	// put all input into memory slice

	var index int
	for programs[index] != 99 {
		idx1 := programs[index+1]
		idx2 := programs[index+2]
		idx3 := programs[index+3]

		switch programs[index] {
		case 1:
			// add
			val := programs[idx1] + programs[idx2]
			programs[idx3] = val
		case 2:
			// multiply
			val := programs[idx1] * programs[idx2]
			programs[idx3] = val
		case 99:
			break
		default:
			log.Fatal("not a valid process")
		}
		index = index + 4
	}

	// fmt.Println("programs: ", programs)
}
