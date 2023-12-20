package main

import (
	"testing"
)

var mockData = []string{
	"broadcaster -> a, b, c",
	"%a -> b",
	"%b -> c",
	"%c -> inv",
	"&inv -> a",
}

var mockData2 = []string{
	"broadcaster -> a",
	"%a -> inv, con",
	"&inv -> b",
	"%b -> con",
	"&con -> end",
	"end -> end",
}

func TestParse(t *testing.T) {
	n, err := parse(mockData)

	if err != nil {
		t.Errorf("parse - unexpected error: %s", err)
	}

	if len(n) != 5 {
		t.Errorf("parse - expected 5 nodes, got %d", len(n))
	}

	// Check if broadcaster has 3 next nodes
	if got := len(n["broadcaster"].(*broadcast).next); got != 3 {
		t.Errorf("parse - expected broadcaster to have 3 next nodes, got %d", got)
	}

}

func TestNode_Next(t *testing.T) {
	n := node{
		next: []nodeInterface{&flip{}, &flip{}},
	}

	if got := len(n.next); got != 2 {
		t.Errorf("node.Next - expected 2 next nodes, got %d", got)
	}

}

func TestNode_AddPrev(t *testing.T) {
	n := node{}
	n.AddPrev(&flip{})

	if got := len(n.prev); got != 1 {
		t.Errorf("node.AddPrev - expected 1 prev node, got %d", got)
	}
}

func TestNode_AddNext(t *testing.T) {
	n := node{}
	n.AddNext(&flip{})

	if got := len(n.next); got != 1 {
		t.Errorf("node.AddNext - expected 1 next node, got %d", got)
	}
}

func TestBroadcast_Status(t *testing.T) {
	n, err := parse(mockData)

	if err != nil {
		t.Errorf("parse - unexpected error: %s", err)
	}

	if got := n["broadcaster"].Status(); got != "" {
		t.Errorf("broadcast.Status - expected empty string, got %s", got)
	}
}

func TestNode_Name(t *testing.T) {
	n := node{name: "test"}

	if got := n.Name(); got != "test" {
		t.Errorf("node.Name - expected test, got %s", got)
	}
}

func TestFlip_Status(t *testing.T) {
	n, err := parse(mockData)

	if err != nil {
		t.Errorf("parse - unexpected error: %s", err)
	}

	if got := n["b"].Status(); got != "0" {
		t.Errorf("flip.Status - expected 0, got %s", got)
	}

	// Change status
	n["b"].(*flip).on = true

	if got := n["b"].Status(); got != "1" {
		t.Errorf("flip.Status - expected 1, got %s", got)
	}
}

func TestConjunction_Status(t *testing.T) {
	n, err := parse(mockData2)

	if err != nil {
		t.Errorf("parse - unexpected error: %s", err)
	}

	if got := n["con"].Status(); got != "00" {
		t.Errorf("conjunction.Status - expected '00', got %s", got)
	}

	// Send high signal to con
	n["con"].Reply(&flip{node: node{name: "a"}}, true)

	if got := n["con"].Status(); got != "10" && got != "01" {
		t.Errorf("conjunction.Status - expected 10, got %s", got)
	}

	// send low signal to con
	n["con"].Reply(&flip{node: node{name: "b"}}, false)

	if got := n["con"].Status(); got != "10" && got != "01" {
		t.Errorf("conjunction.Status - expected 10, got %s", got)
	}

	// send high signal to con
	n["con"].Reply(&flip{node: node{name: "b"}}, true)

	if got := n["con"].Status(); got != "11" {
		t.Errorf("conjunction.Status - expected 11, got %s", got)
	}
}

func TestFlip_ReplyHigh(t *testing.T) {
	n, err := parse(mockData)

	if err != nil {
		t.Errorf("parse - unexpected error: %s", err)
	}

	m := n["a"].Reply(nil, true)

	if got := len(m); got != 0 {
		t.Errorf("flip.Reply - expected 0 message, got %d", got)
	}
}

func TestFlip_ReplyLow(t *testing.T) {
	n, err := parse(mockData)

	if err != nil {
		t.Errorf("parse - unexpected error: %s", err)
	}
	// A is off
	m := n["a"].Reply(nil, false)

	// A is on
	if got := len(m); got != 1 {
		t.Errorf("flip.Reply - expected 1 message, got %d", got)
	}

	if got := m[0].to; got != n["b"] {
		t.Errorf("flip.Reply - expected b, got %s", got)
	}

	// If a became on when receiving a low signal -> out signal is high
	if got := m[0].high; got != true {
		t.Errorf("flip.Reply - expected true, got %t", got)
	}

	m = n["a"].Reply(nil, false)
	// A is off

	// If a became off when receiving a low signal -> out signal is low
	if got := m[0].high; got != false {
		t.Errorf("flip.Reply - expected false, got %t", got)
	}

}

func TestFlip_ReplyPreviousHigh(t *testing.T) {
	n, err := parse(mockData)

	if err != nil {
		t.Errorf("parse - unexpected error: %s", err)
	}

	if got := n["a"].PreviousHigh(); got != false {
		t.Errorf("flip.Reply - expected false, got %t", got)
	}

	n["a"].Reply(nil, true)

	if got := n["a"].PreviousHigh(); got != true {
		t.Errorf("flip.Reply - expected true, got %t", got)
	}

	// b Set to on
	n["b"].(*flip).on = true

	if got := n["b"].PreviousHigh(); got != false {
		t.Errorf("flip.Reply - expected false, got %t", got)
	}

	n["b"].Reply(nil, true)

	if got := n["b"].PreviousHigh(); got != true {
		t.Errorf("flip.Reply - expected true, got %t", got)
	}

}

func TestBroadcast_Reply(t *testing.T) {
	n, err := parse(mockData)

	testCases := []struct {
		name string
		high bool
		want bool
	}{
		{"High", true, true},
		{"Low", false, false},
	}
	if err != nil {
		t.Errorf("parse - unexpected error: %s", err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := n["broadcaster"].Reply(nil, tc.high)

			if got := len(m); got != 3 {
				t.Errorf("broadcast.Reply - expected 3 messages, got %d", got)
			}

			for _, msg := range m {
				if got := msg.to; got != n["a"] && got != n["b"] && got != n["c"] {
					t.Errorf("broadcast.Reply - expected a, b or c, got %s", got)
				}

				if got := msg.high; got != tc.want {
					t.Errorf("broadcast.Reply - expected %t, got %t", tc.want, got)
				}
			}

		})
	}
}

func TestBroadcast_ReplyPreviousHigh(t *testing.T) {
	n, err := parse(mockData2)

	if err != nil {
		t.Errorf("parse - unexpected error: %s", err)
	}

	if got := n["broadcaster"].PreviousHigh(); got != false {
		t.Errorf("conjunction.Reply - expected false, got %t", got)
	}

	n["broadcaster"].Reply(nil, true)

	if got := n["broadcaster"].PreviousHigh(); got != true {
		t.Errorf("conjunction.Reply - expected true, got %t", got)
	}

}

func TestConjunction_Reply(t *testing.T) {
	n, err := parse(mockData2)
	if err != nil {
		t.Errorf("parse - unexpected error: %s", err)
	}

	testMessage := func(messages []message, want bool) {
		for _, msg := range messages {
			if got := msg.high; got != want {
				t.Errorf("conjunction.Reply - expected %t, got %t", want, got)
			}
		}
	}

	// Send a high signal to con from A
	m := n["con"].Reply(&flip{node: node{name: "a"}}, true)
	testMessage(m, true)

	//Send a high signal to con from B
	m = n["con"].Reply(&flip{node: node{name: "b"}}, true)
	testMessage(m, false)

}

func TestConjunction_ReplyPreviousHigh(t *testing.T) {
	n, err := parse(mockData2)

	if err != nil {
		t.Errorf("parse - unexpected error: %s", err)
	}

	if got := n["con"].PreviousHigh(); got != false {
		t.Errorf("conjunction.Reply - expected false, got %t", got)
	}

	n["con"].Reply(&flip{node: node{name: "A"}}, true)

	if got := n["con"].PreviousHigh(); got != true {
		t.Errorf("conjunction.Reply - expected true, got %t", got)
	}
}

func TestPressButton(t *testing.T) {

	testCases := []struct {
		name     string
		data     []string
		wantHigh int
		wantLow  int
	}{
		{"MockData", mockData, 4, 8},
		{"MockData2", mockData2, 4, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			n, err := parse(tc.data)

			if err != nil {
				t.Errorf("parse - unexpected error: %s", err)
			}

			nHigh, nLow := pressButton(n)

			if nHigh != tc.wantHigh {
				t.Errorf("pressButton - expected %d high signals, got %d", tc.wantHigh, nHigh)
			}

			if nLow != tc.wantLow {
				t.Errorf("pressButton - expected %d low signals, got %d", tc.wantLow, nLow)
			}
		})
	}
}

func TestPressButtonN(t *testing.T) {

	testCases := []struct {
		name     string
		data     []string
		wantHigh int
		wantLow  int
	}{
		{"MockData", mockData, 4000, 8000},
		{"MockData2", mockData2, 2750, 4250},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			n, err := parse(tc.data)

			if err != nil {
				t.Errorf("parse - unexpected error: %s", err)
			}

			nHigh, nLow := pressButtonN(n, 1000)

			if nHigh != tc.wantHigh {
				t.Errorf("pressButton - expected %d high signals, got %d", tc.wantHigh, nHigh)
			}

			if nLow != tc.wantLow {
				t.Errorf("pressButton - expected %d low signals, got %d", tc.wantLow, nLow)
			}
		})
	}
}

func TestProductHighLows(t *testing.T) {
	testCases := []struct {
		name string
		data []string
		want int
	}{
		{"MockData", mockData, 32000000},
		{"MockData2", mockData2, 11687500},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			n, err := parse(tc.data)

			if err != nil {
				t.Errorf("parse - unexpected error: %s", err)
			}

			got := productHighLows(n, 1000)

			if got != tc.want {
				t.Errorf("productHighLows - expected %d, got %d", tc.want, got)
			}
		})
	}
}

func TestNPressUntil(t *testing.T) {

	testCases := []struct {
		name string
		high bool
		want int
	}{
		{"con", true, 1},
		{"inv", false, 2},
		{"b", true, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			n, err := parse(mockData2)

			if err != nil {
				t.Errorf("parse - unexpected error: %s", err)
			}

			got := nPressUntil(n, n[tc.name], tc.high)

			if got != tc.want {
				t.Errorf("iterationsUntilLoop - expected %d, got %d", tc.want, got)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	want := 806332748
	got := part1()

	if got != want {
		t.Errorf("part1 - expected %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 228060006554227
	got := part2()

	if got != want {
		t.Errorf("part2 - expected %d, got %d", want, got)
	}
}
