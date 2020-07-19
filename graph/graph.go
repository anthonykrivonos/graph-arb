package graph

import (
	"sync"
)

type Graph struct {
	nodes []*Node
	edges map[Node][]*Edge
	lock sync.RWMutex
}

func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddWeightedEdge(n1 *Node, n2 *Node, weight float64, bidirectional bool) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if g.edges == nil {
		g.edges = make(map[Node][] *Edge)
	}

	g.edges[*n1] = append(g.edges[*n1], NewWeightedEdge(n1, n2, weight))

	if bidirectional {
		g.edges[*n2] = append(g.edges[*n2], NewWeightedEdge(n2, n1, weight))
	}
}

func (g *Graph) AddEdge(n1 *Node, n2 *Node, bidirectional bool) {
	g.AddWeightedEdge(n1, n2, 0.0, bidirectional)
}

func (g *Graph) String() (s string) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	s = ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + ":\n"
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += "  " + near[j].String() + "\n"
		}
	}
	return
}
