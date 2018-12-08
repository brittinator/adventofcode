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
	name     int // line number
	x, y     int
	distance int // for part 2
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
	lineNumba := 1
	for scanner.Scan() {
		line := scanner.Text()
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
				name: lineNumba,
				x:    x,
				y:    y,
			})
		lineNumba++
	}

	fmt.Printf("these are the boundaries: \n %v:%v -- %v:%v\n",
		smallestX, smallestY, largestX, largestY)

	// fmt.Printf("locations: ")
	// for i, l := range locations {
	// 	fmt.Printf("\n%v: %v:%v name: %v\n", i, l.x, l.y, l.name)
	// }

	// make a grid to house state of the world
	grid := make([][]location, largestX)
	for i := range grid {
		grid[i] = make([]location, largestY)
	}

	// partOne(locations, grid, smallestX, smallestY, largestX, largestY)
	partTwo(locations, grid, smallestX, smallestY, largestX, largestY)
}

func partOne(knownLocations []location, grid [][]location,
	// find the largest x, y coordinates, these will be the boundaries
	// automatically those points and any with the same x or y coordinates are infinite so those areas are not in the running
	smallestX, smallestY, largestX, largestY int) {
	type areaInfo struct {
		area     int
		areaList []location
	}

	areas := make(map[location]areaInfo)
	// go thru the grid one by one to see who owns that spot
	for xCoor, loc := range grid {
		for yCoor := range loc {
			grid[xCoor][yCoor].x = xCoor
			grid[xCoor][yCoor].y = yCoor

			ownerLocation := findOwner(&grid[xCoor][yCoor], knownLocations)
			// fmt.Printf("the owner of %v:%v is %v\n", xCoor, yCoor, ownerLocation.name)
			if object, found := areas[ownerLocation]; !found {
				areas[ownerLocation] = areaInfo{
					area: 1,
					areaList: []location{
						location{x: xCoor, y: yCoor},
					},
				}
			} else {
				// add this loc to list
				object.areaList = append(object.areaList, location{x: xCoor, y: yCoor})
				object.area++
				areas[ownerLocation] = object
			}
		}
	}

	// fmt.Println("areaList: ")
	// for l, info := range areas {
	// 	fmt.Printf("loc: %v area: %v\n", l.name, info.area)
	// }

	var largestArea int
	for _, areaInfo := range areas {
		// fmt.Println(location, areaInfo)
		if areaInfo.area > largestArea {
			// check to see if area includes an edge
			setAsLargest := true
			for _, locList := range areaInfo.areaList {
				if locList.x == smallestX || locList.x == largestX {
					fmt.Println("area includes an x edge; skipping ", areaInfo.area)
					setAsLargest = false
					break
				}
				if locList.y == smallestY || locList.y == largestY {
					fmt.Println("area includes an Y edge; skipping ", areaInfo.area)
					setAsLargest = false
					break
				}
			}
			if setAsLargest {
				largestArea = areaInfo.area
			}
		}
	}

	fmt.Println("the largest area is ", largestArea)
}

func partTwo(knownLocations []location, grid [][]location,
	smallestX, smallestY, largestX, largestY int) {
	// What is the size of the region containing all locations
	// which have a total distance to all given coordinates of less than 10000?

	// find the size from 1 point to knownLocations
	// add them up, if >= 1000 discard, don't care

	// find the region with points that all have sizes < 1000
	// return the size of the region
	const limit = 10000

	// go thru the grid one by one to see what the total distance is
	for xCoor, loc := range grid {
		for yCoor := range loc {
			grid[xCoor][yCoor].x = xCoor
			grid[xCoor][yCoor].y = yCoor

			distance := findDistance(&grid[xCoor][yCoor], knownLocations)
			if distance >= limit {
				grid[xCoor][yCoor].distance = 0
				continue
			}

			grid[xCoor][yCoor].distance = distance
		}
	}

	// add them up, if >=limit discard, don't care

	// find the region with points that all have sizes <limit
	// return the size of the region
	var size int
	for _, loc := range grid {
		for _, l := range loc {
			fmt.Println(l.distance)
			if l.distance < limit && l.distance != 0 {
				size++
			}
		}
	}

	fmt.Println("size: ", size)
}

// findDistance finds all of the distances from this point to the known locations.
func findDistance(here *location, knownLocations []location) int {
	var totalDistance int
	for _, spot := range knownLocations {
		totalDistance = totalDistance + manhattanDistance(*here, spot)
	}
	fmt.Println("distance: ", totalDistance)
	return totalDistance
}

func findOwner(here *location, knownLocations []location) location {
	distances := make(map[location]int, len(knownLocations))
	// find distance from here to the known locations
	for _, spot := range knownLocations {
		// fmt.Printf("start: %v\nend: %v\n", here, spot)
		distances[spot] = manhattanDistance(*here, spot)
	}

	smallestDistance := math.MaxInt64
	var closest location
	var sawDuplicate bool

	for spot, currentDistance := range distances {
		if currentDistance < smallestDistance {
			sawDuplicate = false
			closest = spot
			smallestDistance = currentDistance
			continue
		}
		if currentDistance == smallestDistance {
			sawDuplicate = true
		}
	}

	if sawDuplicate {
		// there is no clear winner, so return blank location
		return location{}
	}
	return closest
}

func manhattanDistance(start, end location) int {
	xDistance := end.x - start.x
	yDistance := end.y - start.y
	// fmt.Printf("%v:%v == %v\n", start, end, absolute(xDistance) + absolute(yDistance))
	return absolute(xDistance) + absolute(yDistance)
}

func absolute(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}
