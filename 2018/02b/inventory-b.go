package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../02a/input1")
	defer f.Close()
	if err != nil {
		log.Fatal("error reading file", err)
	}

	scan := bufio.NewScanner(f)
	var lines []string

	for scan.Scan() {
		l := scan.Text()
		// read each line into a ds,
		lines = append(lines, l)
	}

	// go thru data structure, comparing 2 of each char by char
	for _, l1 := range lines {
		for _, l2 := range lines {
			if l1 == l2 {
				// skip the dupe
				continue
			}
			var hamming int

			var prototype []string
			for i, l1val := range l1 {
				l2val := l2[i]

				if string(l1val) != string(l2val) {
					hamming++
				} else {
					// common letters
					prototype = append(prototype, string(l1[i]))
				}
			}
			// fmt.Println("hamming: ", l1, l2, hamming)

			if hamming == 1 {
				// if hamming distance == 1, these are the strings of interest
				fmt.Println("hamming of one!: ", l1, l2)
				// return the string with the runes in common, minus the different line
				fmt.Println(strings.Join(prototype, ""))

				os.Exit(0)
			}
		}
	}

}
