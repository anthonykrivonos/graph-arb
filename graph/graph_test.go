package graph

import (
	"testing"
)

var g Graph

func fillGraph() {
	nA := Node{"A"}
	nB := Node{"B"}
	nC := Node{"C"}
	nD := Node{"D"}
	nE := Node{"E"}
	nF := Node{"F"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB, true)
	g.AddEdge(&nA, &nC, true)
	g.AddEdge(&nB, &nE, true)
	g.AddEdge(&nC, &nE, true)
	g.AddEdge(&nE, &nF, true)
	g.AddEdge(&nD, &nA, true)
}

func TestAdd(t *testing.T) {
	fillGraph()
	g.String()
}
