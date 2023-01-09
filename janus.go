package main

// Janus is a graph data structure implementation in Go.
// It builds out a list of nodes in their order of creation and records
// relationship between two nodes as an edge in the graph.
// Edges are represented as triplet of a start, an end, and a weight to the relationship.

// Graph is our new concrete implementation.
type Graph struct {
	nodes []*GraphNode[any]
}

// Edge is a triplet of start node, end node, and the edge weight.
type Edge [3]int64

// New returns a new Graph.
func New() *Graph {
	return &Graph{
		nodes: []*GraphNode[any]{},
	}
}

// AddNode adds a node to the end of the list with an empty set of edges.
func (n *Graph) AddNode() int64 {
	id := len(n.nodes)
	n.nodes = append(n.nodes, &GraphNode[any]{
		id:    int64(id),
		edges: map[int64]int64{},
		val:   nil,
	})
	return int64(id)
}

// AddEdge adds a start to end edge and gives it a weight of w
func (n *Graph) AddEdge(start int64, end int64, weight int64) {
	n.nodes[start].edges[end] = weight
}

// Neighbors returns a list of all node IDs that share an edge with the node.
func (n *Graph) Neighbors(id int64) []int64 {
	neighbors := []int64{}
	for _, node := range n.nodes {
		// for each edge in the node's list, record its weight.
		for edge := range node.edges {
			if node.id == id {
				neighbors = append(neighbors, edge)
			}
			if edge == id {
				neighbors = append(neighbors, node.id)
			}
		}
	}
	return neighbors
}

// Nodes returns a list of all nodes in the graph.
func (n *Graph) Nodes() []int64 {
	list := make([]int64, 0, int64(len(n.nodes)))
	for _, n := range n.nodes {
		list = append(list, n.id)
	}
	return list
}

// Edges returns the list of edges in the graph.
func (n *Graph) Edges() []Edge {
	// create a slice of edges as large as the list of nodes we have
	edges := make([]Edge, 0, int64(len(n.nodes)))
	// iterate over all nodes and collect their edges
	for nodeID := range n.nodes {
		for peer, weight := range n.nodes[nodeID].edges {
			// loops over the edges in each node in the ID list and
			// creates an Edge entry for each.
			edges = append(edges, Edge{int64(nodeID), peer, weight})
		}
	}

	return edges
}

// Remove deletes a node and all of its edges.
func (n *Graph) Remove(id int64) {
	for idx, node := range n.nodes {
		if node.id == id {
			// slice out of the nodes index
			n.nodes = append(n.nodes[:idx], n.nodes[idx+1:]...)
		}
		if _, ok := node.edges[id]; ok {
			// remove any edges that were connected to that node
			delete(node.edges, id)
		}
	}
}

// Predicate is a function type that takes a node and returns true
// if the node passes the predicate constraints.
type Predicate func(node *GraphNode[any]) bool

// Filter applies a predicate to each node in the Iterator.
// It returns the slice of nodes that passed the predicate.
func Filter(i Iterator, pred Predicate) []*GraphNode[any] {
	filtered := []*GraphNode[any]{}

	// test predicate against every node in the iterator
	for i.HasNext() {
		node := i.Next()
		if ok := pred(node); ok {
			// append to filtered if it passes
			filtered = append(filtered, node)
		}
	}

	return filtered
}
