package graph

import (
	"sync"
)

type Graph struct {
	nodes []*Node
	edges map[*Node][]*Edge
	lock sync.RWMutex
}

func (g *Graph) Nodes() []*Node {
	return g.nodes
}

func (g *Graph) Edges() map[*Node][]*Edge {
	return g.edges
}

func (g *Graph) Add(v interface{}) *Node {
	g.lock.Lock()
	defer g.lock.Unlock()

	n := Node{v}
	g.nodes = append(g.nodes, &n)

	return &n
}

func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddWeightedEdge(n1 *Node, n2 *Node, weight float64) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if g.edges == nil {
		g.edges = make(map[*Node][] *Edge)
	}

	e := NewEdge(n1, n2, weight)

	// Overwrite if n1 is already connected to n2
	includes := false
	for i, edge := range g.edges[n1] {
		if edge.to.value == n2.value {
			includes = true
			g.edges[n1][i] = e
			break
		}
	}

	// Otherwise, add to the list of edges
	if !includes {
		g.edges[n1] = append(g.edges[n1], e)
	}
}

func (g *Graph) AddEdge(n1 *Node, n2 *Node) {
	g.AddWeightedEdge(n1, n2, 0.0)
}

func (g *Graph) AddWeightedBidirectionalEdge(n1 *Node, n2 *Node, weight float64) {
	g.AddWeightedEdge(n1, n2, weight)
	g.AddWeightedEdge(n2, n1, weight)
}

func (g *Graph) AddBidirectionalEdge(n1 *Node, n2 *Node) {
	g.AddWeightedBidirectionalEdge(n1, n2, 0.0)
}

func (g *Graph) Neighbors(n *Node) []*Edge {
	g.lock.RLock()
	defer g.lock.RUnlock()

	if _, ok := g.edges[n]; !ok {
		return make([]*Edge, 0)
	}
	return g.edges[n]
}

func (g *Graph) String() (s string) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + ":\n"
		near := g.edges[g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += "  " + near[j].String() + "\n"
		}
	}
	return
}
