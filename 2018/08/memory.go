package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input0.txt")
	if err != nil {
		log.Fatal("failed to open", err)
	}
	scanner := bufio.NewScanner(f)

	var raw []int
	for scanner.Scan() {
		// entire file
		strArray := strings.Split(scanner.Text(), " ")
		raw = make([]int, len(strArray))
		for i, strNum := range strArray {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				log.Fatal("cannot convert ", strNum, err)
			}
			raw[i] = num
		}
	}

	fmt.Println(len(raw))
	partOne(raw)

	// take the 1st 2 numbas
	// pass into node as headers
	// pop those off array
	// for numchildren,
	// create node with the headers for the headers

	// to create a node,
	// first part (2 numbers) are the headers
	// add that to a new node

	// end condition is either you've reached all the children or then numchildren is zero

}

var sum = 0

func partOne(input []int) {
	// create root node
	rootNode := newNode(input[0], input[1])
	currentNode := rootNode

	pointer := 2
	for i := 0; i < currentNode.numChildren; i++ {
		pointer = buildNodes(input, currentNode, pointer)
	}
	currentNode.parseMetadata(input, pointer)
	fmt.Println(currentNode)
	fmt.Println(sum)
}

func buildNodes(input []int, currentNode *node, pointer int) int {
	if input[pointer] == 0 {
		// first move index +1
		pointer++
		// create a new node
		numMeta := input[pointer]
		thisNode := newNode(0, numMeta)
		currentNode.children = append(currentNode.children, thisNode)

		// reset current node to thisNode
		currentNode = thisNode
		pointer++
		currentNode.parseMetadata(input, pointer)

		pointer = pointer + numMeta
		return pointer
	}

	// it's talking about children or meta of other nodes
	numChildren := input[pointer]
	numMeta := input[pointer+1]
	// make a new node
	thisNode := newNode(numChildren, numMeta)
	currentNode.children = append(currentNode.children, thisNode)
	currentNode = thisNode
	pointer = pointer + 2
	//call this func again numchildren times
	for i := 0; i < currentNode.numChildren; i++ {
		pointer = buildNodes(input, currentNode, pointer)
	}
	//after recursion returns, do the parsemeta on itself
	currentNode.parseMetadata(input, pointer)
	pointer++
	return pointer

}

func (n *node) parseMetadata(input []int, startingIndex int) {
	numMeta := n.numMeta
	for i := 0; i < numMeta; i++ {
		n.metadata[i] = input[startingIndex+i]
		sum = sum + input[startingIndex+i]
		fmt.Println("adding ", input[startingIndex+i])
	}
	fmt.Println("meta done: ", n.metadata)
}

// node is a single node, but can have children.
// root will contain everything.
type node struct {
	numChildren, numMeta int
	children             []*node
	metadata             []int
}

func newNode(numChildren, numMeta int) *node {
	return &node{
		numChildren: numChildren,
		numMeta:     numMeta,
		children:    make([]*node, 0),
		metadata:    make([]int, numMeta),
	}
}
