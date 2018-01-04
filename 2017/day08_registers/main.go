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

func dealWithLine(l string, r map[string]int) map[string]int {
	// ehd dec -241 if fw != 6
	// xn dec -416 if a > -2
	split := strings.Split(l, " ")
	// fmt.Println("====")
	// fmt.Println(split)
	k1 := split[0]
	fxn := split[1] // inc or dec
	v1, _ := strconv.Atoi(split[2])
	// 'if' is 3
	k2 := split[4]
	cond := split[5]
	v2, _ := strconv.Atoi(split[6])

	if _, found := r[k1]; !found {
		// add k to registry
		r[k1] = 0
	}
	if _, found := r[k2]; !found {
		// add k to registry
		r[k2] = 0
	}
	// fmt.Println("registry: ", r)
	var doFxn bool
	// fmt.Println(r[k2], cond, v2)
	switch cond {
	case "<":
		if r[k2] < v2 {
			// fmt.Println("a")
			doFxn = true
		}
	case "<=":
		if r[k2] <= v2 {
			// fmt.Println("b")

			doFxn = true
		}
	case ">":
		if r[k2] > v2 {
			// fmt.Println("c")

			doFxn = true
		}
	case ">=":
		if r[k2] >= v2 {
			// fmt.Println("d")

			doFxn = true
		}
	case "==":
		if r[k2] == v2 {
			// fmt.Println("e")

			doFxn = true
		}
	case "!=":
		if r[k2] != v2 {
			// fmt.Println("f")

			doFxn = true
		}
	}

	if doFxn {
		// fmt.Println("doing fxn")
		if fxn == "dec" {
			v1 = -1 * v1
		}
		r[k1] = r[k1] + v1
	}

	return r
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

	registry := make(map[string]int, 1000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		registry = dealWithLine(line, registry)
	}
	fmt.Println("registry", registry)

	// What is the largest value in any register after completing the instructions in your puzzle input?
	var hi struct {
		k string
		v int
	}
	for k, v := range registry {
		if v > hi.v {
			hi.v = v
			hi.k = k
		}
	}

	fmt.Printf("Highest value after is %v from %v\n", hi.v, hi.k)
}
