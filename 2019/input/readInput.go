package input

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadNumberInput ...
func ReadNumberInput(day string) []int {
	f, err := os.Open(fmt.Sprintf("../input/%v.txt", day))
	defer f.Close()
	if err != nil {
		log.Fatal("cannot open file", err)
	}

	rawBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	numStr := strings.Split(string(rawBytes), ",")
	input := make([]int, len(numStr))

	for i, str := range numStr {
		num, _ := strconv.Atoi(str)
		input[i] = num
	}

	return input
}
