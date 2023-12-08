package utils

import (
	"os"
	"reflect"
	"testing"
)

// TestReadFile tests the ReadFile function
func TestReadFile(t *testing.T) {
	// Setup a test file
	testFileName := "testfile.txt"
	testContent := "line1\nline2\nline3"

	//testFilePath := filepath.Join(os.TempDir(), testFileName)
	err := os.WriteFile(testFileName, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %s", err)
	}

	// Ensure the file is removed after the test
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatalf("Failed to remove test file: %s", err)
		}
	}(testFileName)

	// Call the function
	result := ReadFile(testFileName)

	// Expected result
	expected := []string{"line1", "line2", "line3"}

	// Compare the result with the expected result
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ReadFile() = %v, want %v", result, expected)
	}
}
