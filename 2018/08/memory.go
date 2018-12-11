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
	f, err := os.Open("./input.txt")
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

	rootNode := partOne(raw)
	partTwo(rootNode)
}

var sum = 0

func partTwo(rootNode *node) {
	var sum int
	// What is the value of the root node?

	// get root's metadata,
	// see if they correspond to nodes
	// get their metadata
	// see if they correspond to nodes
	sum = rootNode.getNodeValue(sum)

	fmt.Println("sum part two: ", sum)

}

func (n *node) getNodeValue(value int) int {
	fmt.Println("node ", n, value)
	/*
	   If a node has no child nodes, its value is the sum of its metadata entries.
	   if a node does have child nodes, the metadata entries become indexes
	   which refer to those child nodes.
	*/
	for _, meta := range n.metadata {
		if n.numChildren == 0 {
			value = value + meta
			continue
		}
		if len(n.children) > meta-1 {
			value = n.children[meta-1].getNodeValue(value)
		}
	}

	return value
}

func partOne(input []int) *node {
	// create root node
	dummyNode := newNode(0, 0)
	buildNodes(input, dummyNode, 0)

	fmt.Println("sum >", sum)
	return dummyNode.children[0]
}

func buildNodes(input []int, currentNode *node, index int) int {
	fmt.Println("current node: ", currentNode.numChildren, currentNode.numMeta,
		currentNode.children, currentNode.metadata)
	thisNode, index := parseHeader(input, index)
	currentNode.children = append(currentNode.children, thisNode)
	currentNode = thisNode
	//call this func again num children times
	for i := 0; i < currentNode.numChildren; i++ {
		index = buildNodes(input, currentNode, index)
	}
	//after recursion returns, do the parsemeta on itself
	index = currentNode.parseMetadata(input, index)
	return index
}

func parseHeader(input []int, startingIndex int) (*node, int) {
	numChildren := input[startingIndex]
	numMeta := input[startingIndex+1]

	return newNode(numChildren, numMeta), startingIndex + 2
}

func (n *node) parseMetadata(input []int, startingIndex int) int {
	numMeta := n.numMeta
	for i := 0; i < numMeta; i++ {
		n.metadata[i] = input[startingIndex+i]
		sum = sum + input[startingIndex+i]
		fmt.Println("adding ", input[startingIndex+i])
	}
	return startingIndex + numMeta
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
