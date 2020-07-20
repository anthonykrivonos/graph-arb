package graph

import "fmt"

type Path struct {
	nodes []*Node
	start *Node
	end *Node
	weight float64
}

func (p *Path) String() (s string) {
	s += fmt.Sprintf("Path of len %d: ", len(p.nodes))
	for i := 0; i < len(p.nodes); i++ {
		s += p.nodes[i].String()
		if i < len(p.nodes) - 1 {
			s += " -> "
		} else {
			s += " "
		}
	}
	s += fmt.Sprintf("(%f)", p.weight)
	return
}

func NewPath(nodes []*Node, weight float64) *Path {
	p := new(Path)
	p.nodes = nodes
	p.start = nodes[0]
	p.end = nodes[len(nodes) - 1]
	p.weight = weight
	return p
}
