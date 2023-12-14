package main

import (
	"fmt"
	"slices"
	"strconv"
	"testing"
)

var mockData = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func TestParseData(t *testing.T) {
	records, err := parseData(mockData)
	if err != nil {
		t.Errorf("parseData() returned error: %v", err)
	}

	if len(records) != 6 {
		t.Errorf("parseData() returned %d records, expected 6", len(records))
	}

	if len(records[0].strings) != 7 {
		t.Errorf("parseData() returned %d strings in first record, expected 7", len(records[0].strings))
	}

	if len(records[0].integers) != 3 {
		t.Errorf("parseData() returned %d integers in first record, expected 3", len(records[0].integers))
	}
}

func TestPossiblePattern(t *testing.T) {
	testCases := []struct {
		candidate []string
		pattern   []string
		want      bool
	}{
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "b", "?"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "?", "?"}, true},
		{[]string{"a", "b", "c"}, []string{"?", "?", "?"}, true},
		{[]string{"a", "b", "c"}, []string{"?", "?", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"?", "b", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "?", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "b", "d"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "d", "?"}, false},
		{[]string{"a", "b", "c"}, []string{"d", "?", "?"}, false},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := possiblePattern(tc.candidate, tc.pattern)
			if got != tc.want {
				t.Errorf("possiblePattern() returned %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestSliceSum(t *testing.T) {
	testCases := []struct {
		slice []int
		want  int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 10},
		{[]int{1, 2, 3, 4, 5}, 15},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := sliceSum(tc.slice)
			if got != tc.want {
				t.Errorf("sliceSum() returned %d, expected %d", got, tc.want)
			}
		})
	}
}

func TestRecord_String(t *testing.T) {
	testCases := []struct {
		record record
		want   string
	}{
		{record{[]string{"a", "b", "c"}, []int{1, 2, 3}}, "[a b c] [1 2 3]"},
		{record{[]string{"a", "b", "c"}, []int{1, 2, 3, 4}}, "[a b c] [1 2 3 4]"},
		{record{[]string{}, []int{}}, "[] []"},
		{record{[]string{""}, []int{}}, "[] []"},
		{record{[]string{}, []int{1}}, "[] [1]"},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := tc.record.String()
			if got != tc.want {
				t.Errorf("record.String() returned %s, expected %s", got, tc.want)
			}
		})
	}
}

func TestBuildCandidateString(t *testing.T) {
	testCases := []struct {
		nDot  int
		nHash int
		want  []string
	}{
		{0, 0, []string{"."}},
		{1, 0, []string{".", "."}},
		{0, 1, []string{"#", "."}},
		{1, 1, []string{".", "#", "."}},
		{2, 1, []string{".", ".", "#", "."}},
		{1, 2, []string{".", "#", "#", "."}},
		{2, 2, []string{".", ".", "#", "#", "."}},
		{3, 2, []string{".", ".", ".", "#", "#", "."}},
		{2, 3, []string{".", ".", "#", "#", "#", "."}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v %v", tc.nDot, tc.nHash), func(t *testing.T) {
			got := buildCandidateString(tc.nDot, tc.nHash)
			if !slices.Equal(got, tc.want) {
				t.Errorf("buildCandidateString() returned %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestCountMatches(t *testing.T) {
	testCases := []struct {
		r    record
		want int
	}{
		{record{[]string{"?", "?", "?", ".", "#", "#", "#"}, []int{1, 1, 3}}, 1},
		{record{[]string{".", "?", "?", ".", ".", "?", "?", ".", ".", ".", "?", "#", "#", "."}, []int{1, 1, 3}}, 4},
		{record{[]string{"?", "#", "?", "#", "?", "#", "?", "#", "?", "#", "?", "#", "?", "#", "?"}, []int{1, 3, 1, 6}}, 1},
		{record{[]string{"?", "?", "?", "?", ".", "#", ".", ".", ".", "#", "."}, []int{4, 1, 1}}, 1},
		{record{[]string{"?", "?", "?", "?", ".", "#", "#", "#", "#", "#", "#", ".", ".", "#", "#", "#", "#", "#", "."}, []int{1, 6, 5}}, 4},
		{record{[]string{"?", "#", "#", "#", "?", "?", "?", "?", "?", "?", "?", "?"}, []int{3, 2, 1}}, 10},
	}

	var cache = make(map[string]int)

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v %v", tc.r.strings, tc.r.integers), func(t *testing.T) {
			got := countMatches(cache, tc.r, len(tc.r.strings))
			if got != tc.want {
				t.Errorf("countMatches() returned %d, expected %d", got, tc.want)
			}
		})
	}
}

func TestSumOfMatches(t *testing.T) {
	records, err := parseData(mockData)
	if err != nil {
		t.Errorf("parseData() returned error: %v", err)
	}

	got := sumOfMatches(records)
	want := 21
	if got != want {
		t.Errorf("sumOfMatches() returned %d, expected %d", got, want)
	}
}

func TestUnfold5(t *testing.T) {
	n := 5

	testCases := []struct {
		r    record
		want record
	}{
		{
			record{
				strings:  []string{".", "#"},
				integers: []int{1},
			},
			record{
				strings:  []string{".", "#", "?", ".", "#", "?", ".", "#", "?", ".", "#", "?", ".", "#"},
				integers: []int{1, 1, 1, 1, 1},
			},
		},
		{
			record{
				strings:  []string{"?", "?", "?", ".", "#", "#", "#"},
				integers: []int{1, 1, 3},
			},
			record{
				strings:  []string{"?", "?", "?", ".", "#", "#", "#", "?", "?", "?", "?", ".", "#", "#", "#", "?", "?", "?", "?", ".", "#", "#", "#", "?", "?", "?", "?", ".", "#", "#", "#", "?", "?", "?", "?", ".", "#", "#", "#"},
				integers: []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v %v", tc.r.strings, tc.r.integers), func(t *testing.T) {
			got := unfold(tc.r, n)

			if !slices.Equal(got.strings, tc.want.strings) {
				t.Errorf("unfold() returned %v, expected %v", got.strings, tc.want.strings)
			}

			if !slices.Equal(got.integers, tc.want.integers) {
				t.Errorf("unfold() returned %v, expected %v", got.integers, tc.want.integers)
			}
		})
	}
}

func TestSumOfMatchesUnfolded(t *testing.T) {
	records, err := parseData(mockData)
	if err != nil {
		t.Errorf("parseData() returned error: %v", err)
	}

	records = unfoldAll(records, 5)
	got := sumOfMatches(records)
	want := 525152
	if got != want {
		t.Errorf("sumOfMatchesUnfolded() returned %d, expected %d", got, want)
	}
}
