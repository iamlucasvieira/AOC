package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	lines := []string{"a", "", "b"}
	result := parse(lines)
	if len(result) != 2 {
		t.Errorf("Expected 2, got %d", len(result))
	}
}

func TestFirstUniqueSequence(t *testing.T) {
	testCases := []struct {
		line string
		n    int
		want int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 4, 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11},
	}

	for _, tc := range testCases {
		t.Run(tc.line, func(t *testing.T) {
			got := firstUniqueSequence(tc.line, tc.n)
			if got != tc.want {
				t.Errorf("firstUniqueSequence(%q, %d) = %d; want %d", tc.line, tc.n, got, tc.want)
			}
		})
	}
}

func TestFirstUniqueSequence14(t *testing.T) {
	testCases := []struct {
		line string
		n    int
		want int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 15, 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 14, 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26},
	}

	for _, tc := range testCases {
		t.Run(tc.line, func(t *testing.T) {
			got := firstUniqueSequence(tc.line, tc.n)
			if got != tc.want {
				t.Errorf("firstUniqueSequence14(%q, %d) = %d; want %d", tc.line, tc.n, got, tc.want)
			}
		})
	}
}
