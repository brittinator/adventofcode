package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// fabric is the square piece of fabric for Santa.
type fabric struct {
	dimension         int // both height and width
	array             [][]int
	overlappingClaims int
}

func createSquareArray(size int) fabric {
	f := make([][]int, 1000)
	for i := range f {
		f[i] = make([]int, 1000)
	}

	return fabric{
		dimension: 1000,
		array:     f,
	}
}

func main() {
	// a1 a2 a3 a4 a5
	// b1 b2 b3 b4 b5
	// c1 c2 c3 c4 c5
	// [[a1,b1,c1]][[a2,b2,c2]][[a3,b3,c3]][[a4,b4,c4]][[a5,b5,c5]]
	// so for #123 @ 3,1: 2,2 it would fill b4,b5,c4,c5
	f, err := os.Open("./input")
	defer f.Close()
	if err != nil {
		log.Fatal("error reading file", err)
	}

	// create a 2D array of a 1000x1000 square
	// initialize all with 0
	fabric := createSquareArray(1000)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// go thru each line
		l := scanner.Text()
		// ditch everything preceeding @

		// separate by :
		dimensions := strings.Split(strings.Split(l, "@ ")[1], ": ")
		fmt.Println(l, dimensions)

		// left is how far from edge
		// right is number of spaces to fill
		xOffset, err := strconv.Atoi(strings.Split(dimensions[0], ",")[0])
		if err != nil {
			log.Fatal(err)
		}
		yOffset, err := strconv.Atoi(strings.Split(dimensions[0], ",")[1])
		if err != nil {
			log.Fatal(err)
		}
		xWidth, err := strconv.Atoi(strings.Split(dimensions[1], "x")[0])
		if err != nil {
			log.Fatal(err)
		}
		yHeight, err := strconv.Atoi(strings.Split(dimensions[1], "x")[1])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(xOffset, yOffset, xWidth, yHeight)
		for i := 0; i < yHeight; i++ {
			y := yOffset + i
			// fmt.Println(i, xOffset, y)
			for j := 0; j < xWidth; j++ {
				x := xOffset + j
				fmt.Printf("%v: at %v:%v -->%v\n", j, x, y, fabric.array[x][y])
				fabric.array[x][y]++
			}
		}
	}

	for i := range fabric.array {
		for j := range fabric.array[i] {
			if fabric.array[i][j] > 1 {
				fabric.overlappingClaims++
			}
		}
	}

	// return tally
	fmt.Println("Overlapping: ", fabric.overlappingClaims)

	// to fill:

	// index is num-1, width
	// length is down,or inner index is num-1

	// when a space is filled, increment
	// if the number is >= 2, add to a tally

}
