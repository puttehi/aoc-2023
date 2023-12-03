package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func findFirstAndLastNumber(str string) (string, string) {
	var (
		first string
		last  string
	)

	runes := []rune(str)

	for _, rune := range runes {
		if !unicode.IsDigit(rune) {
			continue
		}

		if first == "" {
			first = string(rune)
		}

		last = string(rune)
	}

	return first, last
}

func sumStringDigits(digitStrings []string) (int, error) {
	var sum int

	for _, str := range digitStrings {
		num, err := strconv.Atoi(str)
		if err != nil {
			return sum, err
		}
		sum += num
	}

	return sum, nil
}

func makeDigitStrings(strings []string) []string {
	var digitStrings = make([]string, len(strings))

	for i := range strings {
		first, last := findFirstAndLastNumber(strings[i])
		digitStrings[i] = fmt.Sprintf("%s%s", first, last)
	}

	return digitStrings
}

func solve(input []string) (int, error) {
	digitStrings := makeDigitStrings(input)

	solution, err := sumStringDigits(digitStrings)

	if err != nil {
		return 0, err
	}

	return solution, nil
}
