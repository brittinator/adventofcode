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
	if err != nil {
		log.Fatal("error reading file", err)
	}
	scanner := bufio.NewScanner(f)

	freq := 0
	for scanner.Scan() {
		num := scanner.Text()
		fmt.Println(num)

		i, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal("trouble converting number ", num, err)
		}
		freq += i
	}

	fmt.Println("frequency: ", freq)
}
