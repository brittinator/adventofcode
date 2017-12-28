package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func Input(filename string) []byte {
	base, err := filepath.Abs(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("reading from ", base)
	// read the whole file at once
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return input
}
