package main

// # DOMAIN MODELING
//
// ## Entity models
//
// Domain modeling starts by first identifying Entities.
// Here we define some Entities, and some relationships between them.

// Location is meant to be a physical location.
type Location struct{}

// Airport defines the interface for an Airport.
// * Airports can have many planes
// * Tickets have an origin and destination Airport.
type Airport struct{}

// Hub represents ground transportation hubs
type Hub struct{}

// Ticket has one Person and one Edge, representing the Trip.
type Ticket struct{}

// Plane has a location, which can be many things, including an airport or a GPS location.
type Plane struct{}

// Person can have many tickets.
type Person struct{}

// We have now defined some rough entity relationships for our domain.
//
// Next, we need to define the lifecycle and possible states of each of these items in our system.
//
// Planes can be under maintenance, or reserved for specific dates.
// Airports can be removed from a route for extreme weather.
// Tickets can be cancelled or changed.
// Routes can be expressed as two different Locations: an origin and destination.
// Tickets can be expressed as the same but for a Person.
//
// Next we build our underlying data structure without the interface depending on it,
// but prepared with a better understanding how all of our entities work together.
// The key problem we're solving is really a graph problem.
// With that in mind, we can craft our underlying data structure without exposing that complexity.

// Graph represents the graph of nodes that Janus uses to map out routes.
type Graph struct {
	Nodes []Node
}

// Node is an abstract interface for implementing a graph data structure.
type Node interface{}

var _ Node = (*GraphNode)(nil)

// Builder defines an interface for building and interacting with a graph representation of a set of locations.
type Builder interface {
	Build() (*Graph, error)
	Add(l Location) (*Graph, error)       // adds a node
	Remove(l Location) (*Graph, error)    // removes a node
	Update(l Location) (*Graph, error)    // updates a node's values
	Link(l1, l2 Location) (*Graph, error) // links two values creating an edge in the graph.
}

// Let's try a second version of this same core resource.
// We should always try multiple designs!
// Let's consider more abstract names and functions, decoupling from our Location terminology.

// NodeBuilder is our new concrete implementation.
type NodeBuilder struct {
	nodes []*GraphNode
}

// Edge is a triplet of start node, end node, and the edge weight.
type Edge [3]int64

// GraphNode is a node in our NodeBuilder
type GraphNode struct {
	id    int64
	edges map[int64]int64 // a connection to another node and the connection's weight
}

// New returns a new NodeBuilder
func New() *NodeBuilder {
	return &NodeBuilder{
		nodes: []*GraphNode{},
	}
}

// AddNode adds a node to the end of the list with an empty set of edges.
func (n *NodeBuilder) AddNode() int64 {
	id := len(n.nodes)
	n.nodes = append(n.nodes, &GraphNode{
		id:    int64(id),
		edges: map[int64]int64{},
	})
	return int64(id)
}

// AddEdge adds a start to end edge and gives it a weight of w
func (n *NodeBuilder) AddEdge(start int64, end int64, weight int64) {
	n.nodes[start].edges[end] = weight
}

// Neighbors returns a list of all node IDs that share an edge with the node.
func (n *NodeBuilder) Neighbors(id int64) []int64 {
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
func (n *NodeBuilder) Nodes() []int64 {
	// TODO: creates a list of all nodes in the graph and returns them
	panic("not impl")
}

// Edges returns the list of edges in the graph.
func (n *NodeBuilder) Edges() []Edge {
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

// And finally, let's define what our external interface should be.
// Remember that we don't want to let our implementation details influence our interface.
// We want our interface to reflect the best way to interact with our core problem.
// For example, functions like `AddNode` or `RemoveHub` are probably reflective of an abstraction leak.
// A Philosophy of Software Design reference here.
// With that in mind, let's define what our booking system should look like.
// This looks suspiciously like a CRUD interface, and that hints to me we've got a resource on our hands.

// BookingAgent exposes our passenger tickets service.
type BookingAgent interface {
	Book(tickets []Ticket) ([]Ticket, error)
	Cancel(tickets []Ticket) error
	Update(tickets []Ticket) ([]Ticket, error)
}

// To prove our abstraction is not leaky, let's test it against another adapter.
// What if we don't have just booking agents at our firm.
// We're an air transit company, people is one option, but cargo is the business to be in.
// So we have a shipping agent too, who works on the same underlying data structure, but looks for different things.

// ShippingTicket represents a shipping ticket in the shipping agent's system.
type ShippingTicket struct{}

// ShippingAgent defines a second system interacting with our same core resource - the graph of nodes.
type ShippingAgent interface {
	Ship(ticket Ticket) (*ShippingTicket, error)
	Find(id string) (*ShippingTicket, error)
	Update(ticket Ticket) (*ShippingTicket, error)
}
