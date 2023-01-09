package main

// NewIterator returns an iterator loaded with the provided graph nodes for traversal.
func NewIterator(graph *Graph) Iterator {
	return &iter{
		nodes: graph.nodes,
		curr:  0,
	}
}

// Iterator returns an interface for fulfilling an iterator adapter pattern.
type Iterator interface {
	Next() *GraphNode[any]
	HasNext() bool
}

// iter fulfills Iterable
type iter struct {
	nodes []*GraphNode[any]
	curr  int64
}

// Next advances the iterator and returns the node at that position
func (i *iter) Next() *GraphNode[any] {
	node := i.nodes[i.curr]
	i.curr++
	return node
}

// HasNext returns true if the iterator has values left to read
func (i *iter) HasNext() bool {
	if i.curr < int64(len(i.nodes)) {
		return true
	}
	return false
}
