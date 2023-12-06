package main

import (
	"github.com/iamlucasvieira/aoc/utils"
	"slices"
	"sort"
	"testing"
)

func TestParseCards(t *testing.T) {

	testCases := []struct {
		input                  string
		expectedWinningNumbers []int
		expectedElfNumbers     []int
	}{
		{
			input:                  "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			expectedWinningNumbers: []int{41, 48, 83, 86, 17},
			expectedElfNumbers:     []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
		{
			input:                  "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			expectedWinningNumbers: []int{13, 32, 20, 16, 61},
			expectedElfNumbers:     []int{61, 30, 68, 82, 17, 32, 24, 19},
		},
		{
			input:                  "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			expectedWinningNumbers: []int{1, 21, 53, 59, 44},
			expectedElfNumbers:     []int{69, 82, 63, 72, 16, 21, 14, 1},
		},
		{
			input:                  "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			expectedWinningNumbers: []int{41, 92, 73, 84, 69},
			expectedElfNumbers:     []int{59, 84, 76, 51, 58, 5, 54, 83},
		},
		{
			input:                  "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			expectedWinningNumbers: []int{87, 83, 26, 28, 32},
			expectedElfNumbers:     []int{88, 30, 70, 12, 93, 22, 82, 36},
		},
		{
			input:                  "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			expectedWinningNumbers: []int{31, 18, 13, 56, 72},
			expectedElfNumbers:     []int{74, 77, 10, 23, 35, 67, 36, 11},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			winningNumbers, elfNumbers, err := parseCard(tc.input)

			// Sort the slices to make the comparison easier
			sort.Ints(winningNumbers)
			sort.Ints(elfNumbers)
			sort.Ints(tc.expectedWinningNumbers)
			sort.Ints(tc.expectedElfNumbers)

			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if !slices.Equal(winningNumbers, tc.expectedWinningNumbers) {
				t.Errorf("Expected winning numbers %v, got %v", tc.expectedWinningNumbers, winningNumbers)
			}

			if !slices.Equal(elfNumbers, tc.expectedElfNumbers) {
				t.Errorf("Expected elf numbers %v, got %v", tc.expectedElfNumbers, elfNumbers)
			}
		})
	}
}

func TestScoreCard(t *testing.T) {
	testCases := []utils.TestCase[string, int]{
		{Input: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", Expected: 8},
		{Input: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", Expected: 2},
		{Input: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", Expected: 2},
		{Input: "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", Expected: 1},
		{Input: "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", Expected: 0},
		{Input: "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", Expected: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			winningNumbers, elfNumbers, err := parseCard(tc.Input)
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			score := scoreCard(winningNumbers, elfNumbers)

			if score != tc.Expected {
				t.Errorf("Expected score %d, got %d", tc.Expected, score)
			}
		})
	}
}

func TestScoreCardCount(t *testing.T) {
	testCases := []utils.TestCase[string, int]{
		{Input: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", Expected: 8},
		{Input: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", Expected: 2},
		{Input: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", Expected: 2},
		{Input: "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", Expected: 1},
		{Input: "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", Expected: 0},
		{Input: "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", Expected: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			winningNumbers, elfNumbers, err := parseCard(tc.Input)
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			score := scoreCard(winningNumbers, elfNumbers)

			if score != tc.Expected {
				t.Errorf("Expected score %d, got %d", tc.Expected, score)
			}
		})
	}
}
func TestScoreMultipleCards(t *testing.T) {
	cards := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
	expectedScore := 13

	score := scoreMultipleCards(cards)

	if score != expectedScore {
		t.Errorf("Expected score %d, got %d", expectedScore, score)
	}
}

func TestCardsWon(t *testing.T) {
	cards := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	expectedScore := 30

	score := cardsWon(cards)

	if score != expectedScore {
		t.Errorf("Expected score %d, got %d", expectedScore, score)
	}

}
