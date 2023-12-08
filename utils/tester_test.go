package utils

import (
	"testing"
)

// TestTestCaseInt tests the TestCase struct with integers
func TestTestCaseInt(t *testing.T) {
	tc := TestCase[int, int]{
		Input:    5,
		Expected: 10,
	}

	if tc.Input != 5 {
		t.Errorf("Expected Input to be 5, got %d", tc.Input)
	}

	if tc.Expected != 10 {
		t.Errorf("Expected Expected to be 10, got %d", tc.Expected)
	}
}

// TestTestCaseString tests the TestCase struct with strings
func TestTestCaseString(t *testing.T) {
	tc := TestCase[string, string]{
		Input:    "hello",
		Expected: "world",
	}

	if tc.Input != "hello" {
		t.Errorf("Expected Input to be 'hello', got %s", tc.Input)
	}

	if tc.Expected != "world" {
		t.Errorf("Expected Expected to be 'world', got %s", tc.Expected)
	}
}

// TestTestCaseMixed tests the TestCase struct with mixed types
func TestTestCaseMixed(t *testing.T) {
	tc := TestCase[string, int]{
		Input:    "test",
		Expected: 42,
	}

	if tc.Input != "test" {
		t.Errorf("Expected Input to be 'test', got %s", tc.Input)
	}

	if tc.Expected != 42 {
		t.Errorf("Expected Expected to be 42, got %d", tc.Expected)
	}
}
