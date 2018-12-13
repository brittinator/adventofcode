package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal("cannot open file", err)
	}

	var initialState []string
	rules := make(map[string]string)

	// read in input
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		const init = "initial state: "
		if strings.Contains(line, init) {
			initialState = strings.Split(strings.TrimPrefix(line, init), "")
		}
		const rocket = " => "
		if strings.Contains(line, rocket) {
			l := strings.Split(line, rocket)
			rules[l[0]] = l[1]
		}
	}

	partOne(initialState, rules)
}

func partOne(state []string, rules map[string]string) {
	// add a few dots to the end so it will have
	// a complete picture.

	// push Front/Unshift
	// a = append([]T{x}, a...)
	const space = 80
	spacer := make([]string, space)
	for i := 0; i < space; i++ {
		spacer[i] = "."
	}
	state = append(spacer, state...)
	state = append(state, spacer...)
	start := space

	fmt.Println("len ", len(state))
	for i := 0; i < 20; i++ {
		fmt.Printf("state>> i: %v\n%v\n", i, strings.Join(state, ""))
		state = generation(state, rules)
	}

	fmt.Printf("FINAL len: %v\n%v   %v\n", len(state), strings.Join(state, ""), start)
	addPots(state, start)
}

func generation(beginningState []string, rules map[string]string) []string {
	newState := []string{".", "."}
	for i := 2; i < len(beginningState)-3; i++ {
		// build up what the local area looks like,
		// i-2,i-1,i,i+1,i+2
		toMatch :=
			beginningState[i-2] +
				beginningState[i-1] +
				beginningState[i] +
				beginningState[i+1] +
				beginningState[i+2]
		if v, ok := rules[toMatch]; !ok {
			// newState = append(newState, ".")
			fmt.Println("panic ", newState)
			log.Panic("rule not found for ", i, toMatch)
		} else {
			newState = append(newState, v)
		}
	}

	return append(newState, ".", ".")
}

func addPots(state []string, start int) int {
	// split into positive and negative
	fmt.Printf("Add pots len: %v start: %v\n", len(state), start)
	fmt.Println(strings.Join(state, ""))
	for j := 0; j < 12; j++ {
		for i := 0; i < 10; i++ {
			fmt.Printf("%v", i)
		}
	}
	fmt.Println(state[start-3], state[start-2], state[start-1], state[start], state[start+1])
	var pos []string
	for i := start; i < len(state)-1; i++ {
		pos = append(pos, state[i])
	}

	var neg []string
	for i := start; i >= 0; i-- {
		neg = append(neg, state[i])
	}

	var sum int

	for i, val := range pos {
		if val == "#" {
			sum = sum + i
		}
	}

	for i, val := range neg {
		if val == "#" {
			sum = sum - i
		}
	}

	fmt.Println("value: ", sum)
	return sum
}
