package main

import (
	"fmt"
	"strconv"
)

func main() {
	var meetsDemand int

	for num := 136818; num <= 685979; num++ {
		if meetsRules(number{
			n:   num,
			str: strconv.Itoa(num),
		}) {
			fmt.Println("meets rules: ", num)
			meetsDemand++
		}
	}
	// fmt.Println(meetsRules(
	// 	number{
	// 		n:   111122,
	// 		str: "111122",
	// 	},
	// ))

	fmt.Println("number that meets the demands: ", meetsDemand)
}

type number struct {
	n   int
	str string
}

func meetsRules(num number) bool {
	if len(num.str) != 6 {
		// fmt.Printf("%v is not 6 digits in length", num)
		return false
	}
	// if num.n < 136818 || num.n > 685979 {
	// 	// fmt.Printf("%v is not in range", num)
	// 	return false
	// }

	var prev number
	valueCount := make(map[int]int)

	for i, c := range num.str {
		n, _ := strconv.Atoi(string(c))
		valueCount[n]++
		if i > 0 {
			// if prev.str == string(c) {
			// 	hasDupe = true
			// }
			if prev.n > n {
				// fmt.Printf("%v is not always incrementing (at char %v, %v)", num, i, string(c))
				return false
			}
		}
		prev = number{n: n, str: string(c)}

	}

	var hasDupe bool
	for _, count := range valueCount {
		if count == 2 {
			hasDupe = true
		}
	}

	if !hasDupe {
		// fmt.Printf("%v no duplicate", num)
		return false
	}

	return true
}
