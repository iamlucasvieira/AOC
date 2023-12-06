package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"slices"
	"sort"
	"testing"
)

var mockData = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestFindRoots(t *testing.T) {
	var tests = []struct {
		a, b, c float64
		x1, x2  float64
	}{
		{1, 0, 0, 0, 0},
		{1, 0, -1, 1, -1},
		{1, 0, 1, 0, 0},
	}

	for _, test := range tests {
		x1, x2, err := findRoots(test.a, test.b, test.c)
		if x1 != test.x1 || x2 != test.x2 {
			t.Errorf("findRoots(%f, %f, %f) = %f, %f, %v", test.a, test.b, test.c, x1, x2, err)
		}
	}
}

func TestDistance(t *testing.T) {
	testCases := []utils.TestCase[[]int, int]{
		{[]int{7, 0}, 0},
		{[]int{7, 1}, 6},
		{[]int{7, 2}, 10},
		{[]int{7, 3}, 12},
		{[]int{7, 4}, 12},
		{[]int{7, 5}, 10},
		{[]int{7, 6}, 6},
		{[]int{7, 7}, 0},
	}

	for _, tc := range testCases {
		speedString := fmt.Sprintf("speed %d", tc.Input[1])
		t.Run(speedString, func(t *testing.T) {
			if got := distance(tc.Input[0], tc.Input[1]); got != tc.Expected {
				t.Errorf("distance(%d, %d) = %d, want %d", tc.Input[0], tc.Input[1], got, tc.Expected)
			}
		})
	}
}

func TestFindSpeed(t *testing.T) {
	testCases := []utils.TestCase[[]int, []float64]{
		{[]int{7, 0}, []float64{0, 7}},
		{[]int{7, 6}, []float64{1, 6}},
		{[]int{7, 10}, []float64{2, 5}},
		{[]int{7, 12}, []float64{3, 4}},
	}

	for _, tc := range testCases {
		speedString := fmt.Sprintf("speed %d", tc.Input[1])
		t.Run(speedString, func(t *testing.T) {
			speed1, speed2, err := findSpeed(tc.Input[0], tc.Input[1])
			if err != nil {
				t.Errorf("findSpeed(%d, %d) = %v", tc.Input[0], tc.Input[1], err)
			}
			if speed1 != tc.Expected[0] || speed2 != tc.Expected[1] {
				t.Errorf("findSpeed(%d, %d) = %f, %f, want %f, %f", tc.Input[0], tc.Input[1], speed1, speed2, tc.Expected[0], tc.Expected[1])
			}
		})
	}
}

func TestBestSpeeds(t *testing.T) {
	testCases := []utils.TestCase[[]int, []int]{
		{[]int{7, 9}, []int{2, 3, 4, 5}},
		{[]int{15, 40}, []int{4, 5, 6, 7, 8, 9, 10, 11}},
		{[]int{30, 200}, []int{11, 12, 13, 14, 15, 16, 17, 18, 19}},
	}

	for _, tc := range testCases {
		speedString := fmt.Sprintf("speed %d", tc.Input[1])
		t.Run(speedString, func(t *testing.T) {
			speeds, err := bestSpeeds(tc.Input[0], tc.Input[1])
			// sort the speeds
			sort.Ints(speeds)
			sort.Ints(tc.Expected)
			if err != nil {
				t.Errorf("bestSpeeds(%d, %d) = %v", tc.Input[0], tc.Input[1], err)
			}
			if !slices.Equal(speeds, tc.Expected) {
				t.Errorf("bestSpeeds(%d, %d) = %v, want %v", tc.Input[0], tc.Input[1], speeds, tc.Expected)
			}
		})
	}
}

func TestParseData(t *testing.T) {
	data := parseData(mockData)

	if len(data) != 3 {
		t.Errorf("Expected data length 3, got %v", len(data))
	}

	if data[0].time != 7 || data[0].distance != 9 {
		t.Errorf("Expected data[0] to be {7, 9}, got %v", data[0])
	}

	if data[1].time != 15 || data[1].distance != 40 {
		t.Errorf("Expected data[1] to be {15, 40}, got %v", data[1])
	}

	if data[2].time != 30 || data[2].distance != 200 {
		t.Errorf("Expected data[2] to be {30, 200}, got %v", data[2])
	}
}

func TestProductNumberBestSpeeds(t *testing.T) {
	data := parseData(mockData)

	if got := productNumberBestSpeeds(data); got != 288 {
		t.Errorf("productNumberBestSpeeds(%v) = %d, want %d", data, got, 288)
	}
}

func TestBuildSingleRace(t *testing.T) {
	data := parseData(mockData)
	singleRace := buildSingleRace(data)

	if singleRace.time != 71530 {
		t.Errorf("Expected singleRace.time to be 71530, got %d", singleRace.time)
	}

	if singleRace.distance != 940200 {
		t.Errorf("Expected singleRace.distance to be 940200, got %d", singleRace.distance)
	}
}

func TestTotalBestSpeeds(t *testing.T) {
	data := parseData(mockData)
	singleRace := buildSingleRace(data)
	speeds, err := bestSpeeds(singleRace.time, singleRace.distance)

	if err != nil {
		t.Errorf("Error getting best speeds: %v", err)
	}

	if len(speeds) != 71503 {
		t.Errorf("Expected speeds length 71503, got %d", len(speeds))
	}
}
