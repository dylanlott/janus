package main

import "sync"

// GraphNode is a node in our NodeBuilder
type GraphNode[node any] struct {
	sync.Mutex

	id    int64
	edges map[int64]int64 // a connection to another node and the connection's weight
	val   node
}

// Get returns the generic value assigned to this GraphNode
func (g *GraphNode[node]) Get() node {
	g.Lock()
	defer g.Unlock()

	return g.val
}

// Set assigns a generic value to this GraphNode
func (g *GraphNode[node]) Set(n node) {
	g.Lock()
	defer g.Unlock()
	g.val = n
}
