package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readUserInput(prompt string, defaultValue string) (string, error) {
	ret := defaultValue

	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return ret, err
	}

	if line != "\n" {
		ret = line
	}

	return ret, nil
}

func readFile(filePath string) ([]string, error) {
	var lines []string

	f, err := os.Open(filePath)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines, nil
}

func main() {
	inputFile, err := readUserInput("Please enter absolute path to input file (./input.txt):", "./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Using input file at %s\n", inputFile)

	inputLines, err := readFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	solutionOne, err := solve(inputLines, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Solution for part 1 is: %d (counted %d strings)\n", solutionOne, len(inputLines))

	solutionTwo, err := solve(inputLines, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Solution for part 2 is: %d (counted %d strings)\n", solutionTwo, len(inputLines))
}
