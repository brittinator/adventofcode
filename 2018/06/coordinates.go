package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
		..........
		.A........
		..........
		........C.
		...D......
		.....E....
		.B........
		..........
		..........
		........F.

here A (tied with B) and F (tied with C) are boundaries


How to plot who the closest neighbor is?

Who owns (4,3)?

		..........
		.A........
		..........
		....X....C.
		...D......
		.....E....
		.B........
		..........
		..........
		........F.
Manhattan distance |x1 – x2| + |y1 – y2|


go thru every point and see which is minimized

A (1, 1) -> (4-1) + (3-1) = 5
B (1, 6) -> 6
C (8, 3) -> 4
D (3, 4) -> 2
E (5, 5) -> 3
F (8, 9) -> 10

The shortest distance is to D, so D owns that spot

if value is equal <-> 2 or more no owner

once all of the points are owned, find the areas by counting up A's, B's, C's, etc

but can discard A, B, F, and C (blacklist?) as their boundaries are infinite

*/

type location struct {
	name string
	x, y int
}

func main() {
	f, err := os.Open("./input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal("problem opening file", err)
	}

	scanner := bufio.NewScanner(f)

	locations := make([]location, 0)

	var largestX, largestY int
	smallestX := math.MaxInt64
	smallestY := math.MaxInt64

	// What is the size of the largest area that isn't infinite?

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		x, err := strconv.Atoi(strings.Split(line, ", ")[0])
		if err != nil {
			log.Fatal("conversion ", err)
		}
		y, err := strconv.Atoi(strings.Split(line, ", ")[1])
		if err != nil {
			log.Fatal("conversion ", err)
		}

		if x > largestX {
			largestX = x
		}
		if y > largestY {
			largestY = y
		}

		if x < smallestX {
			smallestX = x
		}
		if y < smallestY {
			smallestY = y
		}

		// build up the location list
		locations = append(locations,
			location{
				name: line,
				x:    x,
				y:    y,
			})
	}

	// these are out boundaries
	fmt.Printf("these are out boundaries: \n %v:%v -- %v:%v\n",
		smallestX, smallestY, largestX, largestY)

	// fmt.Printf("locations: ")
	// for i, l := range locations {
	// 	fmt.Printf("%v: %v\n", i, l)
	// }

	// make a grid to house state of the world
	grid := make([][]location, largestX)
	for i := range grid {
		grid[i] = make([]location, largestY)
	}
	partOne(locations, grid, smallestX, smallestY, largestX, largestY)
	// find the largest x, y coordinates, these will be the boundaries
	// automatically those points and any with the same x or y coordinates are infinite so those areas are not in the running

}

func partOne(knownLocations []location, grid [][]location,
	smallestX, smallestY, largestX, largestY int) {
	areas := make(map[location]int)
	// go thru the grid one by one to see who owns that spot
	for xCoor, loc := range grid {
		for yCoor := range loc {
			grid[xCoor][yCoor].x = xCoor
			grid[xCoor][yCoor].y = yCoor

			// fmt.Println(grid[xCoor][yCoor])
			ownerLocation := findOwner(&grid[xCoor][yCoor], knownLocations)
			// fmt.Printf("the owner of %v:%v is %v\n", xCoor, yCoor, ownerLocation.name)
			if _, found := areas[ownerLocation]; !found {
				areas[ownerLocation] = 1
			} else {
				areas[ownerLocation]++
			}
		}
	}

	// at this time we should know what the largest area is
	var largestArea int
	fmt.Println(largestX, largestY)

	for location, area := range areas {
		fmt.Println(location, area)

		if area > largestArea {
			largestArea = area
		}
	}

	fmt.Println("the largest area is ", largestArea)
	// is any part of this area at an edge?
}

func findOwner(locale *location, knownLocations []location) location {
	distances := make(map[location]int, len(knownLocations))
	for _, spot := range knownLocations {
		// fmt.Printf("start: %v\nend: %v\n", locale, spot)
		distances[spot] = manhattanDistance(*locale, spot)
	}

	// fmt.Println("distances: ", distances)

	smallestDistance := math.MaxInt64
	var closest location
	for locale, currentDistance := range distances {
		if currentDistance < smallestDistance {
			closest = locale
			smallestDistance = currentDistance
		}
	}

	return closest
}

func manhattanDistance(start, end location) int {
	xDistance := end.x - start.x
	yDistance := end.y - start.y
	// fmt.Printf("%v:%v == %v\n", start, end, int(math.Abs(float64(xDistance))+math.Abs(float64(yDistance))))
	// ugly conversions I know
	return int(math.Abs(float64(xDistance)) + math.Abs(float64(yDistance)))
}
