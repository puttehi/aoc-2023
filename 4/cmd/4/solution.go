package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"regexp"
	"strconv"
	"strings"
)

type InputLister interface {
	GetCards() []string // ["Card ####: ## ## ## ## ## | ## ## 0# ## ## 0# ## ##"]
}

type Solution = int

func getWinningNumbers(card string) ([]int, error) {
	parts := strings.Split(card, "|")
	if len(parts) != 2 {
		return []int{}, fmt.Errorf("Card split to more than two parts: %d", len(parts))
	}

	leftParts := strings.Split(parts[0], ":")
	if len(leftParts) != 2 {
		return []int{}, fmt.Errorf("Left side of card split to more than two parts: %d", len(parts))
	}

	leftNumbers := leftParts[1] // 0 is "Card ##:"
	right := parts[1]

	//cardNoRegex := regexp.MustCompile(`Card (\d+):`)
	numbersRegex := regexp.MustCompile(`(\d+)`)

	//cardNo := cardNoRegex.Find([]byte(left))
	numbersBytes := numbersRegex.FindAll([]byte(leftNumbers), -1)
	if numbersBytes == nil {
		return []int{}, fmt.Errorf("Could not parse chosen numbers from %+v", leftNumbers)
	}

	numbers := make([]int, len(numbersBytes))
	for i, numberBytes := range numbersBytes {
		number, err := strconv.Atoi(string(numberBytes))
		if err != nil {
			return []int{}, fmt.Errorf("Numbers bytes did not parse to an int from %+v: %s", numbersBytes, err.Error())
		}
		numbers[i] = number
	}

	lottoNumbersBytes := numbersRegex.FindAll([]byte(right), -1)
	if lottoNumbersBytes == nil {
		return []int{}, fmt.Errorf("Could not parse chosen numbers from %+v", right)
	}

	var winning []int
	for _, lottoBytes := range lottoNumbersBytes {
		lottoNumber, err := strconv.Atoi(string(lottoBytes))
		if err != nil {
			return winning, fmt.Errorf("Lotto bytes did not parse to an int from %+v: %s", lottoBytes, err.Error())
		}

		// fmt.Printf("winning so far:%+v\nnumbers: %+v\ncomparing: %d\n", winning, numbers, lottoNumber)
		if !slices.Contains(numbers, lottoNumber) {
			continue
		}

		// fmt.Println("hit!")

		winning = append(winning, lottoNumber)
	}

	return winning, nil
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func solve(inputLister InputLister) (Solution, error) {
	input := inputLister.GetCards()

	solution := 0

	for _, card := range input {
		winning, err := getWinningNumbers(card)
		if err != nil {
			return solution, err
		}

		cardPts := 1 << len(winning) / 2
		// fmt.Printf("Total card pts for %d cards: %d\n", len(winning), cardPts)
		solution += cardPts
	}

	return solution, nil
}
