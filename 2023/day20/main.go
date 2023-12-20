package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strings"
)

type message struct {
	from, to nodeInterface
	high     bool
}

type nodeInterface interface {
	AddNext(nodeInterface)
	AddPrev(nodeInterface)
	Status() string
	PreviousHigh() bool
	Reply(nodeInterface, bool) []message
	Name() string
}

type nodesMap map[string]nodeInterface

// node is a node in the circuit
type node struct {
	next, prev   []nodeInterface
	name         string
	previousHigh bool // true if previous signal was high
}

// AddNext adds a next node
func (n *node) AddNext(next nodeInterface) {
	n.next = append(n.next, next)
}

// AddPrev adds a previous node
func (n *node) AddPrev(prev nodeInterface) {
	n.prev = append(n.prev, prev)
}

// PreviousHigh returns true if the previous signal was high
func (n *node) PreviousHigh() bool {
	return n.previousHigh
}

func (n *node) Name() string {
	return n.name
}

var _ nodeInterface = (*flip)(nil)

// flip is a flip-flop node
type flip struct {
	node
	on bool
}

// Status returns the status of the flip-flop
func (f *flip) Status() string {
	if f.on {
		return "1"
	}
	return "0"
}

// Reply returns the message(s) to send to the next node(s)
func (f *flip) Reply(_ nodeInterface, high bool) []message {
	f.previousHigh = high

	if high {
		return []message{}
	}

	nextSignalHigh := true

	// Switch the flip-flop
	f.on = !f.on

	// If node is off, next signal is low
	if !f.on {
		nextSignalHigh = false
	}

	m := make([]message, len(f.next))
	for i, next := range f.next {
		m[i] = message{
			from: f,
			to:   next,
			high: nextSignalHigh,
		}
	}
	return m
}

var _ nodeInterface = (*conjunction)(nil)

// conjunction is a conjunction node
type conjunction struct {
	node
	memory map[string]bool
}

// Status returns the status of the conjunction
func (c *conjunction) Status() string {
	s := ""

	for _, p := range c.prev {
		high, ok := c.memory[p.Name()]
		if !ok {
			s += "0"
			continue
		}
		if high {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}

// Reply returns the message(s) to send to the next node(s)
func (c *conjunction) Reply(from nodeInterface, high bool) []message {
	c.previousHigh = high
	c.memory[from.Name()] = high

	nextSignalHigh := false
	if strings.Contains(c.Status(), "0") {
		nextSignalHigh = true
	}

	m := make([]message, len(c.next))

	for i, next := range c.next {
		m[i] = message{
			from: c,
			to:   next,
			high: nextSignalHigh,
		}
	}

	return m
}

var _ nodeInterface = (*broadcast)(nil)

// broadcast is a broadcast node
type broadcast struct {
	node
}

// Status returns the status of the broadcast
func (b *broadcast) Status() string {
	return ""
}

// Reply returns the message(s) to send to the next node(s)
func (b *broadcast) Reply(_ nodeInterface, high bool) []message {
	m := make([]message, len(b.next))
	b.previousHigh = high

	for i, next := range b.next {
		m[i] = message{
			from: b,
			to:   next,
			high: high,
		}
	}
	return m
}

var _ nodeInterface = (*end)(nil)

type end struct {
	node
}

func (e *end) Status() string {
	return ""
}

func (e *end) Reply(_ nodeInterface, _ bool) []message {
	return []message{}
}

func parse(lines []string) (nodesMap, error) {

	var nodes = make(nodesMap)
	var connections = make(map[string][]string)

	for _, line := range lines {

		if len(line) == 0 {
			continue
		}

		// Remove all spaces from the line
		line = strings.ReplaceAll(line, " ", "")

		// Split the line into two parts
		parts := strings.Split(line, "->")

		if len(parts) != 2 {
			return nil, fmt.Errorf("parse - invalid line: %s", line)
		}

		name := parts[0]
		movingTo := strings.Split(parts[1], ",")

		// The first part is the name of the node
		switch name[0] {
		case 'b':
			nodes[name] = &broadcast{node: node{name: name}}
		case '%':
			name = name[1:]
			nodes[name] = &flip{node: node{name: name}}
		case '&':
			name = name[1:]
			nodes[name] = &conjunction{node: node{name: name}, memory: make(map[string]bool)}
		case 'e':
			nodes[name] = &end{node: node{name: name}}
		}

		// The second part is the name of the node(s) it is connected to
		connections[name] = movingTo
	}

	// Populate the nodes

	for from, to := range connections {
		for _, next := range to {
			_, ok := nodes[next]
			if !ok {
				nodes[next] = &end{node: node{name: next}}
			}
			nodes[from].AddNext(nodes[next])
			nodes[next].AddPrev(nodes[from])
		}
	}

	return nodes, nil
}

// nPressUntil returns the number of press until a node receives a given signal
func nPressUntil(nodes nodesMap, to nodeInterface, high bool) int {
	var count int

	for {
		var queue []message
		m := nodes["broadcaster"].Reply(nil, false)
		queue = append(queue, m...)
		count++

		for len(queue) > 0 {
			m := queue[0]
			queue = queue[1:]

			if m.to.Name() == to.Name() && m.high == high {
				return count
			}

			queue = append(queue, m.to.Reply(m.from, m.high)...)
		}
	}
}

// pressButton presses the button on the broadcaster node
func pressButton(nodes nodesMap) (int, int) {
	nHigh, nLow := 0, 1
	var queue []message

	// Add the first message to the queue
	m := nodes["broadcaster"].Reply(nil, false)

	queue = append(queue, m...)

	for len(queue) > 0 {
		// Get the first message from the queue
		m := queue[0]

		if m.high {
			nHigh++
		} else {
			nLow++
		}

		queue = queue[1:]

		// Send the message
		queue = append(queue, m.to.Reply(m.from, m.high)...)

	}

	return nHigh, nLow
}

// pressButtonN presses the button on the broadcaster node n times
func pressButtonN(nodes nodesMap, n int) (int, int) {
	var nHigh, nLow int
	for i := 0; i < n; i++ {
		h, l := pressButton(nodes)
		nHigh += h
		nLow += l
	}

	return nHigh, nLow
}

// productHighLows returns the product of the number of high and low signals
func productHighLows(nodes nodesMap, n int) int {
	nHigh, nLow := pressButtonN(nodes, n)
	return nHigh * nLow
}

func part1() int {
	fmt.Println("Part 1:")
	n, err := parse(utils.ReadFile("input.txt"))
	if err != nil {
		panic(err)
	}

	prod := productHighLows(n, 1000)
	fmt.Printf("The product of the number of high and low signals is %d\n", prod)
	return prod
}

func part2() int {
	fmt.Println("Part 2:")
	n, err := parse(utils.ReadFile("input2.txt"))
	if err != nil {
		panic(err)
	}
	final := n["rx"]
	var numbers []int
	previous := final.(*end).prev[0].(*conjunction).prev

	for p := range previous {
		n, err := parse(utils.ReadFile("input2.txt"))
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, nPressUntil(n, previous[p], false))
	}

	commonCycle := utils.LCM(numbers...)

	fmt.Printf("First time after pressing %d times\n", commonCycle)
	return commonCycle
}

func main() {
	part1()
	part2()
}
