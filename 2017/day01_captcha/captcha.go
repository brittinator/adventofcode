package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/adventofcode/2017/helpers"
)

// func main() {
// 	file := "../day01_captcha/input"
// 	input := helpers.Input(file)

// 	// part1 := 1
// 	proximity := len(input) / 2
// 	fmt.Println("proximity is ", proximity)
// 	dupes := map[int]int{}
// 	for i := 0; i < len(input)-1; i++ {
// 		num, _ := strconv.Atoi(string(input[i]))

// 		if i+proximity <= len(input)-1 && input[i] == input[i+proximity] {
// 			fmt.Println("matched ", num, input[i+proximity])
// 			// it's a dupe
// 			dupes[num]++
// 		}
// 	}

// 	num, err := strconv.Atoi(string(input[0]))
// 	if err != nil {
// 		log.Fatalf("err : %v", err)
// 	}
// 	if input[0] == input[len(input)-proximity] {
// 		dupes[num]++
// 	}
// 	fmt.Println(dupes)
// 	fmt.Println(addUp(dupes))
// }

func addUp(dupes map[int]int) int {
	var sum int
	if len(dupes) < 1 {
		return 0
	}
	for num, times := range dupes {
		sum += num * times
	}
	return sum
}

func main() {
	//part 2

	file := "../day01_captcha/input"
	input := helpers.Input(file)

	fmt.Println("length is ", len(input))
	proximity := len(input) / 2
	fmt.Println("proximity is ", proximity)
	dupes := map[int]int{}
	for i := 0; i < len(input)-1-proximity; i++ {
		num, _ := strconv.Atoi(string(input[i]))
		if i+proximity <= len(input)-1 && input[i] == input[i+proximity] {
			// it's a dupe
			dupes[num]++
		}
	}
	beginningIndex := 0
	for i := len(input) - proximity; i < len(input)-1; i++ {
		num, _ := strconv.Atoi(string(input[i]))
		if input[i] == input[beginningIndex] {
			fmt.Println("matched ")
			// it's a dupe
			dupes[num]++
		}
		beginningIndex++
	}
	num, err := strconv.Atoi(string(input[0]))
	if err != nil {
		log.Fatalf("err : %v", err)
	}
	if input[0] == input[len(input)-proximity] {
		dupes[num]++
	}
	fmt.Println(dupes)
	fmt.Println(addUp(dupes))
}
