package main

import (
	"fmt"
	"testing"
)

func TestFindFirstAndLast(t *testing.T) {
	type Want = struct {
		first string
		last  string
	}
	var testCases = []struct {
		name  string
		input string
		want  Want
	}{
		{
			"Simple",
			"1abc2",
			Want{"1", "2"},
		},
		{
			"Many numbers",
			"a1b2c3d4e5f",
			Want{"1", "5"},
		},
		{
			"Empty input should return empty strings",
			"",
			Want{"", ""},
		},
		{
			"Single number should return the same for both",
			"1abcd",
			Want{"1", "1"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			first, last := findFirstAndLastNumber(tc.input)

			if first != tc.want.first || last != tc.want.last {
				t.Errorf("got %s, %s, want %s, %s", first, last, tc.want.first, tc.want.last)
			}
		})
	}
}

func TestMakeDigitStrings(t *testing.T) {
	var testInput = []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	var expected = []string{
		"12",
		"38",
		"15",
		"77",
	}

	digitStrings := makeDigitStrings(testInput)

	for i, s := range digitStrings {
		if s != expected[i] {
			t.Errorf("got %s, want %s", s, expected[i])
		}
	}
}

func TestMutateWordsToDigits(t *testing.T) {
	var testInput = []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	var expected = []string{
		"219",
		"823",
		"123",
		"2134",
		"49872",
		"18234",
		"76",
	}

	testInput = wordsToDigits(testInput)

	for i, s := range testInput {
		if s != expected[i] {
			t.Errorf("got %s, want %s", s, expected[i])
		}
	}
}

func TestSolve(t *testing.T) {
	var testCases = []struct {
		name      string
		input     []string
		isPartTwo bool
		want      int
	}{
		{
			"Part one (no conversions)",
			[]string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			false,
			12 + 38 + 15 + 77,
		},
		{
			"Part two (words to digits)",
			[]string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			true,
			29 + 83 + 13 + 24 + 42 + 14 + 76,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			solution, err := solve(tc.input, tc.isPartTwo)

			if err != nil {
				t.Fatalf("returned error %s", err.Error())
			}

			if solution != tc.want {
				t.Errorf("got %d, want %d", solution, tc.want)
			}

			fmt.Println(tc.input, solution)
		})
	}
}
