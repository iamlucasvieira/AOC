package main

import (
	"fmt"
	"testing"
)

var mockData = []string{
	".|...\\....",
	"|.-.\\.....",
	".....|-...",
	"........|.",
	"..........",
	".........\\",
	"..../.\\\\..",
	".-.-/..|..",
	".|....-|.\\",
	"..//.|....",
}

func TestParse(t *testing.T) {

	actual := parse(mockData)

	if len(actual) != 10 {
		t.Errorf("Expected 10, got %d", len(actual))
	}

	if len(actual[0]) != 10 {
		t.Errorf("Expected 10, got %d", len(actual[0]))
	}

	for item := range actual.Iterator() {
		if item.Value.energised {
			t.Errorf("Expected false, got %v", item.Value.energised)
		}
	}
}

func TestMoveBeamAllEnergised(t *testing.T) {

	var g1 = []string{
		"..........",
	}
	var g2 = []string{
		".",
		".",
		".",
		".",
		".",
		".",
		".",
	}

	testCases := []struct {
		g         grid
		direction point
	}{
		{parse(g1), right},
		{parse(g2), down},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.g), func(t *testing.T) {
			moveBeam(point{}, tc.direction, tc.g, make(cache))
			for item := range tc.g.Iterator() {
				if !item.Value.energised {
					t.Errorf("Expected true, got %v", item.Value.energised)
				}
			}
		})
	}

}

func TestNEnergised(t *testing.T) {
	var g1 = []string{
		"..........",
	}
	var g2 = []string{
		".",
		".",
		".",
		".",
		".",
		".",
		".",
	}

	testCases := []struct {
		g         grid
		direction point
		want      int
	}{
		{parse(g1), right, 10},
		{parse(g2), down, 7},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.g), func(t *testing.T) {
			moveBeam(point{}, tc.direction, tc.g, make(cache))
			got := nEnergised(tc.g)
			if got != tc.want {
				t.Errorf("Expected %d, got %d", tc.want, got)
			}
		})
	}
}

func TestMoveBeamSplitVertical(t *testing.T) {
	var g1 = []string{
		"..........",
		".........|",
		"..........",
	}

	g := parse(g1)
	wantEnergised := []point{
		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 4, Y: 1},
		{X: 5, Y: 1},
		{X: 6, Y: 1},
		{X: 7, Y: 1},
		{X: 8, Y: 1},
		{X: 9, Y: 1},
		{X: 9, Y: 0},
		{X: 9, Y: 2},
	}

	moveBeam(point{X: 0, Y: 1}, right, g, make(cache))

	// Check if only the expected points are energised
	for _, p := range wantEnergised {
		if !g.Get(p).energised {
			t.Errorf("Expected true, got %v", g.Get(p).energised)
		}

		if nEnergised(g) != len(wantEnergised) {
			t.Errorf("Expected %d, got %d", len(wantEnergised), nEnergised(g))
		}
	}
}

func TestMoveBeamSplitHorizontal(t *testing.T) {
	var g1 = []string{
		".....|....",
		".....-....",
		"..........",
	}

	g := parse(g1)
	wantEnergised := []point{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 3, Y: 0},
		{X: 4, Y: 0},
		{X: 5, Y: 0},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 4, Y: 1},
		{X: 5, Y: 1},
		{X: 6, Y: 1},
		{X: 7, Y: 1},
		{X: 8, Y: 1},
		{X: 9, Y: 1},
	}

	moveBeam(point{X: 0, Y: 0}, right, g, make(cache))

	for _, p := range wantEnergised {
		if !g.Get(p).energised {
			t.Errorf("Expected true, got %v", g.Get(p).energised)
		}

		if nEnergised(g) != len(wantEnergised) {
			t.Errorf("Expected %d, got %d", len(wantEnergised), nEnergised(g))
		}
	}
}

func TestMoveBeamSplitDiagonal(t *testing.T) {
	var g1 = []string{
		".........\\",
		"/......../",
		"\\.........",
	}

	g := parse(g1)
	want := len(g1) * len(g1[0])

	moveBeam(point{X: 0, Y: 0}, right, g, make(cache))

	if nEnergised(g) != want {
		t.Errorf("Expected %d, got %d", want, nEnergised(g))
	}
}

func TestMoveBeam(t *testing.T) {
	g := parse(mockData)

	moveBeam(point{X: 0, Y: 0}, right, g, make(cache))
	want := 46
	fmt.Println(energyString(g))
	if nEnergised(g) != want {
		t.Errorf("Expected %d, got %d", want, nEnergised(g))
	}

}

func TestMaxEnergy(t *testing.T) {
	value := maxEnergy(mockData)
	want := 51
	if value != want {
		t.Errorf("Expected %d, got %d", want, value)
	}
}
