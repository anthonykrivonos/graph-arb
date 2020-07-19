package graph

import "fmt"

type Edge struct {
	weight float64
	from *Node
	to *Node
}

func (e *Edge) String() string {
	return e.from.String() + " => " + e.to.String() + fmt.Sprintf(" (%f)", e.weight)
}

func NewWeightedEdge(from *Node, to *Node, weight float64) *Edge {
	e := new(Edge)
	e.weight = weight
	e.from = from
	e.to = to
	return e
}
