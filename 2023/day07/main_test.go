package main

import (
	"github.com/iamlucasvieira/aoc/utils"
	"slices"
	"sort"
	"strconv"
	"testing"
)

var mockData = []string{
	"32T3K 765\n",
	"T55J5 684\n",
	"KK677 28\n",
	"KTJJT 220",
	"QQQJA 483",
}

func TestGetCardValue(t *testing.T) {
	tests := []struct {
		label string
		want  int
	}{
		{"A", 14},
		{"K", 13},
		{"Q", 12},
		{"J", 11},
		{"T", 10},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
	}

	for _, tt := range tests {
		got, err := getCardValue(tt.label, labelToValue)
		if err != nil {
			t.Errorf("getCardValue(%q) = %d, want %d", tt.label, got, tt.want)
		}
		if got != tt.want {
			t.Errorf("getCardValue(%q) = %d, want %d", tt.label, got, tt.want)
		}
	}
}

func TestGetCardValueInvalid(t *testing.T) {
	tests := []struct {
		label string
	}{
		{"X"},
		{"1"},
		{"0"},
		{"10"},
	}

	for _, tt := range tests {
		got, err := getCardValue(tt.label, labelToValue)
		if err == nil {
			t.Errorf("getCardValue(%q) = %d, want error", tt.label, got)
		}
	}
}

func TestGetHandValue(t *testing.T) {
	testCases := []utils.TestCase[[]int, int]{
		{[]int{2, 3, 10, 11, 14}, 0},
		{[]int{3, 2, 10, 3, 14}, 1},
		{[]int{13, 13, 6, 7, 7}, 2},
		{[]int{10, 5, 5, 11, 5}, 3},
		{[]int{12, 12, 12, 14, 14}, 4},
		{[]int{13, 13, 13, 13, 14}, 5},
		{[]int{14, 14, 14, 14, 14}, 6},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getHandValue(tc.Input)
			if err != nil {
				t.Errorf("getHandValue(%v) = %d, want %d", tc.Input, got, tc.Expected)
			}
			if got != tc.Expected {
				t.Errorf("getHandValue(%v) = %d, want %d", tc.Input, got, tc.Expected)
			}
		})
	}
}

func TestMakeHand(t *testing.T) {
	testCases := []struct {
		labels      []string
		bid         int
		cardsWanted []int
		valueWanted int
	}{
		{[]string{"A", "K", "Q", "J", "T"}, 100, []int{14, 13, 12, 11, 10}, 0},
		{[]string{"3", "2", "T", "3", "K"}, 765, []int{3, 2, 10, 3, 13}, 1},
		{[]string{"K", "K", "6", "7", "7"}, 1000, []int{13, 13, 6, 7, 7}, 2},
		{[]string{"T", "5", "5", "J", "5"}, 10000, []int{10, 5, 5, 11, 5}, 3},
		{[]string{"Q", "Q", "Q", "K", "K"}, 100000, []int{12, 12, 12, 13, 13}, 4},
		{[]string{"K", "K", "K", "K", "A"}, 1000000, []int{13, 13, 13, 13, 14}, 5},
		{[]string{"A", "A", "A", "A", "A"}, 10000000, []int{14, 14, 14, 14, 14}, 6},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := makeHand(tc.labels, tc.bid)

			if err != nil {
				t.Errorf("makeHand(%v, %d) = %v, want %d", tc.labels, tc.bid, got, tc.bid)
			}
			if got.bid != tc.bid {
				t.Errorf("makeHand(%v, %d) = %v, want %d", tc.labels, tc.bid, got.bid, tc.bid)
			}
			for i, card := range got.cards {
				if card != tc.cardsWanted[i] {
					t.Errorf("makeHand(%v, %d) = %v, want %v", tc.labels, tc.bid, got.cards, tc.cardsWanted)
				}
			}
			if got.value != tc.valueWanted {
				t.Errorf("makeHand(%v, %d) = %v, want %v", tc.labels, tc.bid, got.value, tc.valueWanted)
			}
		})
	}
}

func TestParseData(t *testing.T) {
	data, err := parseData(mockData)

	if err != nil {
		t.Errorf("parseData(%v) = %v, want %v", mockData, err, nil)
	}

	if len(data) != 5 {
		t.Errorf("parseData(%v) = %v, want %v", mockData, data, 5)
	}

	if data[0].bid != 765 {
		t.Errorf("parseData(%v) = %v, want %v", mockData, data[0].bid, 765)
	}

	if data[0].cards[0] != 3 {
		t.Errorf("parseData(%v) = %v, want %v", mockData, data[0].cards[0], 3)
	}

}

func TestSortHands(t *testing.T) {
	var hs hands = []hand{
		{[]int{3, 2, 10, 3, 13}, 1, 1},
		{[]int{10, 5, 5, 11, 5}, 2, 3},
		{[]int{13, 13, 6, 7, 7}, 3, 2},
		{[]int{13, 10, 11, 11, 10}, 4, 2},
		{[]int{12, 12, 12, 11, 14}, 5, 3},
	}

	tc := []utils.TestCase[hand, int]{
		{hs[0], 0},
		{hs[1], 3},
		{hs[2], 2},
		{hs[3], 1},
		{hs[4], 4},
	}

	sort.Sort(hs)

	for i, c := range tc {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			currentHand := c.Input
			expectedIndex := c.Expected
			handAtExpectedIndex := hs[expectedIndex]

			if !slices.Equal(currentHand.cards, handAtExpectedIndex.cards) {
				t.Errorf("Want index %d to be %v, got %v", expectedIndex, handAtExpectedIndex.cards, currentHand.cards)
			}

		})
	}

}

func TestHands_Score(t *testing.T) {
	data, err := parseData(mockData)

	if err != nil {
		t.Errorf("parseData(%v) = %v, want %v", mockData, err, nil)
	}

	score := data.score()

	if score != 6440 {
		t.Errorf("score() = %v, want %v", score, 6440)
	}
}

func TestGetHandValueWildJ(t *testing.T) {
	tc := []utils.TestCase[[]int, int]{
		{[]int{2, 3, 4, 5, 1}, 1}, // 4 different + 1J -> Score should become 1
		{[]int{2, 3, 4, 1, 1}, 3}, // 3 different + 2J -> Score should become 3
		{[]int{2, 3, 1, 1, 1}, 5}, // 2 different + 3J -> Score should become 5
		{[]int{2, 1, 1, 1, 1}, 6}, // 1 different + 4J -> Score should become 6

	}

	for i, c := range tc {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getHandValueWildJ(c.Input)
			if err != nil {
				t.Errorf("getHandValueWildJ(%v) = %d, want %d", c.Input, got, c.Expected)
			}
			if got != c.Expected {
				t.Errorf("getHandValueWildJ(%v) = %d, want %d", c.Input, got, c.Expected)
			}
		})
	}
}

func TestHandsWildJ_Score(t *testing.T) {
	data, err := parseDataWildJ(mockData)

	if err != nil {
		t.Errorf("parseData(%v) = %v, want %v", mockData, err, nil)
	}

	score := data.score()

	if score != 5905 {
		t.Errorf("score() = %v, want %v", score, 5905)
	}
}
