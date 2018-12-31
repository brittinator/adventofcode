package main

import (
	"fmt"
	"log"
)

func main() {
	s := score{
		score:      []int{3, 7},
		elf1Locale: 0,
		elf2Locale: 1,
	}

	for {
		if len(s.score) == 824501+10 {
			break
		}
		fmt.Println("number of recipes: ", len(s.score))
		// create new recipes
		recipes := s.newRecipes()
		// fmt.Println(recipes)

		// add recipes to score
		s.addRecipes(recipes)
		// fmt.Println(s.score)
		// move the elfs
		s.moveElf1()
		s.moveElf2()

		// fmt.Println("new locations: ", s.elf1Locale, s.elf2Locale)
	}

	fmt.Println(s.score[824501:])
}

type score struct {
	score []int
	// in array counting
	elf1Locale int
	elf2Locale int
}

func (s *score) newRecipes() int {
	r1 := s.score[s.elf1Locale]
	r2 := s.score[s.elf2Locale]
	return r1 + r2
}

func (s *score) addRecipes(recipes int) {
	if recipes < 10 {
		s.score = append(s.score, recipes)
		return
	}
	// split into digits
	if recipes < 100 {
		// there are two digits
		r1 := recipes / 10 % 10
		r2 := recipes % 10
		s.score = append(s.score, r1, r2)
		return
	}
	if recipes < 1000 {
		r1 := recipes / 100 % 10
		r2 := recipes / 10 % 10
		r3 := recipes % 10
		s.score = append(s.score, r1, r2, r3)
		return
	}
	if recipes >= 1000 {
		log.Fatal("this recipes is huge: ", recipes)
	}
}

func (s *score) moveElf1() {
	n := s.score[s.elf1Locale] + 1
	// fmt.Println("want to move spaces: ", n)

	s.elf1Locale = s.move(n, s.elf1Locale)
}

func (s *score) move(numToMove int, currentLoc int) int {
	return (numToMove + currentLoc) % len(s.score)
}

func (s *score) moveElf2() {
	n := s.score[s.elf2Locale] + 1
	// fmt.Println("want to move spaces: ", n)

	s.elf2Locale = s.move(n, s.elf2Locale)
}
