package main

import (
	"fmt"
	"slices"
	"testing"
)

var mockData = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

var mockData2 = []string{
	"....#........",
	".........#...",
	"#............",
	".............",
	".............",
	"........#....",
	".#...........",
	"............#",
	".............",
	".............",
	".........#...",
	"#....#.......",
}

func TestParseData(t *testing.T) {
	g := parseData(mockData)

	if height := g.Height(); height != 10 {
		t.Errorf("Height: got %v, want %v", height, 10)
	}

	if width := g.Width(); width != 10 {
		t.Errorf("Width: got %v, want %v", width, 10)
	}
}

func TestGetGalaxies(t *testing.T) {
	g := parseData(mockData)
	galaxies := getGalaxies(g)

	if len(galaxies) != 9 {
		t.Errorf("Galaxies: got %v, want %v", len(galaxies), 9)
	}
}

func TestRowColsWithoutGalaxies(t *testing.T) {
	g := parseData(mockData)
	galaxies := getGalaxies(g)

	rows, cols := rowsColsWithoutGalaxy(g, galaxies)
	var wantRows = []int{3, 7}
	var wantCols = []int{2, 5, 8}

	if !slices.Equal(rows, wantRows) {
		t.Errorf("Rows: got %v, want %v", rows, wantRows)
	}

	if !slices.Equal(cols, wantCols) {
		t.Errorf("Cols: got %v, want %v", cols, wantCols)
	}
}

func TestAppendColumn(t *testing.T) {
	g := parseData(mockData)

	wantValues := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	col := 3
	//rows, cols := rowsColsWithoutGalaxy(g, galaxies)
	g = appendColumn(g, col, wantValues)

	if g.Width() != 11 {
		t.Errorf("Height: got %v, want %v", len(g), 11)
	}

	for y, row := range g {
		if row[col] != wantValues[y] {
			t.Errorf("got %v, want %v", row[col], wantValues[y])
		}
	}
}

func TestAppendRow(t *testing.T) {
	g := parseData(mockData)

	wantValues := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	row := 3
	//rows, cols := rowsColsWithoutGalaxy(g, galaxies)
	g = appendRow(g, row, wantValues)

	if g.Height() != 11 {
		t.Errorf("Height: got %v, want %v", len(g), 11)
	}

	for x, col := range g[row] {
		if col != wantValues[x] {
			t.Errorf("got %v, want %v", col, wantValues[x])
		}
	}
}

func TestAllPairs(t *testing.T) {
	galaxies := getGalaxies(parseData(mockData))

	pairs := allPairs(galaxies)

	if len(pairs) != 36 {
		t.Errorf("got %v, want %v", len(pairs), 36)
	}
}

func TestSumAllDistances(t *testing.T) {
	g := parseData(mockData2)

	sum := sumAllDistances(g)

	if sum != 374 {
		t.Errorf("got %v, want %v", sum, 374)
	}
}

func TestExpand(t *testing.T) {
	g := parseData(mockData)
	want := parseData(mockData2)

	g = expand(g)

	for idx, row := range g {
		if !slices.Equal(row, want[idx]) {
			t.Errorf("got %v, want %v", row, want[idx])
		}
	}
}

func TestExpandAndSumAllDistances(t *testing.T) {
	g := parseData(mockData)
	sum := expandAndSumAllDistances(g)

	if sum != 374 {
		t.Errorf("got %v, want %v", sum, 374)
	}
}

func TestDistance(t *testing.T) {
	g := parseData(mockData)
	rows, cols := rowsColsWithoutGalaxy(g, getGalaxies(g))
	p1 := point{X: 1, Y: 5}
	p2 := point{X: 4, Y: 9}

	want := 9
	got := distance(rows, cols, p1, p2, 2)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumDistancesVirtualExpand(t *testing.T) {
	g := parseData(mockData)

	testCases := []struct {
		n    int
		want int
	}{
		{2, 374},
		{10, 1030},
		{100, 8410},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.n), func(t *testing.T) {
			got := sumDistancesVirtualExpand(g, tc.n)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
