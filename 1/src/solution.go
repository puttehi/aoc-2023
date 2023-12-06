package main

import (
	"fmt"
	"strconv"
	"strings"
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

func makeDigitStrings(stringsSlice []string) []string {
	var digitStrings = make([]string, len(stringsSlice))

	for i := range stringsSlice {
		first, last := findFirstAndLastNumber(stringsSlice[i])
		digitStrings[i] = fmt.Sprintf("%s%s", first, last)
	}

	return digitStrings
}

var digitNamesLookup = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// wordsToDigits converts words like "oneightwo" to "182",
// i.e. converts all the matching words to digits,
// returning a copy of the original `stringsSlice` slice
func wordsToDigits(stringsSlice []string) []string {
	digitsSlice := make([]string, len(stringsSlice))
	for i, s := range stringsSlice {
		if s == "" {
			continue
		}

		sortedString := strings.Repeat(" ", len(s))

		// Sort by index
		for digitName, digit := range digitNamesLookup {
			nameIndex := strings.Index(s, digitName)
			digitIndex := strings.Index(s, digit)

			if nameIndex != -1 {
				// Add digit at index (should use rune)
				sortedString = sortedString[:nameIndex] + digit + sortedString[nameIndex+1:]
			}

			if digitIndex != -1 {
				sortedString = sortedString[:digitIndex] + digit + sortedString[digitIndex+1:]
			}
			// NOTE: Enable for a nice visualization of the process
			// fmt.Printf("s: %s (%d characters)\n-> %+v (%d characters)\ndigit name: %s\n%+v\n%+v\n", s, len(s), sortedString, len(sortedString), digitName, nameIndex, digitIndex)
		}

		sortedString = strings.Join(strings.Fields(sortedString), "")
		digitsSlice[i] = sortedString

		// fmt.Printf("%+v\n%+v\n%+v\n", stringsSlice, sortedString, digitsSlice)
	}

	return digitsSlice
}

// solve solves the puzzle.
// Pass in `convertWords=true` for part 2
func solve(input []string, convertWords bool) (int, error) {
	if convertWords {
		input = wordsToDigits(input)
	}
	digitStrings := makeDigitStrings(input)
	solution, err := sumStringDigits(digitStrings)

	if err != nil {
		return 0, err
	}

	/* for i := range input {
		fmt.Printf("%s => %s\n", input[i], digitStrings[i])
	} */

	return solution, nil
}
