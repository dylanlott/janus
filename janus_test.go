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

func TestIterator(t *testing.T) {
	is := is.New(t)
	g := testGraph(t)
	i := NewIterator(g)

	is.True(i.HasNext())

	first := i.Next()
	is.Equal(first.id, int64(0))

	is.True(i.HasNext())
	_ = i.Next()
	is.True(i.HasNext())
	_ = i.Next()
	is.True(i.HasNext())
	_ = i.Next()

	is.Equal(i.HasNext(), false)
}

func TestFilter(t *testing.T) {
	is := is.New(t)
	g := testGraph(t)
	i := NewIterator(g)

	targetNode := g.nodes[0]
	var edgeFn Predicate = func(n *GraphNode) bool {
		if _, ok := n.edges[targetNode.id]; ok {
			return true
		}
		return false
	}

	got := Filter(i, edgeFn)

	is.Equal(len(got), 1)
	is.Equal(got[0].id, int64(2))
}

func TestRemove(t *testing.T) {
	is := is.New(t)
	g := testGraph(t)

	target := g.nodes[0].id
	g.Remove(target)

	is.Equal(len(g.nodes), 3)

	for _, node := range g.nodes {
		if _, ok := node.edges[target]; ok {
			t.Errorf("edges should not contain target id %+v", target)
		}
	}
}

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
