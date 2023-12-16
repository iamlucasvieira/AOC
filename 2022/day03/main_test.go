package main

import (
	"fmt"
	"testing"
)

var mockData = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestSharedItem(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", "p"},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "L"},
		{"PmmdzqPrVvPwwTWBwg", "P"},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "v"},
		{"ttgJtRGJQctTZtZT", "t"},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", "s"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := sharedItem(tc.input)
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func TestPriority(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"p", 16},
		{"L", 38},
		{"P", 42},
		{"v", 22},
		{"t", 20},
		{"s", 19},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := priority(tc.input)
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func TestPriorityOfSharedItems(t *testing.T) {
	testCases := []struct {
		input []string
		want  int
	}{
		{mockData, 157},
	}

	for _, tc := range testCases {
		t.Run("mockData", func(t *testing.T) {
			got := priorityOfSharedItems(tc.input)
			if got != tc.want {
				t.Errorf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func TestSharedItemStrings(t *testing.T) {
	testCases := []struct {
		input []string
		want  string
	}{
		{mockData[:3], "r"},
		{mockData[3:], "Z"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.want), func(t *testing.T) {
			got := sharedItemStrings(tc.input)
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func TestPriorityOfSharedItemsThree(t *testing.T) {
	sum := priorityOfSharedItemsThree(mockData)
	want := 70

	if sum != want {
		t.Errorf("got %d; want %d", sum, want)
	}
}
