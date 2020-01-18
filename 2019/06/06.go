package main

import (
	"fmt"
	"strings"

	"../input"
)

func main() {
	input := input.ReadLineInput("06")
	nodes, root := createTree(input)

	fmt.Println(traverse(nodes[root], nodes))
}

type node struct {
	name     string
	parent   string
	children []*node
	dist     int
}

func newNode(name string) *node {
	return &node{name: name, children: make([]*node, 0)}
}

func createTree(input []string) (map[string]*node, string) {
	allNodes := make(map[string]*node, len(input))
	for _, line := range input {
		names := strings.Split(line, ")")
		parent := names[0]
		child := names[1]
		var parentNode, childNode *node
		if n, ok := allNodes[parent]; ok {
			parentNode = n
		} else {
			parentNode = newNode(parent)
			allNodes[parent] = parentNode
		}
		if n, ok := allNodes[child]; ok {
			childNode = n
		} else {
			childNode = newNode(child)
			allNodes[child] = childNode
		}
		childNode.parent = parent
		parentNode.children = append(parentNode.children, childNode)
	}

	return allNodes, "COM"
}

func traverse(n *node, allNodes map[string]*node) int {
	// do BFS to visit all same-level nodes at the same time
	var total int
	// start at 0
	// add root to queue
	// add to total
	q := queue{}
	q.enqueue(n)

	for !q.empty() {
		curr := q.dequeue()
		total += curr.dist
		// add children to queue
		// find children, add 1+ parent's value
		for _, child := range curr.children {
			child.dist = curr.dist + 1
			q.enqueue(child)
		}
	}

	return total
}

type queue struct {
	n []*node
}

func (q *queue) empty() bool {
	return len(q.n) == 0
}

func (q *queue) dequeue() *node {
	l := len(q.n) - 1
	node := q.n[l]
	q.n = q.n[0:l]

	return node
}

func (q *queue) enqueue(node *node) {
	q.n = append(q.n, node)
}
