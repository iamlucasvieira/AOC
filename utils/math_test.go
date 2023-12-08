package utils

import "testing"

func TestGCD(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{1, 1, 1},
		{2, 2, 2},
		{2, 4, 2},
		{4, 2, 2},
		{4, 6, 2},
		{6, 4, 2},
		{4, 8, 4},
		{8, 4, 4},
		{4, 12, 4},
		{12, 4, 4},
		{4, 16, 4},
		{16, 4, 4},
		{3, 5, 1},
		{5, 3, 1},
		{3, 6, 3},
		{6, 3, 3},
		{3, 9, 3},
		{9, 3, 3},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := GCD(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("GCD(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	testCases := []struct {
		integers []int
		want     int
	}{
		{[]int{1, 1}, 1},
		{[]int{2, 2}, 2},
		{[]int{2, 4}, 4},
		{[]int{4, 2}, 4},
		{[]int{4, 6}, 12},
		{[]int{6, 4}, 12},
		{[]int{4, 8}, 8},
		{[]int{8, 4}, 8},
		{[]int{4, 12}, 12},
		{[]int{12, 4}, 12},
		{[]int{4, 16}, 16},
		{[]int{16, 4}, 16},
		{[]int{3, 5}, 15},
		{[]int{5, 3}, 15},
		{[]int{3, 6}, 6},
		{[]int{6, 3}, 6},
		{[]int{3, 9, 7}, 63},
		{[]int{2, 5, 3, 7}, 210},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := LCM(tc.integers...)
			if got != tc.want {
				t.Errorf("LCM(%v) = %v, want %v", tc.integers, got, tc.want)
			}
		})
	}
}
