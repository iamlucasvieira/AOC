package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strings"
)

// parse parses the input and returns slices for each column.
func parse(input []string) []string {
	var columns = make([][]string, len(input))
	var columnsStrings = make([]string, len(input))
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		data := strings.Split(line, "")

		for i, c := range data {
			columns[i] = append(columns[i], c)
		}
	}

	// Transform each row into a string
	for i, column := range columns {
		columnsStrings[i] = strings.Join(column, "")
	}

	return columnsStrings
}

type cache map[string]string

// table is a slice of slices of strings.
type table [][]string

// Row returns a string with the row at index i.
func (t table) Row(i int) string {
	return strings.Join(t[i], "")
}

// Column returns a string with the column at index i.
func (t table) Column(i int) string {
	var column []string
	for _, row := range t {
		column = append(column, row[i])
	}
	return strings.Join(column, "")
}

// replaceColumn replaces the column at index i with the string s.
func (t table) replaceColumn(i int, s string) {
	for j, c := range s {
		t[j][i] = string(c)
	}
}

// replaceRow replaces the row at index i with the string s.
func (t table) replaceRow(i int, s string) {
	t[i] = strings.Split(s, "")
}

// score returns the score of the table.
// Each O is worth the height it is at.
func (t table) score() int {
	var result int
	for i, row := range t {
		for _, c := range row {
			if c == "O" {
				result += len(t) - i
			}
		}
	}
	return result
}

func (t table) String() string {
	var result string
	for _, row := range t {
		result += strings.Join(row, "") + "\n"
	}
	return result
}

// parseAsTable parses the input and returns a table.
func parseAsTable(input []string) table {
	var columns [][]string
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		data := strings.Split(line, "")
		columns = append(columns, data)
	}

	return columns
}

// rocksIndex returns the index of all rocks in a column.
func rocksIndex(s string) []int {
	var result []int
	for i, c := range s {
		if c == 'O' {
			result = append(result, i)
		}
	}
	return result
}

func rocksImpactColumn(s string) int {
	rocks := rocksIndex(s)
	allImpact := 0
	for _, rock := range rocks {
		rockImpact := len(s) - rock
		for i := rock - 1; i >= 0; i-- {
			if s[i] == '#' {
				break
			} else if s[i] == '.' {
				rockImpact++
			}
		}
		allImpact += rockImpact
	}
	return allImpact
}

func allImpact(columns []string) int {
	var result int
	for _, column := range columns {
		result += rocksImpactColumn(column)
	}
	return result
}

// Moves all 0s to the left
func moveLeft(s string) string {
	rocks := rocksIndex(s)

	for _, rock := range rocks {
		idx := rock
		for i := rock - 1; i >= 0; i-- {
			if s[i] == '#' {
				break
			} else if s[i] == '.' {
				idx--
			}
		}
		s = s[:idx] + "O" + s[idx:rock] + s[rock+1:]
	}
	return s
}

// Moves all 0s to the right
func moveRight(s string) string {
	// Reverse string
	s = reverseString(s)
	// Move left
	s = moveLeft(s)
	// Reverse again
	s = reverseString(s)
	return s
}

func reverseString(s string) string {
	var result string
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func moveTableHorizontal(t table, c cache, moveFunc func(string) string, symbol string) table {
	for i := range t {
		row := t.Row(i)
		newRow, ok := c[row+symbol]

		if !ok {
			newRow = moveFunc(row)
			c[row+symbol] = newRow
		}
		t.replaceRow(i, newRow)
	}
	return t
}

func moveTableRight(t table, c cache) table {
	return moveTableHorizontal(t, c, moveRight, ">")
}

func moveTableLeft(t table, c cache) table {
	return moveTableHorizontal(t, c, moveLeft, "<")
}

func moveTableVertical(t table, c cache, moveFunc func(string) string, symbol string) table {
	for i := 0; i < len(t[0]); i++ {
		column := t.Column(i)
		newColumn, ok := c[column+symbol]

		if !ok {
			newColumn = moveFunc(column)
			c[column+symbol] = newColumn
		}
		t.replaceColumn(i, newColumn)
	}
	return t
}

func moveTableDown(t table, c cache) table {
	return moveTableVertical(t, c, moveRight, ">")
}

func moveTableUp(t table, c cache) table {
	return moveTableVertical(t, c, moveLeft, "<")
}

// Move table Up, Left, Down and Right
func cycle(t table, c cache) table {
	t = moveTableUp(t, c)
	t = moveTableLeft(t, c)
	t = moveTableDown(t, c)
	t = moveTableRight(t, c)
	return t
}

func cycleNTimes(t table, c cache, n int) table {

	type cachedTable struct {
		t      table
		cycles int
	}
	var cyclesCache = make(map[string]cachedTable)

	for i := 0; i < n; i++ {

		tableRows, ok := cyclesCache[t.String()]

		if ok {
			t = tableRows.t

			// If we found a cycle, we can skip the rest of the iterations

			// Cycle length
			cycleLength := i - tableRows.cycles

			// Cycles left
			cyclesLeft := n - i - 1

			// How many times we can skip the cycle
			skipCycles := cyclesLeft / cycleLength

			// Skip the cycles
			if cyclesLeft%cycleLength == 0 {
				i += skipCycles * cycleLength
			}

		}
		currentT := t.String()
		t = cycle(t, c)
		cyclesCache[currentT] = cachedTable{t, i}

	}
	return t
}

func part1() {
	fmt.Println("Part 1:")
	data := parse(utils.ReadFile("input.txt"))
	impact := allImpact(data)
	fmt.Printf("The impact is %d\n", impact)
}

func part2() {
	fmt.Println("Part 2:")
	data := parseAsTable(utils.ReadFile("input2.txt"))
	c := make(cache)
	result := cycleNTimes(data, c, 1000000000)
	fmt.Printf("The impact is %d\n", result.score())
}
func main() {
	part1()
	part2()
}
