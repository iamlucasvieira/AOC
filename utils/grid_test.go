package utils

import (
	"fmt"
	"testing"
)

func TestPoint_Equal(t *testing.T) {
	testCases := []struct {
		name string
		p1   Point
		p2   Point
		want bool
	}{
		{
			name: "should return true when points are equal",
			p1:   Point{1, 1},
			p2:   Point{1, 1},
			want: true,
		},
		{
			name: "should return false when points are not equal",
			p1:   Point{1, 1},
			p2:   Point{1, 2},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.p1.Equal(tc.p2)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPoint_Add(t *testing.T) {
	testCases := []struct {
		p1   Point
		p2   Point
		want Point
	}{
		{
			p1:   Point{1, 1},
			p2:   Point{1, 1},
			want: Point{2, 2},
		},
		{
			p1:   Point{1, 1},
			p2:   Point{1, 2},
			want: Point{2, 3},
		},
		{
			p1:   Point{-1, 0},
			p2:   Point{1, 0},
			want: Point{0, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v+%v", tc.p1, tc.p2), func(t *testing.T) {
			got := tc.p1.Add(tc.p2)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPoint_Sub(t *testing.T) {
	testCases := []struct {
		p1   Point
		p2   Point
		want Point
	}{
		{
			p1:   Point{1, 1},
			p2:   Point{1, 1},
			want: Point{0, 0},
		},
		{
			p1:   Point{1, 1},
			p2:   Point{1, 2},
			want: Point{0, -1},
		},
		{
			p1:   Point{-1, 0},
			p2:   Point{1, 0},
			want: Point{-2, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v", tc.p1, tc.p2), func(t *testing.T) {
			got := tc.p1.Sub(tc.p2)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPoint_Mul(t *testing.T) {
	testCases := []struct {
		p1   Point
		p2   Point
		want Point
	}{
		{
			p1:   Point{1, 1},
			p2:   Point{1, 1},
			want: Point{1, 1},
		},
		{
			p1:   Point{1, 1},
			p2:   Point{1, 2},
			want: Point{1, 2},
		},
		{
			p1:   Point{-1, 0},
			p2:   Point{1, 0},
			want: Point{-1, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v*%v", tc.p1, tc.p2), func(t *testing.T) {
			got := tc.p1.Mul(tc.p2)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPoint_Distance(t *testing.T) {
	testCases := []struct {
		p1   Point
		p2   Point
		want int
	}{
		{
			p1:   Point{1, 1},
			p2:   Point{1, 1},
			want: 0,
		},
		{
			p1:   Point{1, 1},
			p2:   Point{1, 2},
			want: 1,
		},
		{
			p1:   Point{-1, 0},
			p2:   Point{1, 0},
			want: 2,
		},
		{
			p1:   Point{1, 6},
			p2:   Point{5, 11},
			want: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v", tc.p1, tc.p2), func(t *testing.T) {
			got := tc.p1.Distance(tc.p2)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGrid_Contains(t *testing.T) {
	grid := Grid[int]{
		{1, 2, 3},
		{4, 5, 6},
	}

	testCases := []struct {
		p    Point
		want bool
	}{
		{Point{0, 0}, true},
		{Point{1, 0}, true},
		{Point{2, 0}, true},
		{Point{0, 1}, true},
		{Point{1, 1}, true},
		{Point{2, 1}, true},
		{Point{0, 2}, false},
		{Point{1, 2}, false},
		{Point{2, 2}, false},
		{Point{3, 0}, false},
		{Point{3, 1}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.p), func(t *testing.T) {
			got := grid.Contains(tc.p)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGrid_ContainsXY(t *testing.T) {
	g := Grid[int]{
		{1, 2, 3},
		{4, 5, 6},
	}

	testCases := []struct {
		x    int
		y    int
		want bool
	}{
		{0, 0, true},
		{1, 0, true},
		{2, 0, true},
		{0, 1, true},
		{1, 1, true},
		{2, 1, true},
		{0, 2, false},
		{1, 2, false},
		{2, 2, false},
		{3, 0, false},
		{3, 1, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v,%v", tc.x, tc.y), func(t *testing.T) {
			got := g.ContainsXY(tc.x, tc.y)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGrid_Width(t *testing.T) {
	g := Grid[int]{
		{1, 2, 3},
		{4, 5, 6},
	}

	if got, want := g.Width(), 3; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	emptyGrid := Grid[int]{}
	if got, want := emptyGrid.Width(), 0; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGrid_Height(t *testing.T) {
	g := Grid[int]{
		{1, 2, 3},
		{4, 5, 6},
	}

	if got, want := g.Height(), 2; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	emptyGrid := Grid[int]{}
	if got, want := emptyGrid.Height(), 0; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGrid_Iterator(t *testing.T) {
	g := Grid[int]{
		{1, 2, 3},
		{4, 5, 6},
	}

	var sum int
	// Check if iterator goes over all items
	for e := range g.Iterator() {
		sum += e.Value
	}

	if got, want := sum, 21; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
