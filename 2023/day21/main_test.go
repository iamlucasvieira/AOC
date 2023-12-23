package main

import (
	"testing"
)

var mockData = []string{
	"...........",
	".....###.#.",
	".###.##..#.",
	"..#.#...#..",
	"....#.#....",
	".##..S####.",
	".##..#...#.",
	".......##..",
	".##.#.####.",
	".##..##.##.",
	"...........",
}

func TestParse(t *testing.T) {
	g, err := parse(mockData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if g.Width() != 11 {
		t.Errorf("expected width to be 11, got %d", g.Width())
	}

	if g.Height() != 11 {
		t.Errorf("expected height to be 11, got %d", g.Height())
	}

	if g.Get(point{X: 5, Y: 5}) != "S" {
		t.Errorf("expected S at 5,5, got %s", g.Get(point{5, 5}))
	}
}

func TestNeighbours(t *testing.T) {
	g, err := parse(mockData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	testCases := []struct {
		p    point
		want []point
	}{
		{
			p:    point{X: 5, Y: 5},
			want: []point{{4, 5}, {5, 4}},
		}, {
			p:    point{X: 4, Y: 2},
			want: []point{{X: 4, Y: 1}},
		},
	}

	for _, tc := range testCases {
		got := neighbours(tc.p, g)

		if len(got) != len(tc.want) {
			t.Errorf("expected %d neighbours, got %d", len(tc.want), len(got))
		}

		for i := range got {
			if got[i] != tc.want[i] {
				t.Errorf("expected neighbour %d to be %v, got %v", i, tc.want[i], got[i])
			}
		}
	}
}

func TestFindStart(t *testing.T) {
	g, err := parse(mockData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	start := findStart(g)
	want := point{X: 5, Y: 5}

	if !start.Equal(want) {
		t.Errorf("expected start to be %v, got %v", want, start)
	}
}

func TestPossibleEnd(t *testing.T) {
	g, err := parse(mockData)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	testCases := []struct {
		name string
		n    int
		want int
	}{
		{"Moving 1", 1, 2},
		{"Moving 2", 2, 4},
		{"Moving 3", 3, 6},
		{"Moving 4", 6, 16},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			start := findStart(g)
			got := possibleEnd(start, g, tc.n)

			if len(got) != tc.want {
				t.Errorf("expected %d possible end points, got %d", tc.want, len(got))
			}
		})
	}
}

func TestPart1(t *testing.T) {
	want := 3740
	got := part1()

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}
