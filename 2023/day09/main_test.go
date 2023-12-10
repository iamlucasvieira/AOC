package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"slices"
	"testing"
)

var mockData = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func TestParseData(t *testing.T) {
	data, err := parseData(mockData)
	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	if len(data) != 3 {
		t.Errorf("Expected 3 lists, got %d", len(data))
	}

	if len(data[0]) != 6 {
		t.Errorf("Expected 6 numbers in first list, got %d", len(data[0]))
	}
}

func TestDiff(t *testing.T) {
	testCases := []utils.TestCase[[]int, []int]{
		{[]int{1, 2, 3, 4, 5}, []int{1, 1, 1, 1}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 1, 1, 1}},
		{[]int{1, 4, 7, 10, 20}, []int{3, 3, 3, 10}},
		{[]int{1, 5, 42, 95, 2940}, []int{4, 37, 53, 2845}},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got := diff(tc.Input)
			if !slices.Equal(got, tc.Expected) {
				t.Errorf("Expected %v, got %v", tc.Expected, got)
			}
		})
	}
}

func TestAllZeros(t *testing.T) {
	testCases := []utils.TestCase[[]int, bool]{
		{[]int{0, 0, 0, 0, 0}, true},
		{[]int{0, 0, 0, 0, 1}, false},
		{[]int{0, 0, 0, 1, 0}, false},
		{[]int{0, 0, 1, 0, 0}, false},
		{[]int{0, 1, 0, 0, 0}, false},
		{[]int{1, 0, 0, 0, 0}, false},
		{[]int{1, 1, 1, 1, 1}, false},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got := allZeros(tc.Input)
			if got != tc.Expected {
				t.Errorf("Expected %v, got %v", tc.Expected, got)
			}
		})
	}
}

func TestDiffIterationsEnd(t *testing.T) {
	testCases := []utils.TestCase[[]int, []int]{
		{[]int{0, 3, 6, 9, 12, 15}, []int{3, 0}},
		{[]int{1, 3, 6, 10, 15, 21}, []int{6, 1, 0}},
		{[]int{1, 3, 6, 10, 15, 21, 28}, []int{7, 1, 0}},
		{[]int{10, 13, 16, 21, 30, 45, 68}, []int{23, 8, 2, 0}},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got := diffIterations(tc.Input, true)
			if !slices.Equal(got, tc.Expected) {
				t.Errorf("Expected %v, got %v", tc.Expected, got)
			}
		})
	}
}

func TestDiffIterationsStart(t *testing.T) {
	testCases := []utils.TestCase[[]int, []int]{
		{[]int{0, 3, 6, 9, 12, 15}, []int{3, 0}},
		{[]int{1, 3, 6, 10, 15, 21}, []int{2, 1, 0}},
		{[]int{10, 13, 16, 21, 30, 45, 68}, []int{3, 0, 2, 0}},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got := diffIterations(tc.Input, false)
			if !slices.Equal(got, tc.Expected) {
				t.Errorf("Expected %v, got %v", tc.Expected, got)
			}
		})
	}
}

func TestNextNumber(t *testing.T) {
	testCases := []utils.TestCase[[]int, int]{
		{[]int{0, 3, 6, 9, 12, 15}, 18},
		{[]int{1, 3, 6, 10, 15, 21}, 28},
		{[]int{10, 13, 16, 21, 30, 45}, 68},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got := nextNumber(tc.Input)
			if got != tc.Expected {
				t.Errorf("Expected %v, got %v", tc.Expected, got)
			}
		})
	}
}

func TestSumNextNumbers(t *testing.T) {
	data, err := parseData(mockData)
	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	got := sumNextNumbers(data)
	expected := 114
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPreviousNumber(t *testing.T) {
	testCases := []utils.TestCase[[]int, int]{
		{[]int{0, 3, 6, 9, 12, 15}, -3},
		{[]int{1, 3, 6, 10, 15, 21}, 0},
		{[]int{10, 13, 16, 21, 30, 45}, 5},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got := previousNumber(tc.Input)
			if got != tc.Expected {
				t.Errorf("Expected %v, got %v", tc.Expected, got)
			}
		})
	}
}

func TestSumPreviousNumbers(t *testing.T) {
	data, err := parseData(mockData)
	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	got := sumPreviousNumbers(data)
	expected := 2
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
