package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestGraph(t *testing.T) {
	is := is.New(t)

	n := New()

	firstNode := n.AddNode()
	secondNode := n.AddNode()
	thirdNode := n.AddNode()

	is.Equal(len(n.nodes), 3)
	is.True(firstNode == 0)
	is.True(secondNode == 1)
	is.True(thirdNode == 2)

	n.AddEdge(firstNode, secondNode, 100)
	n.AddEdge(secondNode, thirdNode, 100)
	n.AddEdge(thirdNode, firstNode, 50)

	neighbors := n.Neighbors(firstNode)
	is.Equal(len(neighbors), 2)
	is.Equal(neighbors[0], int64(1))
	is.Equal(neighbors[1], int64(2))

	edges := n.Edges()
	is.Equal(len(edges), 3)

	list := n.Nodes()
	is.Equal(len(list), 3)
}

// func TestIteration(t *testing.T) {
// is := is.New(t)
//
// g := testGraph(t)
//
// iter := g.Iter()
// }

func testGraph(t *testing.T) *Graph {
	n := New()

	firstNode := n.AddNode()  // 0
	secondNode := n.AddNode() // 1
	thirdNode := n.AddNode()  // 2
	fourthNode := n.AddNode() // 3

	n.AddEdge(firstNode, secondNode, 100) // 0 => 1
	n.AddEdge(secondNode, thirdNode, 100) // 1 => 2
	n.AddEdge(thirdNode, firstNode, 50)   // 2 => 0
	n.AddEdge(fourthNode, thirdNode, 100) // 3 => 2

	return n
}
