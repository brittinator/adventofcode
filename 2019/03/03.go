package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../inputs/03.txt")
	defer f.Close()
	if err != nil {
		log.Fatal("cannot open file", err)
	}

	rawBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawBytes), "\n")
	w1 := lines[0]
	w2 := lines[1]
	// w1 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	// w2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	// fmt.Println("> ", w1, w2)
	path1 := plot(w1)
	path2 := plot(w2)

	point := findClosestIntersection(path1, path2)

	fmt.Println("closest is: ", point, manhattanDistance(point.x, point.y))

	shortest := findShortestIntersection(path1, path2)

	fmt.Println("shortest is: ", shortest)

}

type coor struct {
	x, y int
}

func plot(fullDirections string) map[coor]int {
	directions := strings.Split(fullDirections, ",")
	path := make(map[coor]int)

	var steps int
	current := coor{x: 0, y: 0}
	for _, dir := range directions {
		times, _ := strconv.Atoi(dir[1:])
		switch {
		case dir[0] == 'U':
			for i := 0; i < times; i++ {
				steps++
				current.x++
				path[current] = steps
			}
		case dir[0] == 'D':
			for i := 0; i < times; i++ {
				steps++
				current.x--
				path[current] = steps
			}
		case dir[0] == 'L':
			for i := 0; i < times; i++ {
				steps++
				current.y++
				path[current] = steps
			}
		case dir[0] == 'R':
			for i := 0; i < times; i++ {
				steps++
				current.y--
				path[current] = steps
			}
		}
	}

	// for _, v := range path {
	// fmt.Println(v)
	// }
	return path
}

func findClosestIntersection(path1, path2 map[coor]int) coor {
	closest := coor{x: math.MaxInt16, y: math.MaxInt16}

	for point := range path1 {
		if _, found := path2[point]; found {
			// intersection
			if manhattanDistance(point.x, point.y) < manhattanDistance(closest.x, closest.y) {
				closest = point
			}
		}
	}

	return closest
}

func findShortestIntersection(path1, path2 map[coor]int) int {
	shortest := math.MaxInt16
	for p1, dist1 := range path1 {
		if dist2, found := path2[p1]; found {
			// intersection
			distance := dist1 + dist2
			if distance < shortest {
				shortest = distance
			}
		}
	}

	return shortest
}

// plot first path
// plot 2nd path, comparing with 1st path
// if in both
// see if manhattanDis < min, if so make min == new manDist

func manhattanDistance(x, y int) int {
	if x < 0 {
		x = -1 * x
	}
	if y < 0 {
		y = -1 * y
	}
	return x + y
}
