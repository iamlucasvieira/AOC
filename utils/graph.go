package utils

import (
	"container/heap"
)

type NodeInterface interface {
	GetValue() int
	SetValue(int)
	Equal(n NodeInterface) bool
	Next() NodeInterface
	IsEmpty() bool
	String() string
}

// Node is a struct that represents a node in a graph.
type Node struct {
	Point
	Value int
	Prev  *Node
}

func (n *Node) GetValue() int {
	return n.Value
}

func (n *Node) SetValue(value int) {
	n.Value = value
}

func (n *Node) Equal(node NodeInterface) bool {
	// Type assertion to check if node is of type Node
	other, ok := node.(*Node)
	if !ok {
		return false
	}
	return n.Point.Equal(other.Point)
}

func (n *Node) Next() NodeInterface {
	return n.Prev
}

func (n *Node) IsEmpty() bool {
	return n == nil
}

func (n *Node) String() string {
	return n.Point.String()
}

// Neighbors is a function that returns the neighbors of a node in a graph.
func Neighbors(n *Node, g Graph[*Node]) []*Node {
	directions := []Point{
		{0, -1}, // up
		{0, 1},  // down
		{-1, 0}, // left
		{1, 0},  // right
	}

	neighbors := make([]*Node, 0)

	for _, direction := range directions {
		neighbor := n.Add(direction)
		contains := g.Nodes.Contains(neighbor)
		prev := n.Prev
		if prev != nil {
			contains = contains && !neighbor.Equal(prev.Point)
		}
		if contains {
			neighbour := g.Get(neighbor.X, neighbor.Y)
			neighbors = append(neighbors, &Node{
				Point: neighbour.Point,
				Value: n.Value + neighbour.Value,
				Prev:  n,
			})
		}
	}

	return neighbors
}

// NewNode is a function that returns a new Node.
func NewNode(x, y, value int) Node {
	return Node{Point{x, y}, value, nil}
}

// Graph is a struct that represents a graph.
type Graph[T NodeInterface] struct {
	Nodes Grid[T]
}

// Get is a method that returns a node in a graph.
func (g Graph[T]) Get(x, y int) T {
	return g.Nodes[y][x]
}

// NewGraph is a function that returns a new Graph.
func NewGraph(input Grid[int]) Graph[*Node] {
	nodes := make(Grid[*Node], input.Height())
	for y := range input {
		nodes[y] = make([]*Node, input.Width())
		for x := range input[y] {
			nodes[y][x] = &Node{Point{x, y}, input[y][x], nil}
		}
	}
	return Graph[*Node]{Nodes: nodes}
}

type NodeQueue[T NodeInterface] []T

func (nq NodeQueue[T]) Len() int           { return len(nq) }
func (nq NodeQueue[T]) Less(i, j int) bool { return nq[i].GetValue() < nq[j].GetValue() }
func (nq NodeQueue[T]) Swap(i, j int)      { nq[i], nq[j] = nq[j], nq[i] }

// Push is a method that pushes a node onto a NodeQueue.
func (nq *NodeQueue[T]) Push(x interface{}) {
	*nq = append(*nq, x.(T))
}

// Pop is a method that pops a node from a NodeQueue.
func (nq *NodeQueue[T]) Pop() interface{} {
	old := *nq
	n := len(old)
	x := old[n-1]
	*nq = old[0 : n-1]
	return x
}

// NewNodeQueue is a function that returns a new NodeQueue.
func NewNodeQueue[T NodeInterface]() *NodeQueue[T] {
	nq := make(NodeQueue[T], 0)
	return &nq
}

// DijkstraTemplate is the base function for Dijkstra's algorithm.
func DijkstraTemplate[T NodeInterface](g Graph[T], start T, end T, NeighborFunc func(T, Graph[T]) []T) []T {
	visited := make(map[string]bool)
	queue := NewNodeQueue[T]()

	start.SetValue(0)

	heap.Push(queue, start)

	for queue.Len() > 0 {
		current := heap.Pop(queue).(T)
		key := current.String()
		if visited[key] {
			continue
		}
		visited[key] = true

		if current.Equal(end) {
			return ReconstructPath[T](current)
		}

		for _, neighbor := range NeighborFunc(current, g) {
			if !visited[neighbor.String()] {
				heap.Push(queue, neighbor)
			}
		}
	}
	return nil
}

// Dijkstra is a function that returns the shortest path between two nodes.
func Dijkstra(g Graph[*Node], start *Node, end *Node) []*Node {
	return DijkstraTemplate[*Node](g, start, end, Neighbors)
}

// ReconstructPath is a function that reconstructs a path from a node.
func ReconstructPath[T NodeInterface](end T) []T {
	path := make([]T, 0)
	for n := end; !n.IsEmpty(); n = n.Next().(T) {
		path = append(path, n)
	}

	// Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// DijkstraShortestDistance is a function that returns the shortest distance between two nodes.
func DijkstraShortestDistance(g Graph[*Node], start *Node, end *Node) int {
	path := Dijkstra(g, start, end)
	if path == nil {
		return -1
	}
	return path[len(path)-1].Value
}
