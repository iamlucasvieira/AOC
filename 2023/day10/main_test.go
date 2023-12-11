package main

import (
	"fmt"
	"testing"
)

var mockData = []string{
	".....",
	".S-7.",
	".|.||",
	".L-J.",
	".....",
}

var mockData2 = []string{
	"..F7.",
	".FJ|.",
	"SJ.L7",
	"|F--J",
	"LJ...",
}

var mockData3 = []string{
	"...........",
	".S-------7.",
	".|F-----7|.",
	".||.....||.",
	".||.....||.",
	".|L-7.F-J|.",
	".|..|.|..|.",
	".L--J.L--J.",
	"...........",
}

var mockData4 = []string{
	".F----7F7F7F7F-7....",
	".|F--7||||||||FJ....",
	".||.FJ||||||||L7....",
	"FJL7L7LJLJ||LJ.L-7..",
	"L--J.L7...LJS7F-7L7.",
	"....F-J..F7FJ|L7L7L7",
	"....L7.F7||L7|.L7L7|",
	".....|FJLJ|FJ|F7|.LJ",
	"....FJL-7.||.||||...",
	"....L---J.LJ.LJLJ...",
}

var mockData5 = []string{
	"FF7FSF7F7F7F7F7F---7",
	"L|LJ||||||||||||F--J",
	"FL-7LJLJ||||||LJL-77",
	"F--JF--7||LJLJ7F7FJ-",
	"L---JF-JLJ.||-FJLJJ7",
	"|F|F-JF---7F7-L7L|7|",
	"|FFJF7L7F-JF7|JL---7",
	"7-L-JL7||F7|L7F-7F7|",
	"L.L7LFJ|||||FJL7||LJ",
	"L7JLJL-JLJLJL--JLJ.L",
}

func TestMakePiece(t *testing.T) {
	testCases := []struct {
		name            string
		point           point
		wantConnections []point
	}{
		{
			name:            "L",
			point:           point{3, 4},
			wantConnections: []point{{4, 4}, {3, 3}},
		},
		{
			name:            "7",
			point:           point{3, 7},
			wantConnections: []point{{2, 7}, {3, 8}},
		},
		{
			name:            "J",
			point:           point{1, 30},
			wantConnections: []point{{0, 30}, {1, 29}},
		},
		{
			name:            "F",
			point:           point{24, 18},
			wantConnections: []point{{24, 19}, {25, 18}},
		},
		{
			name:            "|",
			point:           point{5, 8},
			wantConnections: []point{{5, 7}, {5, 9}},
		},
		{
			name:            "-",
			point:           point{20, 30},
			wantConnections: []point{{19, 30}, {21, 30}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := makePiece(tc.name, tc.point)
			if err != nil {
				t.Errorf("Error creating piece: %v", err)
			}

			for _, connection := range got.connections {
				if connection != tc.wantConnections[0] && connection != tc.wantConnections[1] {
					t.Errorf("Expected %v to be in %v", connection, tc.wantConnections)
				}
			}
		})
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		a    point
		b    point
		want point
	}{
		{point{1, 2}, point{3, 4}, point{4, 6}},
		{point{-1, 2}, point{3, -4}, point{2, -2}},
		{point{0, 2}, point{3, 0}, point{3, 2}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := tc.a.add(tc.b)
			if got != tc.want {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		a    point
		b    point
		want bool
	}{
		{point{1, 2}, point{1, 2}, true},
		{point{1, 2}, point{2, 1}, false},
		{point{1, 2}, point{1, 3}, false},
		{point{1, 2}, point{3, 2}, false},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got := tc.a.equal(tc.b)
			if got != tc.want {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestGetNext(t *testing.T) {
	testCases := []struct {
		from     point
		name     string
		position point
		want     point
	}{
		{point{1, 2}, "L", point{1, 3}, point{2, 3}},
		{point{1, 2}, "L", point{0, 2}, point{0, 1}},
		{point{1, 2}, "7", point{2, 2}, point{2, 3}},
		{point{1, 2}, "7", point{1, 1}, point{0, 1}},
		{point{1, 2}, "J", point{2, 2}, point{2, 1}},
		{point{1, 2}, "J", point{1, 3}, point{0, 3}},
		{point{1, 2}, "F", point{0, 2}, point{0, 3}},
		{point{1, 2}, "F", point{1, 1}, point{2, 1}},
		{point{1, 2}, "|", point{1, 1}, point{1, 0}},
		{point{1, 2}, "|", point{1, 3}, point{1, 4}},
		{point{1, 2}, "-", point{0, 2}, point{-1, 2}},
		{point{1, 2}, "-", point{2, 2}, point{3, 2}},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got, err := getNext(tc.from, tc.name, tc.position)
			if err != nil {
				t.Errorf("Error getting next: %v", err)
			}
			if !got.equal(tc.want) {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestHasConnection(t *testing.T) {
	lPiece, _ := makePiece("L", point{1, 1})
	testCases := []struct {
		p2   point
		want bool
	}{
		{point{1, 2}, false},
		{point{1, 0}, true},
		{point{2, 1}, true},
		{point{0, 1}, false},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			got := lPiece.hasConnection(tc.p2)
			if got != tc.want {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestParseData(t *testing.T) {
	data, start, err := parseData(mockData)
	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	if len(data) != 5 {
		t.Errorf("Expected 3 lists, got %d", len(data))
	}

	if len(data[0]) != 5 {
		t.Errorf("Expected 6 numbers in first list, got %d", len(data[0]))
	}

	if !start.equal(point{1, 1}) {
		t.Errorf("Expected start to be %v, got %v", point{1, 1}, start)
	}
}

func TestFindFirstStep(t *testing.T) {
	data, start, err := parseData(mockData)
	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	var got point
	got, err = findFirstStep(data, start)
	if err != nil {
		t.Errorf("Error finding first step: %v", err)
	}

	if !got.equal(point{1, 2}) {
		t.Errorf("Expected first step to be %v, got %v", point{1, 2}, got)
	}
}

func TestPath(t *testing.T) {
	data, start, err := parseData(mockData)
	if err != nil {
		t.Errorf("Error parsing data: %v", err)
	}

	var got []point
	got, err = path(data, start)
	if err != nil {
		t.Errorf("Error finding first step: %v", err)
	}

	want := []point{
		{1, 1},
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 3},
		{3, 2},
		{3, 1},
		{2, 1},
		{1, 1},
	}

	for idx, p := range want {
		if !p.equal(got[idx]) {
			t.Errorf("Expected %v, got %v", p, got[idx])
		}
	}
}

func TestLargestDistance(t *testing.T) {

	testCases := []struct {
		data []string
		want int
	}{
		{mockData, 4},
		{mockData2, 8},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			data, start, err := parseData(tc.data)
			if err != nil {
				t.Errorf("Error parsing data: %v", err)
			}

			var got int
			got, err = largestDistance(data, start)
			if err != nil {
				t.Errorf("Error finding largest distance: %v", err)
			}

			if got != tc.want {
				t.Errorf("Expected %d, got %d", tc.want, got)
			}
		})
	}
}

func TestPointInPolygon(t *testing.T) {

	type values []struct {
		p    point
		want bool
	}

	testCases := []struct {
		data   []string
		values values
	}{
		{mockData,
			values{
				{point{0, 0}, false},
				{point{2, 2}, true},
				{point{4, 1}, false},
			},
		},
		{
			mockData2,
			values{
				{point{0, 0}, false},
				{point{2, 2}, true},
			},
		},
		{
			mockData3,
			values{
				{point{0, 0}, false},
				{point{3, 3}, false},
				{point{5, 5}, false},
				{point{2, 6}, true},
				{point{7, 6}, true},
			},
		},
		{
			mockData4,
			values{
				{point{0, 0}, false},
				{point{3, 2}, false},
				{point{4, 4}, false},
				{point{7, 4}, true},
				{point{6, 6}, true},
				{point{14, 6}, true},
				{point{0, 6}, false},
				{point{9, 8}, false},
			},
		},
		{
			mockData5,
			values{
				{point{10, 4}, true},
				{point{1, 8}, false},
				{point{14, 3}, true},
			},
		},
	}
	for idx, tc := range testCases {
		data, start, err := parseData(tc.data)

		if err != nil {
			t.Errorf("Error parsing data: %v", err)
		}

		var polygon []point
		polygon, err = path(data, start)
		if err != nil {
			t.Errorf("Error finding first step: %v", err)
		}

		for idx2, v := range tc.values {
			t.Run(fmt.Sprintf("TC %d.%d", idx, idx2), func(t *testing.T) {
				got := pointInPolygon(v.p, data, polygon)
				if got != v.want {
					t.Errorf("Expected (%v) %v, got %v", v.p, v.want, got)
				}
			})
		}
	}

}

func TestTilesEnclosed(t *testing.T) {
	testCases := []struct {
		data []string
		want int
	}{
		{mockData, 1},
		{mockData2, 1},
		{mockData3, 4},
		{mockData4, 8},
		{mockData5, 10},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			data, start, err := parseData(tc.data)
			if err != nil {
				t.Errorf("Error parsing data: %v", err)
			}

			var polygon []point
			polygon, err = path(data, start)
			if err != nil {
				t.Errorf("Error finding polygon: %v", err)
			}

			var got int
			got = tilesEnclosed(data, polygon)

			if got != tc.want {
				t.Errorf("Expected %d, got %d", tc.want, got)
			}
		})
	}
}
