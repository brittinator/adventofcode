package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTree(t *testing.T) {
	input := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	nodes, root := createTree(input)
	fmt.Println("root", root)
	for _, v := range nodes {
		fmt.Printf("%v, parent %v len child %v\n", v.name, v.parent, len(v.children))
	}

	assert.Equal(t, traverse(nodes[root], nodes), 42)
}
