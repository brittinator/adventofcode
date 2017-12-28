package main

func main() {

}

func populate(endNum int) [][]int {
	/*
		Using these sets of rules, populate a anti-clockwise
		grid of incrementing numbers in a spiraling outwards fashion.
	*/

	// if incrementing x, stop when (x, y+1) is empty

	// if incrementing y, stop when (x-1,y) is empty

	// if decrementing x, stop when (x, y-1) is empty

	// if decrementing y, stop when (x+1, y) is empty

	// stop all when num == endNum
}
