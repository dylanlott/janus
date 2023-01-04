package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestNodeBuilder(t *testing.T) {
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
}
