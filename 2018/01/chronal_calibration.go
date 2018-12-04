package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input1")
	defer f.Close()

	if err != nil {
		log.Fatal("error reading file", err)
	}
	scanner := bufio.NewScanner(f)

	var freq int

	var changes []int

	seenFreaks := map[int]int{}
	seenFreaks[0] = 0

	for scanner.Scan() {
		num := scanner.Text()
		fmt.Printf("%v + %v\n", freq, num)

		i, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal("trouble converting number ", num, err)
		}

		// for step 2
		for _, value := range changes {
			if value == i {
				fmt.Println("found repeating sequence: ", changes)
				break
			}
		}
		fmt.Println("adding to changes")
		changes = append(changes, i)

		freq += i
	}

	// fmt.Println("frequency: ", freq)

	// fmt.Println("for step 2 changes: ", changes, len(changes))
	// step 2

	seenFreaks = map[int]int{}
	seenFreaks[0] = 0
	var repeatFreq int

	// os.Exit(0)
	for {
		fmt.Println("a")
		for _, num := range changes {
			// fmt.Println("b")
			// fmt.Printf("%v + %v\n", repeatFreq, num)
			repeatFreq += num

			if _, found := seenFreaks[repeatFreq]; found {
				// fmt.Println("seenFreaks> ", seenFreaks)
				// fmt.Println("changes> ", changes)
				fmt.Println("found first repeat of frequency: ", repeatFreq)
				os.Exit(0)
			} else {
				// add it to the map
				seenFreaks[repeatFreq] = repeatFreq
			}

			// fmt.Println("c")
		}
	}
}
