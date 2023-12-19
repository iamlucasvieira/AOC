package main

import "testing"

var mockData = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestParse(t *testing.T) {
	stacks, rules, err := parse(mockData)
	if err != nil {
		t.Errorf("parse: %s", err)
	}

	if len(stacks) != 3 {
		t.Errorf("parse: Expected 3 stacks, got %d", len(stacks))
	}

	if len(rules) != 4 {
		t.Errorf("parse: Expected 4 rules, got %d", len(rules))
	}

	if len(stacks[0]) != 2 {
		t.Errorf("parse: Expected 1 box in stack 1, got %d", len(stacks[0]))
	}

	if len(stacks[1]) != 3 {
		t.Errorf("parse: Expected 1 box in stack 2, got %d", len(stacks[1]))
	}

	if len(stacks[2]) != 1 {
		t.Errorf("parse: Expected 1 box in stack 3, got %d", len(stacks[2]))
	}

	v, ok := stacks[0].Pop()

	if !ok {
		t.Errorf("parse: Expected a value, got none")
	}

	if v != "N" {
		t.Errorf("parse: Expected D, got %s", v)
	}

	v, ok = stacks[0].Pop()

	if !ok {
		t.Errorf("parse: Expected a value, got none")
	}

	if v != "Z" {
		t.Errorf("parse: Expected Z, got %s", v)
	}
}

func TestMoveBoxes(t *testing.T) {
	stacks, rules, err := parse(mockData)
	if err != nil {
		t.Errorf("parse: %s", err)
	}

	err = moveBoxes(&stacks, rules)
	if err != nil {
		t.Errorf("moveBoxes: %s", err)
	}

	if len(stacks[0]) != 1 {
		t.Errorf("moveBoxes: Expected 1 boxes in stack 1, got %d", len(stacks[0]))
	}

	if len(stacks[1]) != 1 {
		t.Errorf("moveBoxes: Expected 1 box in stack 2, got %d", len(stacks[1]))
	}

	if len(stacks[2]) != 4 {
		t.Errorf("moveBoxes: Expected 4 boxes in stack 3, got %d", len(stacks[2]))
	}

	v, ok := stacks[0].Pop()

	if !ok {
		t.Errorf("parse: Expected a value, got none")
	}

	if v != "C" {
		t.Errorf("parse: Expected C, got %s", v)
	}

	v, ok = stacks[1].Pop()

	if !ok {
		t.Errorf("parse: Expected a value, got none")
	}

	if v != "M" {
		t.Errorf("parse: Expected M, got %s", v)
	}

	v, ok = stacks[2].Pop()

	if !ok {
		t.Errorf("parse: Expected a value, got none")
	}

	if v != "Z" {
		t.Errorf("parse: Expected Z, got %s", v)
	}
}

func TestTopMessage(t *testing.T) {
	stacks, rules, err := parse(mockData)
	if err != nil {
		t.Errorf("parse: %s", err)
	}

	err = moveBoxes(&stacks, rules)
	if err != nil {
		t.Errorf("moveBoxes: %s", err)
	}

	msg := topMessage(stacks)
	if msg != "CMZ" {
		t.Errorf("topMessage: Expected C M Z, got %s", msg)
	}
}

func TestPart1(t *testing.T) {
	msg := part1()
	if msg != "BSDMQFLSP" {
		t.Errorf("part1: Expected BSDMQFLSP, got %s", msg)
	}
}

func TestMoveMultiple(t *testing.T) {
	var stackFrom, stackTo stack
	stackFrom.Push("A")
	stackFrom.Push("B")

	err := moveMultiple(2, &stackFrom, &stackTo)

	if err != nil {
		t.Errorf("moveMultiple: %s", err)
	}

	if len(stackFrom) != 0 {
		t.Errorf("moveMultiple: Expected 0 boxes in stackFrom, got %d", len(stackFrom))
	}

	if len(stackTo) != 2 {
		t.Errorf("moveMultiple: Expected 2 boxes in stackTo, got %d", len(stackTo))
	}

	v, ok := stackTo.Pop()

	if !ok {
		t.Errorf("moveMultiple: Expected a value, got none")
	}

	if v != "B" {
		t.Errorf("moveMultiple: Expected B, got %s", v)
	}

	v, ok = stackTo.Pop()

	if !ok {
		t.Errorf("moveMultiple: Expected a value, got none")
	}

	if v != "A" {
		t.Errorf("moveMultiple: Expected A, got %s", v)
	}

}

func TestMoveBoxesWithMultiples(t *testing.T) {
	stacks, rules, err := parse(mockData)
	if err != nil {
		t.Errorf("parse: %s", err)
	}

	err = moveBoxesWithMultiples(&stacks, rules)

	if err != nil {
		t.Errorf("moveBoxesWithMultiples: %s", err)
	}

	if len(stacks[0]) != 1 {
		t.Errorf("moveBoxesWithMultiples: Expected 1 boxes in stack 1, got %d", len(stacks[0]))
	}

	if len(stacks[1]) != 1 {
		t.Errorf("moveBoxesWithMultiples: Expected 1 box in stack 2, got %d", len(stacks[1]))
	}

	if len(stacks[2]) != 4 {
		t.Errorf("moveBoxesWithMultiples: Expected 4 boxes in stack 3, got %d", len(stacks[2]))
	}

	v, ok := stacks[2].Pop()

	if !ok {
		t.Errorf("parse: Expected a value, got none")
	}

	if v != "D" {
		t.Errorf("parse: Expected D, got %s", v)
	}

	v, ok = stacks[2].Pop()

	if !ok {
		t.Errorf("parse: Expected a value, got none")
	}

	if v != "N" {
		t.Errorf("parse: Expected N, got %s", v)
	}
}

func TestTopMessageWithMultiples(t *testing.T) {
	stacks, rules, err := parse(mockData)
	if err != nil {
		t.Errorf("parse: %s", err)
	}

	err = moveBoxesWithMultiples(&stacks, rules)
	if err != nil {
		t.Errorf("moveBoxesWithMultiples: %s", err)
	}

	msg := topMessage(stacks)
	if msg != "MCD" {
		t.Errorf("topMessage: Expected D N Z, got %s", msg)
	}
}

func TestPart2(t *testing.T) {
	msg := part2()
	if msg != "PGSQBFLDP" {
		t.Errorf("part2: Expected BSDMQFLSP, got %s", msg)
	}
}
