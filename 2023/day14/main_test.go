package main

import (
	"slices"
	"testing"
)

var mockData = []string{
	"O....#....",
	"O.OO#....#",
	".....##...",
	"OO.#O....O",
	".O.....O#.",
	"O.#..O.#.#",
	"..O..#O..O",
	".......O..",
	"#....###..",
	"#OO..#....",
}

var mockDataUp = []string{
	"OOOO.#.O..",
	"OO..#....#",
	"OO..O##..O",
	"O..#.OO...",
	"........#.",
	"..#....#.#",
	"..O..#.O.O",
	"..O.......",
	"#....###..",
	"#....#....",
}

var mockDataCycle1 = []string{
	".....#....",
	"....#...O#",
	"...OO##...",
	".OO#......",
	".....OOO#.",
	".O#...O#.#",
	"....O#....",
	"......OOOO",
	"#...O###..",
	"#..OO#....",
}

var mockDataCycle2 = []string{
	".....#....",
	"....#...O#",
	".....##...",
	"..O#......",
	".....OOO#.",
	".O#...O#.#",
	"....O#...O",
	".......OOO",
	"#..OO###..",
	"#.OOO#...O",
}

var mockDataCycle3 = []string{
	".....#....",
	"....#...O#",
	".....##...",
	"..O#......",
	".....OOO#.",
	".O#...O#.#",
	"....O#...O",
	".......OOO",
	"#...O###.O",
	"#.OOO#...O",
}

func TestParse(t *testing.T) {
	data := parse(mockData)

	if len(data) != 10 {
		t.Errorf("Expected 10 columns, got %d", len(data))
	}
}

func TestRocksIndex(t *testing.T) {
	data := parse(mockData)

	testCases := []struct {
		index    int
		expected []int
	}{
		{0, []int{0, 1, 3, 5}},
		{1, []int{3, 4, 9}},
		{2, []int{1, 6, 9}},
		{3, []int{1}},
		{4, []int{3}},
		{5, []int{5}},
		{6, []int{6}},
		{7, []int{4, 7}},
		{8, []int{}},
		{9, []int{3, 6}},
	}

	for _, tc := range testCases {
		result := rocksIndex(data[tc.index])

		if !slices.Equal(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func TestRocksImpactColumn(t *testing.T) {
	data := parse(mockData)

	testCases := []struct {
		index    int
		expected int
	}{
		{0, 34},
		{1, 27},
		{2, 17},
		{3, 10},
		{4, 8},
		{5, 7},
		{6, 7},
		{7, 14},
		{8, 0},
		{9, 12},
	}

	for _, tc := range testCases {
		t.Run(data[tc.index], func(t *testing.T) {
			result := rocksImpactColumn(data[tc.index])

			if result != tc.expected {
				t.Errorf("idx: %v, Expected %d, got %d", tc.index, tc.expected, result)
			}
		})
	}
}

func TestAllImpact(t *testing.T) {
	data := parse(mockData)

	result := allImpact(data)
	expected := 136

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestParseTable(t *testing.T) {
	data := parseAsTable(mockData)

	if len(data) != 10 {
		t.Errorf("Expected 10 columns, got %d", len(data))
	}

	if len(data[0]) != 10 {
		t.Errorf("Expected 10 rows, got %d", len(data[0]))
	}
}

func TestTable_Column(t *testing.T) {
	data := parseAsTable(mockData)

	testCases := []struct {
		index    int
		expected string
	}{
		{0, "OO.O.O..##"},
		{1, "...OO....O"},
	}

	for _, tc := range testCases {
		t.Run(data.Column(tc.index), func(t *testing.T) {
			result := data.Column(tc.index)

			if result != tc.expected {
				t.Errorf("idx: %v, Expected %s, got %s", tc.index, tc.expected, result)
			}
		})
	}
}

func TestTable_Row(t *testing.T) {
	data := parseAsTable(mockData)

	testCases := []struct {
		index    int
		expected string
	}{
		{0, "O....#...."},
		{1, "O.OO#....#"},
		{2, ".....##..."},
		{3, "OO.#O....O"},
		{4, ".O.....O#."},
		{5, "O.#..O.#.#"},
		{6, "..O..#O..O"},
		{7, ".......O.."},
		{8, "#....###.."},
		{9, "#OO..#...."},
	}

	for _, tc := range testCases {
		t.Run(data.Row(tc.index), func(t *testing.T) {
			result := data.Row(tc.index)

			if result != tc.expected {
				t.Errorf("idx: %v, Expected %s, got %s", tc.index, tc.expected, result)
			}
		})
	}

}

func TestTable_ReplaceRow(t *testing.T) {
	data := parseAsTable(mockData)

	data.replaceRow(0, "OOOOOOOOOO")

	if data.Row(0) != "OOOOOOOOOO" {
		t.Errorf("Expected OOOOOOOOOO, got %s", data.Row(0))
	}
}

func TestTable_ReplaceColumn(t *testing.T) {
	data := parseAsTable(mockData)

	data.replaceColumn(0, "OOOOOOOOOO")

	if data.Column(0) != "OOOOOOOOOO" {
		t.Errorf("Expected OOOOOOOOOO, got %s", data.Column(0))
	}
}

func TestMoveLeft(t *testing.T) {

	testCases := []struct {
		data string
		want string
	}{
		{"..O..#O..O", "O....#OO.."},
		{"........O", "O........"},
		{"O.O......", "OO......."},
		{"O.O.O.O.O", "OOOOO...."},
	}

	for _, tc := range testCases {
		t.Run(tc.data, func(t *testing.T) {
			result := moveLeft(tc.data)

			if result != tc.want {
				t.Errorf("Expected %s, got %s", tc.want, result)
			}
		})
	}
}

func TestMoveRight(t *testing.T) {

	testCases := []struct {
		data string
		want string
	}{
		{"..O..#O..O", "....O#..OO"},
		{"........O", "........O"},
		{"O.O......", ".......OO"},
		{"O.O.O.O.O", "....OOOOO"},
	}

	for _, tc := range testCases {
		t.Run(tc.data, func(t *testing.T) {
			result := moveRight(tc.data)

			if result != tc.want {
				t.Errorf("Expected %s, got %s", tc.want, result)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	testCases := []struct {
		data string
		want string
	}{
		{"..O..#O..O", "O..O#..O.."},
		{"........O", "O........"},
		{"O.O......", "......O.O"},
		{"O.O.O.O.O", "O.O.O.O.O"},
	}

	for _, tc := range testCases {
		t.Run(tc.data, func(t *testing.T) {
			result := reverseString(tc.data)

			if result != tc.want {
				t.Errorf("Expected %s, got %s", tc.want, result)
			}
		})
	}
}

func TestMoveTableRight(t *testing.T) {
	data := [][]string{
		{"O", ".", "."},
		{"O", ".", "."},
		{"O", ".", "."},
		{"O", ".", "."},
	}

	want := [][]string{
		{".", ".", "O"},
		{".", ".", "O"},
		{".", ".", "O"},
		{".", ".", "O"},
	}
	c := make(cache)
	result := moveTableRight(data, c)

	for i := range result {
		if !slices.Equal(result[i], want[i]) {
			t.Errorf("Expected %v, got %v", want[i], result[i])
		}
	}
}

func TestMoveTableLeft(t *testing.T) {
	data := [][]string{
		{".", ".", "O"},
		{".", ".", "O"},
		{".", ".", "O"},
		{".", ".", "O"},
	}

	want := [][]string{
		{"O", ".", "."},
		{"O", ".", "."},
		{"O", ".", "."},
		{"O", ".", "."},
	}
	c := make(cache)
	result := moveTableLeft(data, c)

	for i := range result {
		if !slices.Equal(result[i], want[i]) {
			t.Errorf("Expected %v, got %v", want[i], result[i])
		}
	}
}

func TestMoveTableDown(t *testing.T) {
	data := [][]string{
		{"O", "O", "O", "O"},
		{".", "#", ".", "."},
		{".", ".", ".", "."},
	}

	want := [][]string{
		{".", "O", ".", "."},
		{".", "#", ".", "."},
		{"O", ".", "O", "O"},
	}
	c := make(cache)
	result := moveTableDown(data, c)

	for i := range result {
		if !slices.Equal(result[i], want[i]) {
			t.Errorf("Expected %v, got %v", want[i], result[i])
		}
	}
}

func TestMoveTableUp(t *testing.T) {
	data := parseAsTable(mockData)
	want := parseAsTable(mockDataUp)

	c := make(cache)
	result := moveTableUp(data, c)

	for i := range result {
		if !slices.Equal(result[i], want[i]) {
			t.Errorf("Expected %v, got %v", want[i], result[i])
		}
	}
}

func TestMoveAndScore(t *testing.T) {
	data := parseAsTable(mockData)

	c := make(cache)
	result := moveTableUp(data, c)
	score := result.score()

	if score != 136 {
		t.Errorf("Expected 136, got %v", score)
	}
}

func TestCycle(t *testing.T) {
	data := [][]string{
		{"#", ".", "."},
		{".", ".", "."},
		{"O", ".", "O"},
	}

	want := [][]string{
		{"#", ".", "."},
		{".", ".", "."},
		{".", "O", "O"},
	}

	c := make(cache)
	result := cycle(data, c)

	for i := range result {
		if !slices.Equal(result[i], want[i]) {
			t.Errorf("Expected %v, got %v", want[i], result[i])
		}
	}
}

func TestCycleNTimes(t *testing.T) {
	testCases := []struct {
		n    int
		data [][]string
		want [][]string
	}{
		{1, parseAsTable(mockData), parseAsTable(mockDataCycle1)},
		{2, parseAsTable(mockData), parseAsTable(mockDataCycle2)},
		{3, parseAsTable(mockData), parseAsTable(mockDataCycle3)},
	}

	for _, tc := range testCases {
		c := make(cache)
		result := cycleNTimes(tc.data, c, tc.n)

		for i := range result {
			if !slices.Equal(result[i], tc.want[i]) {
				t.Errorf("idx: %v, Expected %v, got %v", i, tc.want[i], result[i])
			}
		}
	}
}

func TestCycleAndScore(t *testing.T) {
	data := parseAsTable(mockData)

	c := make(cache)
	result := cycleNTimes(data, c, 1000)
	score := result.score()

	if score != 64 {
		t.Errorf("Expected 64, got %v", score)
	}
}
