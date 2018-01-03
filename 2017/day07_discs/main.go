package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// oenxsfm (275) -> cuhrgp, qlgme, bbhyth, dnzdqz, ezdhr
// name (wt) (optional) -> child1, child2
// names len are variable
// wt is in parens
// children are after '->'

// form up a node as below

//  form a graph
// make current name root
// make wt attr on the node
// (if any) add children as children

// go to the first child of node
// search hash for the child
// repeat forming of graph

// Graph is the whole representation of the spinning discs field.
type Graph struct {
	Nodes map[string]*Node
}

// Node is an individual 'program'.
type Node struct {
	Children   map[string]string
	Name       string
	Parent     string // will be string empty if unknown or no parent (which means it's the root)
	PathLength int
	Weight     int
}

func newNode(name string, wt int) Node {
	return Node{
		Children:   make(map[string]string, 0),
		Name:       name,
		PathLength: 1,
		Weight:     wt,
	}
}

func (g *Graph) addNodeToGraph(n *Node) {
	if _, found := g.Nodes[n.Name]; found {
		return
	}
	g.Nodes[n.Name] = n
}

func (g *Graph) parseAsNode(line string) Node {
	// fmt.Println("parseAsNode ", line)
	nodeInfo := strings.FieldsFunc(line, Split)
	// splits into ex: len=6 [fwft  72    ktlj  cntj  xhth]

	wt, err := strconv.Atoi(nodeInfo[1])
	if err != nil {
		log.Fatal(err)
	}
	n := newNode(nodeInfo[0], wt)
	if len(nodeInfo) > 2 {
		// add info about children
		children := map[string]string{}
		for i := 3; i < len(nodeInfo); i++ {
			name := nodeInfo[i]
			n.Children[name] = name
		}
		n.Children = children
	}
	g.addNodeToGraph(&n)
	// fmt.Println("n is: ", n)
	return n
}

func Split(r rune) bool {
	return r == ',' || r == '-' || r == '>' || r == '(' || r == ')' || r == ' '
}

func (g *Graph) findLength(nde string) Node {
	if len(g.Nodes[nde].Children) <= 0 {
		return *g.Nodes[nde]
	}
	nodes := g.Nodes
	actual := nodes[nde]
	actual.PathLength++
	for childName := range g.Nodes[nde].Children {
		g.findLength(childName)
	}
	return Node{}
}

func addToMap(mapKey map[string]int, line string) map[string]int {
	lineSplit := strings.FieldsFunc(line, Split)
	// splits into ex: len=6 [fwft  72    ktlj  cntj  xhth]
	fmt.Println(lineSplit)
	for _, item := range lineSplit {
		item = strings.TrimSpace(item)
		if strings.ContainsAny(item, "->() 1234567890") {
			fmt.Println("character ", item, len(item))
			continue
		}
		fmt.Println("character>> ", item)
		// add to map
		if _, found := mapKey[item]; found {
			mapKey[item]++
			continue
		}
		mapKey[item] = 1
	}
	return mapKey
}

func main() {
	fileName := "input"
	base, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(base)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// g := Graph{Nodes: make(map[string]*Node, 10)}

	scanner := bufio.NewScanner(file)
	mapKey := make(map[string]int, 0)
	for scanner.Scan() {
		// separates out by new line
		line := scanner.Text()
		fmt.Println("parsing", line)
		mapKey = addToMap(mapKey, line)
		// node := g.parseAsNode(line)
		// g.Nodes[node.Name] = &node
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(mapKey)

	// find the key with value of only 1
	for node, v := range mapKey {
		if v == 1 {
			fmt.Println("base is: ", node)
			return
		}
	}
}
