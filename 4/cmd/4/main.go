package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var testInput = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

type StringLister struct {
	multiLineString string
	InputLister
}

func (sl StringLister) GetCards() []string {
	// fmt.Println(sl.multiLineString)
	return strings.Split(sl.multiLineString, "\n")
}

func newStringListerFromFile(filepath string) (*StringLister, error) {
	body, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Unable to read file: %v", err)
	}

	return &StringLister{
		multiLineString: strings.TrimSpace(string(body)),
	}, nil
}

func main() {
	testInputter := StringLister{
		multiLineString: testInput,
	}
	testSolution, err := solve(testInputter)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Test input:\n%+v\n\nSolution:\n%d\n", testInputter.GetCards(), testSolution)

	inputter, err := newStringListerFromFile("./input.txt")
	if err != nil {
		panic(err)
	}

	solution, err := solve(inputter)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Real solution:\n%d\n", solution)

	os.Exit(0)
}
