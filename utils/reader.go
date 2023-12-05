package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadFile(fileName string) []string {
	// Get the caller's file path
	_, callerFilePath, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("Could not get caller's file path")
	}

	// Get the directory of the caller
	callerDir := filepath.Dir(callerFilePath)

	// Construct the full path of the target file
	fullPath := filepath.Join(callerDir, fileName)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	// Transform into list of strings (one string per line)
	lines := strings.Split(string(data), "\n")

	return lines
}
