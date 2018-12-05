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
	f, err := os.Open("./input")
	defer f.Close()
	if err != nil {
		log.Fatal("error reading file", err)
	}

	// create a 2D array of a 1000x1000 square
	// initialize all with 0
	fabric := createSquareArray(1000)

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		// go thru each line
		l := scanner.Text()
		lines = append(lines, l)
		// ditch everything preceeding @

		// separate by :
		dimensions := strings.Split(strings.Split(l, "@ ")[1], ": ")
		// fmt.Println(l, dimensions)

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
				// fmt.Printf("%v: at %v:%v -->%v\n", j, x, y, fabric.array[x][y])
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

	// find the non-overlapping claim
	for _, l := range lines {
		// grab the id
		id := strings.Split(l, " @")[0]
		// fmt.Println("next scan ", id)

		// separate by :
		dimensions := strings.Split(strings.Split(l, "@ ")[1], ": ")
		// fmt.Println(l, dimensions)

		// left is how far from edge
		// right is number of spaces to fill
		xOffset, _ := strconv.Atoi(strings.Split(dimensions[0], ",")[0])
		yOffset, _ := strconv.Atoi(strings.Split(dimensions[0], ",")[1])
		xWidth, _ := strconv.Atoi(strings.Split(dimensions[1], "x")[0])
		yHeight, _ := strconv.Atoi(strings.Split(dimensions[1], "x")[1])

		if !fabric.doIOverlap(xOffset, yOffset, xWidth, yHeight) {
			fmt.Println(id)
			os.Exit(0)
		}
	}
}

func (f *fabric) doIOverlap(xStart, yStart, width, height int) bool {
	for y := yStart; y < yStart+height; y++ {
		for x := xStart; x < xStart+width; x++ {
			// fmt.Println(x, y)
			if f.array[x][y] >= 2 {
				// fmt.Println("true: ", x, y, f.array[x][y])
				return true
			}
			// fmt.Println("false: ", x, y, f.array[x][y])
		}
	}

	return false
}
