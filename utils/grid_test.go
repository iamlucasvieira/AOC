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

func TestPoint_String(t *testing.T) {
	testCases := []struct {
		p    Point
		want string
	}{
		{Point{0, 0}, "(0, 0)"},
		{Point{1, 0}, "(1, 0)"},
		{Point{0, 1}, "(0, 1)"},
		{Point{1, 1}, "(1, 1)"},
		{Point{2, 2}, "(2, 2)"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.p), func(t *testing.T) {
			got := tc.p.String()
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

func TestGrid_Get(t *testing.T) {
	mockGrid := Grid[int]{
		{1, 2, 3},
		{4, 5, 6},
	}

	testCases := []struct {
		p   Point
		val int
	}{
		{Point{0, 0}, 1},
		{Point{1, 0}, 2},
		{Point{2, 0}, 3},
		{Point{0, 1}, 4},
		{Point{1, 1}, 5},
		{Point{2, 1}, 6},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.p), func(t *testing.T) {
			got := mockGrid.Get(tc.p)
			if got != tc.val {
				t.Errorf("got %v, want %v", got, tc.val)
			}
		})
	}
}

func TestGrid_Set(t *testing.T) {
	mockGrid := Grid[int]{
		{1, 2, 3},
		{4, 5, 6},
	}

	testCases := []struct {
		p   Point
		val int
	}{
		{Point{0, 0}, 10},
		{Point{1, 0}, 20},
		{Point{2, 0}, 30},
		{Point{0, 1}, 40},
		{Point{1, 1}, 50},
		{Point{2, 1}, 60},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.p), func(t *testing.T) {
			mockGrid.Set(tc.p, tc.val)
			got := mockGrid.Get(tc.p)
			if got != tc.val {
				t.Errorf("got %v, want %v", got, tc.val)
			}
		})
	}
}

var mockGrid = Grid[int]{
	{0, 0, 0, 0, 1, 1, 0, 0},
	{0, 1, 1, 1, 1, 1, 1, 0},
	{0, 1, 0, 0, 0, 0, 1, 0},
	{0, 1, 0, 1, 1, 0, 1, 0},
	{0, 1, 1, 1, 1, 1, 1, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
}

var mockPolygon = []Point{
	{1, 1},
	{1, 2},
	{1, 3},
	{1, 4},
	{2, 4},
	{3, 4},
	{3, 3},
	{4, 3},
	{4, 4},
	{5, 4},
	{6, 4},
	{6, 3},
	{6, 2},
	{6, 1},
	{5, 1},
	{5, 0},
	{4, 0},
	{4, 1},
	{3, 1},
	{2, 1},
	{1, 1},
}

func TestInsidePolygon(t *testing.T) {

	testCases := []struct {
		p    Point
		want bool
	}{
		{Point{0, 0}, false},
		{Point{1, 0}, false},
		{Point{2, 0}, false},
		{Point{3, 0}, false},
		{Point{4, 0}, false},
		{Point{5, 0}, false},
		{Point{6, 0}, false},
		{Point{7, 0}, false},
		{Point{0, 1}, false},
		{Point{1, 1}, false},
		{Point{2, 1}, false},
		{Point{3, 1}, false},
		{Point{4, 1}, false},
		{Point{5, 1}, false},
		{Point{6, 1}, false},
		{Point{7, 1}, false},
		{Point{0, 2}, false},
		{Point{1, 2}, false},
		{Point{2, 2}, true},
		{Point{3, 2}, true},
		{Point{4, 2}, true},
		{Point{5, 2}, true},
		{Point{6, 2}, false},
		{Point{7, 2}, false},
		{Point{0, 3}, false},
		{Point{1, 3}, false},
		{Point{2, 3}, true},
		{Point{3, 3}, false},
		{Point{4, 3}, false},
		{Point{5, 3}, true},
		{Point{6, 3}, false},
		{Point{7, 3}, false},
		{Point{0, 4}, false},
		{Point{1, 4}, false},
		{Point{2, 4}, false},
		{Point{3, 4}, false},
		{Point{4, 4}, false},
		{Point{5, 4}, false},
		{Point{6, 4}, false},
		{Point{7, 4}, false},
		{Point{0, 5}, false},
		{Point{1, 5}, false},
		{Point{2, 5}, false},
		{Point{3, 5}, false},
		{Point{4, 5}, false},
		{Point{5, 5}, false},
		{Point{6, 5}, false},
		{Point{7, 5}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.p), func(t *testing.T) {
			got, err := InsidePolygon(tc.p, mockGrid, mockPolygon)
			if err != nil {
				t.Errorf("got error %v", err)
			}
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestInsidePolygonWrongPolygon(t *testing.T) {
	// Polygon not closed
	if _, err := InsidePolygon(Point{}, mockGrid, mockPolygon[1:]); err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestInsidePolygonPointOutsideGrid(t *testing.T) {
	if _, err := InsidePolygon(Point{100, 100}, mockGrid, mockPolygon); err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestNewGrid(t *testing.T) {
	g := NewGrid[int](3, 2, 0)

	if got, want := g.Width(), 3; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := g.Height(), 2; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	for e := range g.Iterator() {
		if e.Value != 0 {
			t.Errorf("got %v, want 0", e.Value)
		}
	}
}
