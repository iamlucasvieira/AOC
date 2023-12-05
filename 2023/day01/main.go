package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"log"
	"strconv"
	"strings"
)

var wordsMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// isNumeric returns true if the given character is a number
func isNumeric(c rune) bool {
	return c >= '0' && c <= '9'
}

// decodeCalibration decodes a calibration code and returns the resulting value
func decodeCalibration(encodedCode string) int {
	var code string

	runes := []rune(encodedCode)

	// Find first numeric character in the code
	for _, c := range runes {
		if isNumeric(c) {
			code += string(c)
			break // stop after the first match
		}
	}

	// Find last numeric character in the code
	for i := len(runes) - 1; i >= 0; i-- {
		if isNumeric(runes[i]) {
			code += string(runes[i])
			break // stop after the first match
		}
	}

	// Convert string to int
	codeInt, err := strconv.Atoi(code)
	if err != nil {
		log.Fatal(err)
	}

	return codeInt

}

type decoder func(string) int

// sumCodes sums the values of the given codes. The decoder function is used to decode each code
func sumCodes(codes []string, decode decoder) int {
	var sum int = 0

	for _, code := range codes {

		if len(code) == 0 {
			continue
		}
		sum += decode(code)
	}

	return sum
}

func possibleWords(character string, endsWith bool) []string {

	var words []string

	for word := range wordsMap {
		if endsWith && !strings.HasSuffix(word, character) {
			continue
		} else if !endsWith && !strings.HasPrefix(word, character) {
			continue
		}
		words = append(words, word)
	}
	return words
}

// return true if the given word is present at the given index of the encoded code
func isPresentAt(encodedCode string, word string, i int) bool {
	// Check length
	if len(encodedCode) < i+len(word) || i < 0 {
		return false
	}
	return encodedCode[i:i+len(word)] == word
}

// decodeCalibrationWritten decodes a calibration code in case written numbers are part of the encoded code
func decodeCalibrationWritten(encodedCode string) int {

	var code string

	// Loop through index and character of encoded code
	for i, c := range encodedCode {

		if isNumeric(c) {
			// If the character is numeric, skip it
			code += string(c)
			break
		}

		// Get possible words for the character
		possibleWords := possibleWords(string(c), false)

		for _, word := range possibleWords {

			if isPresentAt(encodedCode, word, i) {
				code += strconv.Itoa(wordsMap[word])
				break
			}
		}

		if len(code) == 1 {
			break
		}
	}

	runes := []rune(encodedCode)

	// Get Final Number
	for i := len(encodedCode) - 1; i >= 0; i-- {
		if isNumeric(runes[i]) {
			code += string(runes[i])
			break // stop after the first match
		}

		// Get possible words for the character
		possibleWords := possibleWords(string(runes[i]), true)

		for _, word := range possibleWords {

			if isPresentAt(encodedCode, word, i-len(word)+1) {
				code += strconv.Itoa(wordsMap[word])
				break
			}
		}

		if len(code) == 2 {
			break
		}
	}

	// Convert string to int
	codeInt, err := strconv.Atoi(code)
	if err != nil {
		log.Fatal(err)
	}
	return codeInt
}

// part1 solves part 1 of challenge
func part1() {
	fmt.Println("Running part 1")
	data := utils.ReadFile("input.txt")
	totalSum := sumCodes(data, decodeCalibration)
	fmt.Printf("Total sum: %d\n", totalSum)
}

// part2 solves part 2 of challenge
func part2() {
	fmt.Println("Running part 2")
	data := utils.ReadFile("input2.txt")
	totalSum := sumCodes(data, decodeCalibrationWritten)
	fmt.Printf("Total sum: %d\n", totalSum)
}

func main() {
	part1()
	part2()
}
