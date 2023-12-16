package main

import "testing"

var mockData = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestRound(t *testing.T) {
	testCases := []struct {
		player   string
		opponent string
		want     int
	}{
		{"R", "R", 3},
		{"R", "S", 6},
		{"R", "P", 0},
		{"S", "R", 0},
		{"S", "S", 3},
		{"S", "P", 6},
		{"P", "R", 6},
		{"P", "S", 0},
		{"P", "P", 3},
	}

	for _, tc := range testCases {
		t.Run(tc.player+tc.opponent, func(t *testing.T) {
			got := round(tc.opponent, tc.player)
			if got != tc.want {
				t.Errorf("round(%s, %s) = %d; want %d", tc.opponent, tc.player, got, tc.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	got, err := parse(mockData)
	if err != nil {
		t.Errorf("parse(%v) = %v; want nil", mockData, err)
	}

	want := [][2]string{
		{"R", "P"},
		{"P", "R"},
		{"S", "S"},
	}

	for i, tc := range want {
		if got[i][0] != tc[0] || got[i][1] != tc[1] {
			t.Errorf("parse(%v) = %v; want %v", mockData, got, want)
		}
	}
}

func TestScore(t *testing.T) {
	rounds, err := parse(mockData)
	if err != nil {
		t.Errorf("parse(%v) = %v; want nil", mockData, err)
	}
	got := score(rounds)
	want := 15

	if got != want {
		t.Errorf("score(%v) = %d; want %d", mockData, got, want)
	}
}

func TestPlayerChoice(t *testing.T) {
	testCases := []struct {
		opponent string
		order    string // R (Win), P (Draw), S (Lose)
		want     string
	}{
		{"R", "R", "S"},
		{"R", "P", "R"},
		{"R", "S", "P"},
		{"P", "R", "R"},
		{"P", "P", "P"},
		{"P", "S", "S"},
		{"S", "R", "P"},
		{"S", "P", "S"},
		{"S", "S", "R"},
	}

	for _, tc := range testCases {
		t.Run(tc.opponent+tc.order, func(t *testing.T) {
			got := playerChoice(tc.opponent, tc.order)
			if got != tc.want {
				t.Errorf("playerChoice(%s, %s) = %s; want %s", tc.opponent, tc.order, got, tc.want)
			}
		})
	}
}

func TestScorePlayerChoice(t *testing.T) {
	rounds, err := parse(mockData)
	if err != nil {
		t.Errorf("parse(%v) = %v; want nil", mockData, err)
	}
	got := scorePlayerChoice(rounds)
	want := 12

	if got != want {
		t.Errorf("scorePlayerChoice(%v) = %d; want %d", mockData, got, want)
	}
}
