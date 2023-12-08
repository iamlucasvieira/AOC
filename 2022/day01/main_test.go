package main

import "testing"

var mockData = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

func TestParseData(t *testing.T) {
	data, err := parseData(mockData)

	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	if len(data) != 5 {
		t.Errorf("Expected 5 slices, got %d", len(data))
	}
}

func TestTopThreeSum(t *testing.T) {
	data, err := parseData(mockData)

	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	sum := topThreeSum(data)

	if sum != 45000 {
		t.Errorf("Expected 45000, got %d", sum)
	}

}
