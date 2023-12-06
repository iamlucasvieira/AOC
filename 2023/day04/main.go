package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// parseCard parses the input string card. Returns two slices.
// One with the winning numbers and another with the elf numbers
func parseCard(input string) ([]int, []int, error) {

	cardParts := strings.SplitN(input, ":", 2)

	if len(cardParts) != 2 {
		return nil, nil, fmt.Errorf("invalid card format: %s", input)
	}

	numberParts := strings.SplitN(cardParts[1], "|", 2)

	if len(numberParts) != 2 {
		return nil, nil, fmt.Errorf("invalid card format for numbers: %s", cardParts[1])
	}

	extractNumbers := func(n string) ([]int, error) {
		nRegex := regexp.MustCompile(`\d+`)
		var numbers []int

		for _, match := range nRegex.FindAllString(n, -1) {
			number, err := strconv.Atoi(match)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, number)
		}
		return numbers, nil
	}

	winningNumbers, err := extractNumbers(numberParts[0])
	if err != nil {
		return nil, nil, err
	}

	elfNumbers, err := extractNumbers(numberParts[1])
	if err != nil {
		return nil, nil, err
	}

	return winningNumbers, elfNumbers, nil
}

// scoreFunc is a function that returns the score of a card
type scoreFunc func(int) int

// scoreCardTemplate is a template for scoring a card
func scoreCardTemplate(winningNumbers, elfNumbers []int, f scoreFunc) int {

	var count int

	for _, number := range elfNumbers {
		if slices.Contains(winningNumbers, number) {
			count++
		}
	}

	score := f(count)
	return score
}

// scoreCard returns the score of a card, which is the sum of 2^(n-1) where n is the number of winning numbers
func scoreCard(winningNumbers, elfNumbers []int) int {
	sFunc := func(n int) int {
		return int(math.Pow(2, float64(n)-1))
	}
	return scoreCardTemplate(winningNumbers, elfNumbers, sFunc)
}

// scoreCardCount returns the score of a card, which is the number of winning numbers
func scoreCardCount(winningNumbers, elfNumbers []int) int {
	sFunc := func(n int) int {
		return n
	}
	return scoreCardTemplate(winningNumbers, elfNumbers, sFunc)
}

// scoreMultipleCards returns the score of multiple cards combined
func scoreMultipleCards(cards []string) int {
	var score int

	for _, card := range cards {
		if card == "" {
			continue
		}
		winningNumbers, elfNumbers, err := parseCard(card)
		if err != nil {
			panic(err)
		}
		score += scoreCard(winningNumbers, elfNumbers)
	}

	return score
}

func scoreCumulative(cardIdx int, cardsScore, cumulativeScores map[int]int) int {
	// Depends on n cards
	nDepends := cardsScore[cardIdx]

	// Cumulative score of the card
	score, ok := cumulativeScores[cardIdx]

	if !ok {
		score = 1

		if nDepends == 0 {
			cumulativeScores[cardIdx] = 1
			return 1
		}

		for _, idx := range idxCardsWon(cardIdx, nDepends) {
			score += scoreCumulative(idx, cardsScore, cumulativeScores)
		}

		cumulativeScores[cardIdx] = score
	}

	return score
}

func idxCardsWon(idx int, score int) []int {
	var idxCards = make([]int, score)

	for i := 0; i < score; i++ {
		idxCards[i] = idx + i + 1
	}

	return idxCards
}

func cardsWon(cards []string) int {
	cardsScore := make(map[int]int)

	for idx, card := range cards {
		if card == "" {
			continue
		}
		winningNumbers, elfNumbers, err := parseCard(card)
		if err != nil {
			panic(err)
		}
		score := scoreCardCount(winningNumbers, elfNumbers)
		cardsScore[idx] = score
	}

	cumulativeScores := make(map[int]int)

	var score int
	for idx := range cardsScore {
		score += scoreCumulative(idx, cardsScore, cumulativeScores)
	}
	return score
}

func part1() {
	fmt.Println("Part 1:")
	data := utils.ReadFile("input.txt")
	score := scoreMultipleCards(data)
	fmt.Println("Score:", score)
}

func part2() {
	fmt.Println("Part 2:")
	data := utils.ReadFile("input2.txt")
	score := cardsWon(data)
	fmt.Println("Score:", score)
}

func main() {
	part1()
	part2()
}
