package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./input1")
	defer f.Close()
	if err != nil {
		log.Fatal("error reading file", err)
	}

	scan := bufio.NewScanner(f)

	var double, triple int
	for scan.Scan() {
		// read line
		line := scan.Text()
		alpha := make(map[rune]int, len(line))
		for _, r := range line {
			if i, found := alpha[r]; found {
				// increment by one
				alpha[r] = i + 1
			} else {
				alpha[r] = 1
			}
		}

		fmt.Println("alpha: ", alpha)
		// increment no more than once per line
		var doneDouble, doneTriple bool
		for key, val := range alpha {

			// if double, add to double
			if val == 2 && !doneDouble {
				fmt.Println("incrementing double because ", key)
				double++
				doneDouble = true
			}
			// if triple, add to triple
			if val == 3 && !doneTriple {
				fmt.Println("incrementing triple because ", key)
				triple++
				doneTriple = true
			}
		}

		fmt.Println("double triple values: ", double, triple)
	}

	fmt.Println("double triple values: ", double, triple)

	// at end, multiple numbers together
	finalNum := double * triple
	fmt.Println("checksum: ", finalNum)
}
