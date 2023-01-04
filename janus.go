package main

// Node is an abstract interface for implementing a graph data structure.
type Node interface{}

var _ Node = (*GraphNode)(nil)

// Graph is our new concrete implementation.
type Graph struct {
	nodes []*GraphNode
}

// Edge is a triplet of start node, end node, and the edge weight.
type Edge [3]int64

// GraphNode is a node in our NodeBuilder
type GraphNode struct {
	id    int64
	edges map[int64]int64 // a connection to another node and the connection's weight
}

// New returns a new Graph.
func New() *Graph {
	return &Graph{
		nodes: []*GraphNode{},
	}
}

// AddNode adds a node to the end of the list with an empty set of edges.
func (n *Graph) AddNode() int64 {
	id := len(n.nodes)
	n.nodes = append(n.nodes, &GraphNode{
		id:    int64(id),
		edges: map[int64]int64{},
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
