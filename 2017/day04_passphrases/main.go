package main

import (
	"fmt"
	"sort"

	"github.com/adventofcode/2017/helpers"
)

func main() {
	fileName := "../day04_passphrases/input"
	fmt.Println("reading ", fileName)
	inputBytes := helpers.Input(fileName)

	var numValid int
	phrases := splitIntoPhrases(inputBytes, true)
	// fmt.Println(phrases)
	for _, phrase := range phrases {
		if isValid(phrase) {
			numValid++
		}
	}
	fmt.Println("numValid is ", numValid)

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// part two
	var numValidNoAna int
	phrasesAgain := splitIntoPhrases(inputBytes, false)
	for _, phrase := range phrasesAgain {
		if isValid(phrase) {
			numValidNoAna++
		}
	}
	fmt.Println("numValid with no anagrams is ", numValidNoAna)
}

func splitIntoPhrases(inputBytes []byte, wantAnagrams bool) [][]string {
	var passphraseList [][]string
	var word string
	var phrase []string

	for i := 0; i <= len(inputBytes); i++ {
		if i == len(inputBytes) {
			// add last word to phrase
			if !wantAnagrams {
				wordSorted := Sort(word)
				phrase = append(phrase, wordSorted)
			} else {
				phrase = append(phrase, word)
			}
			passphraseList = append(passphraseList, phrase)
			continue
		}
		// convert to string
		str := string(inputBytes[i])
		switch str {
		case " ":
			// this is a word!
			if !wantAnagrams {
				wordSorted := Sort(word)
				phrase = append(phrase, wordSorted)
			} else {
				phrase = append(phrase, word)
			}
			// clean phrase and word for the next one
			word = ""
		case "\n":
			// add last word to phrase
			if !wantAnagrams {
				wordSorted := Sort(word)
				phrase = append(phrase, wordSorted)
			} else {
				// fmt.Println("word :", word)
				phrase = append(phrase, word)
			}
			passphraseList = append(passphraseList, phrase)
			word = ""
			phrase = make([]string, 0)
		default:
			// fmt.Println("character! Adding to str", str)
			word += str
		}
	}
	return passphraseList
}

func isValid(phrase []string) bool {
	phrasesCount := make(map[string]int, len(phrase))

	for _, word := range phrase {
		if _, ok := phrasesCount[word]; ok {
			// fmt.Printf("not valid because %v seen more than once in %v\n", word, phrase)
			return false
		}
		phrasesCount[word] = 1
	}
	return true
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func Sort(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
