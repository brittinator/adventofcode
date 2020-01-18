package main

import (
	"fmt"
	"strings"

	"../input"
)

func main() {
	input := input.ReadLineInput("06")
	nodes, _ := createTree(input)

	// fmt.Println(traverse(nodes[root], nodes))

	fmt.Println(findSantaDistance(nodes))
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

func findSantaDistance(allNodes map[string]*node) int {
	q := queue{}
	// start at YOU's parent or child, dist 0
	// BFS until find SAN, then subtract 1 from dist
	you := allNodes["YOU"]
	start := allNodes[allNodes["YOU"].parent]
	q.enqueue(start)
	for _, child := range you.children {
		q.enqueue(child)
	}
	visited := make(map[string]struct{}, 0)
	visited["YOU"] = struct{}{}

	for !q.empty() {
		curr := q.dequeue()
		if _, ok := visited[curr.name]; ok {
			continue
		}
		visited[curr.name] = struct{}{}
		if curr.name == "SAN" {
			return curr.dist - 1
		}
		parent := allNodes[curr.parent]
		parent.dist = curr.dist + 1
		q.enqueue(parent)
		for _, child := range curr.children {
			child.dist = curr.dist + 1
			q.enqueue(child)
		}
	}

	return -1

}
