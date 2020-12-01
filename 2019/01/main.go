package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal("cannot open file", err)
	}

	var totalFuel int

	// read in input
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		fuel := getFuel(mass)
		totalFuel += fuel
		// get fuel for the fuel mass we just added
		for fuel/3-2 > 0 {
			fuel = getFuel(fuel)
			totalFuel += fuel
		}
	}

	fmt.Println("fuel needed: ", totalFuel)
}

func getFuel(mass int) int {
	return (mass / 3) - 2
}
