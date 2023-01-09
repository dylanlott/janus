package main

import (
	"testing"

	"github.com/matryer/is"
)

type MockNode struct {
	id      int64
	payload map[string]interface{}
}

func TestNode(t *testing.T) {
	is := is.New(t)
	g := &GraphNode[*MockNode]{
		edges: make(map[int64]int64, 0),
		val: &MockNode{
			id: 1,
		},
	}
	one := g.Get()
	is.Equal(one.id, int64(1))
}
