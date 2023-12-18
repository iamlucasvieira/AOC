package main

import "testing"

var mockData = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestParse(t *testing.T) {
	data, err := parse(mockData)
	if err != nil {
		t.Errorf("parse() error = %v", err)
	}

	if len(data) != len(mockData) {
		t.Errorf("parse() len(data) = %v, want %v", len(data), len(mockData))
	}

	if data[0][0].start != 2 {
		t.Errorf("parse() data[0][0].start = %v, want %v", data[0][0].start, 2)
	}

	if data[0][0].end != 4 {
		t.Errorf("parse() data[0][0].end = %v, want %v", data[0][0].end, 4)
	}

	if data[0][1].start != 6 {
		t.Errorf("parse() data[0][1].start = %v, want %v", data[0][1].start, 6)
	}

	if data[0][1].end != 8 {
		t.Errorf("parse() data[0][1].end = %v, want %v", data[0][1].end, 8)
	}
}

func TestIsWithin(t *testing.T) {
	want := []bool{false, false, false, true, true, false}
	data, err := parse(mockData)

	if err != nil {
		t.Errorf("parse() error = %v", err)
	}

	for i, d := range data {
		t.Run(mockData[i], func(t *testing.T) {
			if got := isWithin(d); got != want[i] {
				t.Errorf("isWithin() = %v, want %v", got, want[i])
			}
		})
	}
}

func TestNWithin(t *testing.T) {
	want := 2
	data, err := parse(mockData)

	if err != nil {
		t.Errorf("parse() error = %v", err)
	}

	if got := nWithin(data); got != want {
		t.Errorf("nWithin() = %v, want %v", got, want)
	}
}

func TestHasOverlap(t *testing.T) {
	want := []bool{false, false, true, true, true, true}

	data, err := parse(mockData)

	if err != nil {
		t.Errorf("parse() error = %v", err)
	}

	for i, d := range data {
		t.Run(mockData[i], func(t *testing.T) {
			if got := hasOverlap(d); got != want[i] {
				t.Errorf("hasOverlap() = %v, want %v", got, want[i])
			}
		})
	}
}

func TestNOverlap(t *testing.T) {
	want := 4
	data, err := parse(mockData)

	if err != nil {
		t.Errorf("parse() error = %v", err)
	}

	if got := nOverlap(data); got != want {
		t.Errorf("nOverlap() = %v, want %v", got, want)
	}
}

func TestPart1(t *testing.T) {
	want := 651
	got := part1()
	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 956
	got := part2()
	if got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}
