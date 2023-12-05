package main

import (
	"github.com/iamlucasvieira/aoc/utils"
	"slices"
	"sort"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	testCases := []utils.TestCase[rune, bool]{
		{'0', true},
		{'1', true},
		{'5', true},
		{'9', true},
		{'a', false},
		{'z', false},
		{'A', false},
		{'$', false},
		{' ', false},
	}

	for _, tc := range testCases {
		t.Run(string(tc.Input), func(t *testing.T) {
			result := isNumeric(tc.Input)
			if result != tc.Expected {
				t.Errorf("isNumeric(%q) = %v; want %v", tc.Input, result, tc.Expected)
			}
		})
	}
}

func TestDecodeCalibration(t *testing.T) {
	testCases := []utils.TestCase[string, int]{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			result := decodeCalibration(tc.Input)
			if result != tc.Expected {
				t.Errorf("decodeCalibration(%q) = %v; want %v", tc.Input, result, tc.Expected)
			}
		})
	}
}

func TestSumCodes(t *testing.T) {
	type NewTestCase struct {
		utils.TestCase[[]string, int]
		Decode decoder
	}
	testCases := []NewTestCase{
		{
			TestCase: utils.TestCase[[]string, int]{
				Input:    []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
				Expected: 142,
			},
			Decode: decodeCalibration,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Input[0], func(t *testing.T) {
			result := sumCodes(tc.Input, decodeCalibration)
			if result != tc.Expected {
				t.Errorf("sumCodes(%q) = %v; want %v", tc.Input, result, tc.Expected)
			}
		})
	}
}

func TestDecodeCalibrationWritten(t *testing.T) {
	testCases := []utils.TestCase[string, int]{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			result := decodeCalibrationWritten(tc.Input)
			if result != tc.Expected {
				t.Errorf("decodeCalibrationWritten(%q) = %v; want %v", tc.Input, result, tc.Expected)
			}
		})
	}
}

func TestPossibleWordsEnds(t *testing.T) {
	testCases := []utils.TestCase[string, []string]{
		{"o", []string{"two"}},
		{"e", []string{"one", "three", "five", "nine"}},
		{"n", []string{"seven"}},
	}

	for _, tc := range testCases {
		t.Run(string(tc.Input), func(t *testing.T) {
			result := possibleWords(tc.Input, true)
			sort.Strings(result)
			sort.Strings(tc.Expected)
			if !slices.Equal(result, tc.Expected) {
				t.Errorf("possibleWords(%q) = %v; want %v", tc.Input, result, tc.Expected)
			}
		})
	}
}

func TestPossibleWordsStarts(t *testing.T) {
	testCases := []utils.TestCase[string, []string]{
		{"t", []string{"two", "three"}},
		{"o", []string{"one"}},
		{"f", []string{"four", "five"}},
	}

	for _, tc := range testCases {
		t.Run(string(tc.Input), func(t *testing.T) {
			result := possibleWords(tc.Input, false)
			sort.Strings(result)
			sort.Strings(tc.Expected)
			if !slices.Equal(result, tc.Expected) {
				t.Errorf("possibleWords(%q) = %v; want %v", tc.Input, result, tc.Expected)
			}
		})
	}
}
