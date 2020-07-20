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

func (e *Edge) From() *Node {
	return e.from
}

func (e *Edge) To() *Node {
	return e.to
}

func (e *Edge) Weight() float64 {
	return e.weight
}

func NewEdge(from *Node, to *Node, weight float64) *Edge {
	e := new(Edge)
	e.weight = weight
	e.from = from
	e.to = to
	return e
}
