package main

import (
	"github.com/iamlucasvieira/aoc/utils"
	"strconv"
	"testing"
)

func TestNewGame(t *testing.T) {
	game := "Game 1: 3 blue, 7 green, 10 red; 4 green, 4 red; 1 green, 7 blue, 5 red; 8 blue, 10 red; 7 blue, 19 red, 1 green"

	parsedGame, err := newGame(game)

	if err != nil {
		t.Errorf("newGame(%q) = %v; want %v", game, err, nil)
	}

	if parsedGame.id != 1 {
		t.Errorf("newGame(%q) = %v; want %v", game, parsedGame.id, 1)
	}

	if len(parsedGame.rounds) != 5 {
		t.Errorf("newGame(%q) = %v; want %v", game, len(parsedGame.rounds), 5)
	}

	// Test Colors in first round
	if parsedGame.rounds[0].blue != 3 {
		t.Errorf("newGame(%q) = %v; want %v", game, parsedGame.rounds[0].blue, 3)
	}

	if parsedGame.rounds[0].green != 7 {
		t.Errorf("newGame(%q) = %v; want %v", game, parsedGame.rounds[0].green, 7)
	}

	if parsedGame.rounds[0].red != 10 {
		t.Errorf("newGame(%q) = %v; want %v", game, parsedGame.rounds[0].red, 10)
	}
}

func TestIsGameValid(t *testing.T) {
	testCases := []utils.TestCase[string, bool]{
		{
			Input:    "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Expected: true,
		},
		{
			Input:    "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Expected: true,
		},
		{
			Input:    "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Expected: false,
		},
		{
			Input:    "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Expected: false,
		},
		{
			Input:    "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Expected: true,
		},
	}

	for _, tc := range testCases {
		game, err := newGame(tc.Input)
		if err != nil {
			t.Errorf("newGame(%d) = %v; want %v", game.id, err, nil)
		}
		t.Run(strconv.Itoa(game.id), func(t *testing.T) {

			result := isGameValid(game, 12, 13, 14)
			if result != tc.Expected {
				t.Errorf("isGameValid(%d) = %v; want %v", game.id, result, tc.Expected)
			}
		})
	}
}

func TestFewestCubes(t *testing.T) {

	testCases := []utils.TestCase[string, []int]{
		{
			Input:    "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Expected: []int{4, 2, 6},
		},
		{
			Input:    "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Expected: []int{1, 3, 4},
		},
		{
			Input:    "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Expected: []int{20, 13, 6},
		},
		{
			Input:    "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Expected: []int{14, 3, 15},
		},
		{
			Input:    "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Expected: []int{6, 3, 2},
		},
	}

	for _, tc := range testCases {
		game, err := newGame(tc.Input)
		if err != nil {
			t.Errorf("newGame(%d) = %v; want %v", game.id, err, nil)
		}
		t.Run(strconv.Itoa(game.id), func(t *testing.T) {
			red, green, blue := fewestCubes(game)
			if red != tc.Expected[0] {
				t.Errorf("RED: fewestCubes(%d) = %v; want %v", game.id, red, tc.Expected[0])
			}
			if green != tc.Expected[1] {
				t.Errorf("GREEN: fewestCubes(%d) = %v; want %v", game.id, green, tc.Expected[1])
			}
			if blue != tc.Expected[2] {
				t.Errorf("BLUE: fewestCubes(%d) = %v; want %v", game.id, blue, tc.Expected[2])
			}
		})
	}
}
