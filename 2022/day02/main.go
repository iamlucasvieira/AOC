package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strings"
)

var opponent = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
}

var player = map[string]string{
	"X": "R",
	"Y": "P",
	"Z": "S",
}

// round returns the result of a single round of rock-paper-scissors.
// Returns:
// 3 if the players tie
// 6 if player wins
// 0 if opponent wins
// Plus 1 if player is "R", 2 if "P", 4 if "S"
func round(opponent, player string) int {
	if player == opponent {
		return 3
	}

	if player == "R" && opponent == "S" ||
		player == "S" && opponent == "P" ||
		player == "P" && opponent == "R" {
		return 6
	}

	return 0
}

// playerChoice returns the player's choice given the opponent's choice and an order
// Order
// S: Win
// P: Draw
// R: Lose
func playerChoice(opponent, order string) string {
	switch order {
	case "S": // Win
		switch opponent {
		case "R":
			return "P"
		case "P":
			return "S"
		case "S":
			return "R"
		}
	case "P": // Draw
		return opponent
	case "R": // Lose
		switch opponent {
		case "R":
			return "S"
		case "P":
			return "R"
		case "S":
			return "P"
		}
	}
	return ""
}

// scoreTemplate is the base function that scores multiple rounds of rock-paper-scissors.
func scoreTemplate(rounds [][2]string, playerMove func(string, string) string) int {
	var result int
	for _, r := range rounds {
		p := playerMove(r[0], r[1])
		result += round(r[0], p)

		switch p {
		case "R":
			result += 1
		case "P":
			result += 2
		case "S":
			result += 3
		}
	}
	return result
}

// score returns the score of multiple rounds of rock-paper-scissors.
func score(rounds [][2]string) int {
	pc := func(opponent, order string) string {
		return order
	}
	return scoreTemplate(rounds, pc)
}

// scorePlayerChoice returns the score of multiple rounds of rock-paper-scissors by first determining the player's choice.
func scorePlayerChoice(rounds [][2]string) int {
	return scoreTemplate(rounds, playerChoice)
}

func parse(lines []string) ([][2]string, error) {
	var result [][2]string
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		var p, o string
		var ok bool
		p, ok = player[parts[1]]
		if !ok {
			return nil, fmt.Errorf("invalid player: %s", parts[0])
		}

		o, ok = opponent[parts[0]]
		if !ok {
			return nil, fmt.Errorf("invalid opponent: %s", parts[1])
		}

		result = append(result, [2]string{o, p})
	}

	return result, nil
}

func part1() {
	fmt.Println("Part 1")
	rounds, err := parse(utils.ReadFile("input.txt"))
	if err != nil {
		panic(err)
	}
	s := score(rounds)
	fmt.Printf("Score: %d\n", s)
}

func part2() {
	fmt.Println("Part 2")
	rounds, err := parse(utils.ReadFile("input2.txt"))
	if err != nil {
		panic(err)
	}
	s := scorePlayerChoice(rounds)
	fmt.Printf("Score: %d\n", s)
}
func main() {
	part1()
	part2()
}
