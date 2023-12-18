package main

import (
	"fmt"
	"testing"
)

var mockData = []string{
	"R 6 (#70c710)",
	"D 5 (#0dc571)",
	"L 2 (#5713f0)",
	"D 2 (#d2c081)",
	"R 2 (#59c680)",
	"D 2 (#411b91)",
	"L 5 (#8ceee2)",
	"U 2 (#caa173)",
	"L 1 (#1b58a2)",
	"U 2 (#caa171)",
	"R 2 (#7807d2)",
	"U 3 (#a77fa3)",
	"L 2 (#015232)",
	"U 2 (#7a21e3)",
}

// Goes to negative x and y
var mockData2 = []string{
	"L 2 (#70c710)",
	"U 2 (#0dc571)",
	"R 2 (#5713f0)",
	"D 2 (#d2c081)",
}

func TestParse(t *testing.T) {
	commands, err := parse(mockData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(commands) != 14 {
		t.Fatalf("expected 14 commands, got %d", len(commands))
	}

	if commands[0].direction != (point{X: 1}) {
		t.Fatalf("expected direction to be (1, 0), got %v", commands[0].direction)
	}

	if commands[0].steps != 6 {
		t.Fatalf("expected steps to be 6, got %d", commands[0].steps)
	}

	if commands[0].color != "#70c710" {
		t.Fatalf("expected color to be #70c710, got %s", commands[0].color)
	}
}

func TestBuildPath(t *testing.T) {
	commands, err := parse(mockData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	path := buildPath(commands)

	if len(path) != 38 {
		t.Fatalf("expected 38 points, got %d", len(path))
	}
}

func TestPolygonArea(t *testing.T) {
	commands, err := parse(mockData2)
	path := buildPath(commands)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	area := polygonArea(path[1:])
	if area != 4 {
		t.Fatalf("expected 4, got %d", area)
	}

}

func TestPickleTheorem(t *testing.T) {

	testCases := []struct {
		commands []string
		want     int
	}{
		{mockData, 62},
		{mockData2, 9},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.commands), func(t *testing.T) {
			commands, err := parse(tc.commands)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			path := buildPath(commands)
			got := pickleTheorem(path)
			if got != tc.want {
				t.Fatalf("expected %d, got %d", tc.want, got)
			}
		})
	}
}

func TestCommandsFromHex(t *testing.T) {
	commands, err := parse(mockData)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	testCases := []struct {
		index         int
		wantDirection point
		wantSteps     int
	}{
		{0, point{X: 1}, 461937},
		{1, point{Y: 1}, 56407},
		{2, point{X: 1}, 356671},
		{3, point{Y: 1}, 863240},
		{4, point{X: 1}, 367720},
		{5, point{Y: 1}, 266681},
		{6, point{X: -1}, 577262},
		{7, point{Y: -1}, 829975},
		{8, point{X: -1}, 112010},
		{9, point{Y: 1}, 829975},
		{10, point{X: -1}, 491645},
		{11, point{Y: -1}, 686074},
		{12, point{X: -1}, 5411},
		{13, point{Y: -1}, 500254},
	}
	var newCommands []command

	newCommands, err = commandsFromHex(commands)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(newCommands) != 14 {
		t.Fatalf("expected 14 commands, got %d", len(newCommands))
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.index), func(t *testing.T) {
			if newCommands[tc.index].direction != tc.wantDirection {
				t.Fatalf("expected direction to be %v, got %v", tc.wantDirection, newCommands[tc.index].direction)
			}

			if newCommands[tc.index].steps != tc.wantSteps {
				t.Fatalf("expected steps to be %d, got %d", tc.wantSteps, newCommands[tc.index].steps)
			}
		})
	}

}

func TestPickleTheoremHex(t *testing.T) {
	commands, err := parse(mockData)
	want := 952408144115
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	commands, err = commandsFromHex(commands)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	path := buildPath(commands)
	got := pickleTheorem(path)
	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}

func TestPart1(t *testing.T) {
	value := part1()
	want := 34329
	if value != want {
		t.Fatalf("expected %d, got %d", want, value)
	}
}
