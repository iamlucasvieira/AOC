package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
)

type grid = utils.Grid[string]
type point = utils.Point

func parseData(data []string) grid {
	var parsedData grid

	for _, line := range data {
		var row []string

		if len(line) == 0 {
			continue
		}

		for _, char := range line {
			row = append(row, string(char))
		}
		parsedData = append(parsedData, row)
	}

	return parsedData
}

// rowsColsWithoutGalaxy is a function that returns the rows and columns idx of a grid that have no galaxy.
func rowsColsWithoutGalaxy(g grid, galaxies []point) (rows, cols []int) {
	w, h := g.Width(), g.Height()
	rows, cols = make([]int, 0, h), make([]int, 0, w)

	hasGalaxy := func(_galaxies []point, value int, inY bool) bool {
		for _, galaxy := range _galaxies {
			var galaxyCoordinate int
			if inY {
				galaxyCoordinate = galaxy.Y
			} else {
				galaxyCoordinate = galaxy.X
			}

			if galaxyCoordinate == value {
				return true
			}
		}
		return false
	}
	for y := 0; y < h; y++ {
		if hasGalaxy(galaxies, y, true) {
			continue
		}
		rows = append(rows, y)
	}

	for x := 0; x < w; x++ {
		if hasGalaxy(galaxies, x, false) {
			continue
		}
		cols = append(cols, x)
	}

	return rows, cols
}

func getGalaxies(g grid) []point {
	var galaxies []point
	for e := range g.Iterator() {
		if e.Value == "#" {
			galaxies = append(galaxies, e.Index)
		}
	}
	return galaxies
}

// appendColumn is a function that appends a column to a grid.
func appendColumn(g grid, col int, colValues []string) grid {
	newG := make(grid, len(g))

	for y, row := range g {
		// Create a new row with space for the new element
		newRow := make([]string, 0, len(row)+1)

		// Add the elements before the new column
		newRow = append(newRow, row[:col]...)

		// Add the new column value
		newRow = append(newRow, colValues[y])

		// Add the elements after the new column
		newRow = append(newRow, row[col:]...)

		// Assign the new row to the new grid
		newG[y] = newRow
	}
	return newG
}

// appendRow is a function that appends a row to a grid.
func appendRow(g grid, row int, rowValues []string) grid {
	newG := make(grid, 0, len(g)+1)

	// Add the elements before the new row
	newG = append(newG, g[:row]...)

	// Add the new row
	newG = append(newG, rowValues)

	// Add the elements after the new row
	newG = append(newG, g[row:]...)

	return newG
}

// allPairs is a function that returns all pairs of a slice without repetition.
func allPairs(s []point) [][]point {
	var pairs [][]point
	for i, v := range s {
		for j := i + 1; j < len(s); j++ {
			pairs = append(pairs, []point{v, s[j]})
		}
	}
	return pairs
}

func sumAllDistances(g grid) int {
	galaxies := getGalaxies(g)

	gPairs := allPairs(galaxies)

	var sum int

	for _, pair := range gPairs {
		sum += pair[0].Distance(pair[1])
	}

	return sum
}

// sumDistancesVirtualExpand is a function that sums all distances between galaxies in a grid with virtual expansion.
// n is the number of rows and columns replacing empty rows and columns
func sumDistancesVirtualExpand(g grid, n int) int {
	galaxies := getGalaxies(g)
	rows, cols := rowsColsWithoutGalaxy(g, getGalaxies(g))
	var sum int

	pairs := allPairs(galaxies)
	for _, pair := range pairs {
		sum += distance(rows, cols, pair[0], pair[1], n)
	}

	return sum
}

// expand appends a row or column of "." to rows and columns without galaxies.
func expand(g grid) grid {

	galaxies := getGalaxies(g)

	rows, cols := rowsColsWithoutGalaxy(g, galaxies)

	for idx, row := range rows {
		// The row is made of "." with the same width as the grid
		rowValues := make([]string, g.Width())
		for i := range rowValues {
			rowValues[i] = "."
		}
		g = appendRow(g, row+idx, rowValues)
	}

	for idx, col := range cols {
		// The column is made of "." with the same height as the grid
		colValues := make([]string, g.Height())
		for i := range colValues {
			colValues[i] = "."
		}
		g = appendColumn(g, col+idx, colValues)
	}

	return g
}

// distance is a function that returns the distance between two points, if rows and columns in between have no galaxy they add n.
func distance(rows, cols []int, p1, p2 point, n int) int {
	minX := min(p1.X, p2.X)
	maxX := max(p1.X, p2.X)
	minY := min(p1.Y, p2.Y)
	maxY := max(p1.Y, p2.Y)

	var d = p1.Distance(p2)

	for _, row := range rows {
		if row < minY {
			continue
		}

		if row > maxY {
			break
		}

		d += n - 1 // -1 because the initial empty row is already counted
	}

	for _, col := range cols {
		if col < minX {
			continue
		}

		if col > maxX {
			break
		}
		d += n - 1 // -1 because the initial empty column is already counted
	}

	return d
}

// expandAndSumAllDistances is a function that expands a grid and sums all distances between galaxies.
func expandAndSumAllDistances(g grid) int {
	g = expand(g)
	return sumAllDistances(g)
}

func part1() {
	fmt.Print("Part 1: ")
	data := parseData(utils.ReadFile("input.txt"))
	sum := expandAndSumAllDistances(data)
	fmt.Printf("The sum of all distances is %d\n", sum)
}

func part2() {
	fmt.Print("Part 2: ")
	data := parseData(utils.ReadFile("input2.txt"))
	sum := sumDistancesVirtualExpand(data, 1000000)
	fmt.Printf("The sum of all distances is %d\n", sum)
}
func main() {
	part1()
	part2()
}
