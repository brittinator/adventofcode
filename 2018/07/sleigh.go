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

	instructions := make(map[string]step, 26)
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, r := range alpha {
		instructions[string(r)] = newStep(string(r))
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// Step C must be finished before step F can begin.
		prereq := line[5]
		name := line[36]

		oldStep := instructions[string(name)]
		oldStep.prerequisites[string(prereq)] = false
		instructions[string(name)] = oldStep
	}

	fmt.Println("instructions: ", instructions)
	partOne(instructions)

}

func newStep(name string) step {
	return step{
		name:          name,
		prerequisites: make(map[string]bool),
	}
}

type step struct {
	name          string
	prerequisites map[string]bool // true when prereq fulfilled
}

func partOne(instructions map[string]step) {
	orderOfAssembly := []string{}
	// In what order should the steps in your instructions be completed?

	for {
		if len(instructions) == 0 {
			break
		}
		nextInstruction := findNextStep(instructions)
		fmt.Println(nextInstruction)
		orderOfAssembly = append(orderOfAssembly, nextInstruction.name)

		delete(instructions, nextInstruction.name)
		fulfillPrereq(instructions, nextInstruction.name)
	}

	fmt.Println("assembly instructions: ", strings.Join(orderOfAssembly, ""))
}

// fulfillPrereq goes through the instruction list and sets the
// instruction in each other steps to true.
func fulfillPrereq(instructions map[string]step, name string) {
	for stepName, stepInfo := range instructions {
		for prereq := range stepInfo.prerequisites {
			if prereq == name {
				// set it to true, it's done.
				stepInfo.prerequisites[name] = true
				instructions[stepName] = stepInfo
			}
		}
	}
}

func findNextStep(instructions map[string]step) step {
	nextStep := step{name: "Z"}

	// find all steps with no prereqs
	for stepName, stepInfo := range instructions {
		// compare their name order, earlier alpha goes first
		// fmt.Println(len(stepInfo.prerequisites), stepName < nextStep.name)
		if len(stepInfo.prerequisites) == 0 && stepName < nextStep.name {
			// this is the first step
			nextStep = stepInfo
			continue
		}
		allPrereqsDone := true
		for _, v := range stepInfo.prerequisites {
			if v == false {
				allPrereqsDone = false
			}
		}
		if allPrereqsDone && stepName < nextStep.name {
			nextStep = stepInfo
		}
	}

	// fmt.Println("find next: ", nextStep)
	return nextStep
}
