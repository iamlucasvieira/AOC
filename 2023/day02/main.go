package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"regexp"
	"strconv"
	"strings"
)

// Game is a game of cubes
type game struct {
	id     int
	rounds []round
}

type round struct {
	id    int
	blue  int
	red   int
	green int
}

// newGame creates a new game from a string
func newGame(gameString string) (game, error) {
	// Split game and rounds by :
	gameParts := strings.Split(gameString, ":")

	// Define regexes
	gameIDRegex := regexp.MustCompile(`Game (\d+)`)
	colorRegex := regexp.MustCompile(`(\d+) (blue|red|green)`)

	gameID, err := strconv.Atoi(gameIDRegex.FindStringSubmatch(gameParts[0])[1])
	if err != nil {
		return game{}, fmt.Errorf("error parsing game ID: %v", err)
	}

	g := game{id: gameID}

	// Loop through rounds
	for i, roundString := range strings.Split(gameParts[1], ";") {
		r := round{id: i}
		// Get number of cubes of each color
		matches := colorRegex.FindAllStringSubmatch(roundString, -1)
		for _, match := range matches {
			count, err := strconv.Atoi(match[1])
			if err != nil {
				return game{}, fmt.Errorf("error parsing color count: %v", err)
			}
			switch match[2] {
			case "blue":
				r.blue = count
			case "red":
				r.red = count
			case "green":
				r.green = count
			}
		}
		g.rounds = append(g.rounds, r)
	}

	return g, nil
}

// isGameValid checks if a game is valid
func isGameValid(g game, nRed int, nGreen int, nBlue int) bool {
	// Loop through rounds
	for _, round := range g.rounds {
		if round.blue > nBlue {
			return false
		}
		if round.red > nRed {
			return false
		}
		if round.green > nGreen {
			return false
		}
	}
	return true
}

// part1 solves part 1 of day 2
func part1() {
	fmt.Println("Part 1:")
	data := utils.ReadFile("input.txt")
	sum := 0

	for _, gameString := range data {
		if gameString != "" {
			game, err := newGame(gameString)
			if err != nil {
				fmt.Printf("Error parsing game: %v\n", err)
			}
			if isGameValid(game, 12, 13, 14) {
				sum += game.id
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func fewestCubes(g game) (int, int, int) {
	var maxRed, maxGreen, maxBlue int

	for _, round := range g.rounds {
		if round.red > maxRed {
			maxRed = round.red
		}
		if round.green > maxGreen {
			maxGreen = round.green
		}
		if round.blue > maxBlue {
			maxBlue = round.blue
		}
	}
	return maxRed, maxGreen, maxBlue
}

// part2 solves part 2 of day 2
func part2() {
	fmt.Println("Part 2:")
	data := utils.ReadFile("input2.txt")
	sum := 0
	for _, gameString := range data {
		if gameString != "" {
			game, err := newGame(gameString)
			if err != nil {
				fmt.Printf("Error parsing game: %v\n", err)
			}
			red, green, blue := fewestCubes(game)
			sum += red * green * blue
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func main() {
	part1()
	part2()
}
