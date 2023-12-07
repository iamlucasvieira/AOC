package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"sort"
	"strconv"
	"strings"
)

var labelToValue = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
}

var labelToValueWildJ = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1,
	"T": 10,
}

type labelMap map[string]int

// getCardValue returns the value of a card.
func getCardValue(label string, lm labelMap) (int, error) {
	if value, ok := lm[label]; ok {
		return value, nil
	}

	// If the label is not in the map, it must be a number.
	value, err := strconv.Atoi(label)
	if err != nil {
		return 0, err
	}

	if value < 2 || value > 9 {
		return 0, fmt.Errorf("invalid card value: %d", value)
	}

	return value, nil
}

type handValueFunc func([]int) (int, error)

func getHandValueWildJ(cards []int) (int, error) {
	// Get card list without wild cards J:
	var cardsNoJ []int
	var cardsFrequency = make(map[int]int)
	for _, card := range cards {
		if card != 1 {
			cardsNoJ = append(cardsNoJ, card)
			cardsFrequency[card]++
		}
	}

	// If Hand has no wild cards J, return the conventional hand value:
	if len(cardsNoJ) == len(cards) {
		return getHandValue(cards)
	}

	// Get card that appears more often in the hand:
	var mostFrequent int
	var frequency int

	for c, f := range cardsFrequency {
		if f > frequency {
			mostFrequent = c
			frequency = f
		} else if f == frequency {
			if c > mostFrequent {
				mostFrequent = c
				frequency = f
			}
		}
	}

	//
	var cardsWithReplacedJ = make([]int, len(cards))
	for i, card := range cards {
		if card == 1 {
			cardsWithReplacedJ[i] = mostFrequent
		} else {
			cardsWithReplacedJ[i] = card
		}

	}

	return getHandValue(cardsWithReplacedJ)
}

// getHandValue returns the value of a hand.
// The values are:
// 0: All cards are different.
// 1: One pair.
// 2: Two pairs.
// 3: Three of a kind.
// 4: Full house: Three of a kind and a pair.
// 5: Four of a kind.
// 6: Five of a kind.
func getHandValue(cards []int) (int, error) {
	valuesToCount := make(map[int]int)

	for _, card := range cards {
		valuesToCount[card]++
	}

	switch len(valuesToCount) {
	// 0: All cards are different.
	case 5:
		return 0, nil
	// 1: One pair.
	case 4:
		return 1, nil
	// 2 or 3: Two pairs or three of a kind.
	case 3:
		for _, count := range valuesToCount {
			if count == 3 {
				return 3, nil
			}
		}
		return 2, nil
	// 4 or 5: Full house or four of a kind.
	case 2:
		for _, count := range valuesToCount {
			if count == 4 {
				return 5, nil
			}
		}
		return 4, nil
	// 6: Five of a kind.
	case 1:
		return 6, nil
	}

	return 0, fmt.Errorf("invalid number of cards: %d", len(cards))
}

// hand represents a hand of cards.
type hand struct {
	cards []int
	bid   int
	value int
}

type hands []hand

// Len returns the number of hands.
func (h hands) Len() int {
	return len(h)
}

// Swap swaps two hands.
func (h hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hands) score() int {
	var score int
	sort.Sort(h)

	for idx, currentHand := range h {
		score += currentHand.bid * (idx + 1)
	}

	return score
}

// Less returns true if the first hand is less than the second.
func (h hands) Less(i, j int) bool {
	if h[i].value == h[j].value {
		// Check cards values.
		for k := 0; k < 5; k++ {
			if h[i].cards[k] == h[j].cards[k] {
				continue
			}
			return h[i].cards[k] < h[j].cards[k]
		}
	}
	return h[i].value < h[j].value
}

type handMaker func([]string, int) (hand, error)

// makeHandTemplate returns a hand from a slice of card labels with a given map.
func makeHandTemplate(labels []string, bid int, lm labelMap, vf handValueFunc) (hand, error) {

	if len(labels) != 5 {
		return hand{}, fmt.Errorf("invalid number of cards: %d", len(labels))
	}

	cards := make([]int, len(labels))
	for i, label := range labels {
		cardValue, err := getCardValue(label, lm)
		if err != nil {
			return hand{}, err
		}
		cards[i] = cardValue
	}

	value, err := vf(cards)
	if err != nil {
		return hand{}, err
	}

	return hand{
		cards: cards,
		bid:   bid,
		value: value,
	}, nil
}

// makeHand returns a hand from a slice of card labels.
func makeHand(cards []string, bid int) (hand, error) {
	return makeHandTemplate(cards, bid, labelToValue, getHandValue)
}

// makeHandWildJ returns a hand from a slice of card labels.
func makeHandWildJ(cards []string, bid int) (hand, error) {
	return makeHandTemplate(cards, bid, labelToValueWildJ, getHandValueWildJ)
}

// parseDataTemplate returns a slice of hands from a slice of strings.
func parseDataTemplate(lines []string, hm handMaker) (hands, error) {
	var hands []hand
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		// Strip the trailing newline.
		line = strings.TrimSuffix(line, "\n")

		// Split the line into the cards and the bid.
		parts := strings.Split(line, " ")

		var cards = make([]string, len(parts[0]))

		// Create a slice of the cards.
		for i, card := range parts[0] {
			cards[i] = string(card)
		}

		// Convert the bid to an int.
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		var h hand
		h, err = hm(cards, bid)
		if err != nil {
			return nil, err
		}

		hands = append(hands, h)
	}
	return hands, nil
}

// parseData returns a slice of hands from a slice of strings.
func parseData(lines []string) (hands, error) {
	return parseDataTemplate(lines, makeHand)
}

// parseDataWildJ returns a slice of hands from a slice of strings.
func parseDataWildJ(lines []string) (hands, error) {
	return parseDataTemplate(lines, makeHandWildJ)
}

func part1() {
	fmt.Println("Part 1:")
	data, err := parseData(utils.ReadFile("input.txt"))
	if err != nil {
		panic(err)
	}
	score := data.score()
	fmt.Printf("The score is %d\n", score)
}

func part2() {
	fmt.Println("Part 2:")
	data, err := parseDataWildJ(utils.ReadFile("input2.txt"))
	if err != nil {
		panic(err)
	}
	score := data.score()
	fmt.Printf("The score is %d\n", score)
}
func main() {
	part1()
	part2()
}
