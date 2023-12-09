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
		return []int{}, fmt.Errorf("Card split to wrong amount of parts: %d, card: %s", len(parts), card)
	}

	leftParts := strings.Split(parts[0], ":")
	if len(leftParts) != 2 {
		return []int{}, fmt.Errorf("Left side of card split to wrong amount of parts: %d, left side: %s", len(leftParts), parts[0])
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

// Recursive solution:
// - Copy func input list (card list)
// - Go through copy
// - Builds list of copy cards in a list
// - When done, recurse back with the new, current copy
// In the end, copies of copies of copies of copies should be returned when popping stack
// and it should be possible to len(copy of input)
func recursive(allCards []string, input []string, previousIndex int) []string {
	total := make([]string, len(input))
	copy(total, input) // Always first 25 cards are included in total

	for i, card := range input {
		copies := []string{}
		winning, err := getWinningNumbers(card)

		// cardInfo := fmt.Sprintf("Original card at %d: %s\nwon: %d card copies\npreviousIndex: %d", i, card, len(winning), previousIndex)
		// infoLen := len(cardInfo)
		// fmt.Println(strings.Repeat("-", infoLen))
		// fmt.Println(cardInfo)
		// fmt.Println(strings.Repeat("\\/", int(infoLen/2)))

		if err != nil {
			fmt.Printf("Couldn't determine winning numbers for card %s at index %d: %s\n", card, i, err.Error())
			return []string{}
		}

		winningCount := len(winning)
		startIndex := previousIndex + 1
		for j := startIndex; j < startIndex+winningCount; j++ {
			// fmt.Println("ALL:\t\t", allCards)
			// fmt.Printf("Copying winning card and recursing...\n  at %d:\t  %s\n", j+1, allCards[j+1])
			// fmt.Println("From:\t", copies)
			copies = append(copies, allCards[j+1]) // j needs to be in allcards index range
			// fmt.Println("To:\t", copies)
		}
		// fmt.Printf("Oh no Pooh, you are eating recursion!:\nRecursing to:\t--> %s <<-\n", strings.Join(copies, ",\n\t\t    "))
		total = append(total, recursive(allCards, copies, startIndex)...)
		// fmt.Println(strings.Repeat("/\\", int(infoLen/2)))
		// fmt.Println(strings.Repeat("-", infoLen))
        previousIndex = startIndex
	}

	return total
}

func solve(inputLister InputLister) (Solution, Solution, error) {
	input := inputLister.GetCards()

	partOneSolution := 0

	for _, card := range input {
		winning, err := getWinningNumbers(card)
		if err != nil {
			return partOneSolution, -1, err
		}

		winningCount := len(winning)

		cardPts := 1 << winningCount / 2
		// fmt.Printf("Total card pts for %d cards: %d\n", len(winning), cardPts)
		partOneSolution += cardPts
	}

	partTwoCards := recursive(input, input, -1)
	// fmt.Printf("Part two?\n%+v\n", strings.Join(partTwoCards, "\n"))
	partTwoSolution := len(partTwoCards)
	// fmt.Printf("Part two Solution: %d\n", partTwoSolution)

	return partOneSolution, partTwoSolution, nil
}
