package main

import (
	"fmt"
	"slices"
	"testing"
)

var mockData = []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
}

func TestParse(t *testing.T) {
	data, err := parse(mockData)

	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	if len(data) != 2 {
		t.Errorf("Expected 2 patterns, got %v", len(data))
	}

	if len(data[0]) != 7 {
		t.Errorf("Expected 7 rows in pattern 1, got %v", len(data[0]))
	}

	if len(data[0][0]) != 9 {
		t.Errorf("Expected 9 columns in pattern 1, got %v", len(data[0][0]))
	}
}

func TestIsReflective(t *testing.T) {
	row := []string{".", ".", "#", "#", ".", ".", "#", "#", "."}

	testCases := []struct {
		x        int
		expected bool
	}{
		{0, false},
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.x), func(t *testing.T) {
			actual := isReflective(row, tc.x)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}

}

func TestAsRow(t *testing.T) {
	matrix := [][]string{
		{"A", "B", "C", "1"},
		{"D", "E", "F", "2"},
		{"G", "H", "I", "3"},
	}

	testCases := []struct {
		idx  int
		want []string
	}{
		{0, []string{"A", "D", "G"}},
		{1, []string{"B", "E", "H"}},
		{2, []string{"C", "F", "I"}},
		{3, []string{"1", "2", "3"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.idx), func(t *testing.T) {
			got := getColumn(matrix, tc.idx)
			if len(got) != len(tc.want) {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}

			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("Expected %v, got %v", tc.want, got)
				}
			}
		})
	}
}
func TestIsReflectiveColumn(t *testing.T) {
	col := [][]string{
		{"."},
		{"."},
		{"#"},
		{"#"},
		{"."},
		{"."},
		{"#"},
		{"#"},
		{"."},
	}

	testCases := []struct {
		x        int
		expected bool
	}{
		{0, false},
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.x), func(t *testing.T) {
			row := getColumn(col, 0)
			actual := isReflective(row, tc.x)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestPossibleReflections(t *testing.T) {
	row := []string{".", ".", "#", "#", ".", ".", "#", "#", "."}
	want := []int{1, 3, 5, 7}

	got := possibleReflections(row)

	if !slices.Equal(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestPossibleReflectionsColumn(t *testing.T) {
	col := [][]string{
		{"."},
		{"."},
		{"#"},
		{"#"},
		{"."},
		{"."},
		{"#"},
		{"#"},
		{"."},
	}

	want := []int{1, 3, 5, 7}

	got := possibleReflections(getColumn(col, 0))

	if !slices.Equal(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestColumnReflectionPoint(t *testing.T) {
	data, err := parse(mockData)

	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	want := 5

	got := columnReflectionPoint(data[0])

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestRowReflectionPoint(t *testing.T) {
	data, err := parse(mockData)

	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	want := 4

	got := rowReflectionPoint(data[1])

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestPatternScore(t *testing.T) {
	data, err := parse(mockData)

	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	testCases := []struct {
		pattern [][]string
		want    int
	}{
		{data[0], 5},
		{data[1], 400},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.pattern), func(t *testing.T) {
			got, err := patternScore(tc.pattern)
			if err != nil {
				t.Errorf("Error calculating pattern score: %v", err)
			}
			if got != tc.want {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestSumScore(t *testing.T) {
	data, err := parse(mockData)

	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	want := 405

	got := sumScore(data)

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestNonReflectivePoints(t *testing.T) {
	row := []string{".", ".", "#", "#", ".", ".", "#", "#", "."}

	testCases := []struct {
		x        int
		expected int
	}{
		{1, 0},
		{2, 2},
		{3, 0},
		{4, 4},
		{5, 0},
		{6, 3},
		{7, 0},
		{8, 1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.x), func(t *testing.T) {
			actual, err := nonReflectivePoints(row, tc.x)
			if err != nil {
				t.Errorf("Error calculating non-reflective points: %v", err)
			}
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestRowReflectionPointFixing(t *testing.T) {
	data, err := parse(mockData)
	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	want := 3
	got := rowReflectionPointFixing(data[0])
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}

	want = 1
	got = rowReflectionPointFixing(data[1])
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestColumnReflectionPointFixing(t *testing.T) {
	data := [][]string{
		{".", ".", "."},
		{"#", ".", "."},
		{"#", ".", "#"},
	}

	want := 2

	got := columnReflectionPointFixing(data)
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestPatterScoreFixing(t *testing.T) {
	data, err := parse(mockData)

	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	testCases := []struct {
		idx  int
		want int
	}{
		{0, 300},
		{1, 100},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.idx), func(t *testing.T) {
			got, err := patternScoreFixing(data[tc.idx])
			if err != nil {
				t.Errorf("Error calculating pattern score: %v", err)
			}
			if got != tc.want {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestSumScoreFixing(t *testing.T) {
	data, err := parse(mockData)

	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	want := 400

	got := sumScoreFixing(data)

	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
