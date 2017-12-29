package main

import (
	"fmt"

	"github.com/adventofcode/2017/helpers"
)

func main() {
	fileName := "../day04_passphrases/input"
	fmt.Println("reading ", fileName)
	inputBytes := helpers.Input(fileName)

	var numValid int
	phrases := splitIntoPhrases(inputBytes)
	// fmt.Println(phrases)
	for _, phrase := range phrases {
		if isValid(phrase) {
			numValid++
		}
	}
	fmt.Println("numValid is ", numValid)
}

func splitIntoPhrases(inputBytes []byte) [][]string {
	var passphraseList [][]string
	var word string
	var phrase []string

	for i := 0; i <= len(inputBytes); i++ {
		if i == len(inputBytes) {
			// add last word to phrase
			phrase = append(phrase, word)
			passphraseList = append(passphraseList, phrase)
			continue
		}
		// convert to string
		str := string(inputBytes[i])
		switch str {
		case " ":
			// this is a word!
			fmt.Println("word :", word)
			phrase = append(phrase, word)
			word = ""
		case "\n":
			// add last word to phrase
			phrase = append(phrase, word)
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
	// keys := make([]string, 0, len(phrases))

	for _, word := range phrase {
		if _, ok := phrasesCount[word]; ok {
			return false
		}
		phrasesCount[word] = 1
	}
	return true
}
