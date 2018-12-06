package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func main() {
	f, err := os.Open("./input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal("problem opening file", err)
	}

	scanner := bufio.NewScanner(f)

	var polymer string
	for scanner.Scan() {
		polymer = scanner.Text()
	}
	fmt.Println(len(polymer))

	partTwo(polymer)
}

func partTwo(polymer string) {
	// What is the length of the shortest polymer you can produce by
	// removing all units of exactly one type and fully reacting the result?

	polymerSizes := make(map[rune]int, 26)
	for _, letter := range alphabet {
		// replace both upper and lower case letter
		newPolymer := strings.Replace(polymer, string(letter), "", -1)
		newPolymer = strings.Replace(newPolymer, strings.ToUpper(string(letter)), "", -1)
		size := partOne(newPolymer)
		polymerSizes[letter] = size
	}

	smallestSize := len(polymer)
	var letter rune

	for k, v := range polymerSizes {
		if v < smallestSize {
			smallestSize = v
			letter = k
		}
	}
	fmt.Printf("best size using %v: %v\n", string(letter), smallestSize)
}

func partOne(polymer string) int {
	// How many units remain after fully reacting the polymer you scanned?

	// reduce matching pairs until you can do no more
	// [a:A, b:B, B:b, A:a]

	// use a for loop
	// break out when no more matches, match == false

	// when find a match, excise the pair
	// a = append(a[:matchz], a[matchZ+1:]...)
	// make match = true
	dictionary := createDictionary()

	for {
		foundMatch := false

		for _, match := range dictionary {
			newPolymer := strings.Replace(polymer, match, "", -1)
			if newPolymer != polymer {
				// fmt.Println("replaced ", match)
				foundMatch = true
				polymer = newPolymer
			}
		}
		if foundMatch == false {
			break
		}
	}

	fmt.Println("leftover polymer: ", len(polymer))
	return len(polymer)
}

func createDictionary() []string {
	dictionary := make([]string, 0)

	for _, char := range alphabet {
		dictionary = append(
			dictionary,
			strings.ToUpper(string(char))+string(char),
			string(char)+strings.ToUpper(string(char)),
		)
	}
	// fmt.Println(dictionary)
	return dictionary
}
