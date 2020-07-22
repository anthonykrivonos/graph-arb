package algos

import (
	"github.com/anthonykrivonos/graph-arb/graph"
	"math"
)

func Bellman(g *graph.Graph, source *graph.Node) []*graph.Path {
	// List of distances from the source node
	dists := make(map[*graph.Node] float64)

	// List of previous nodes in the path to the current node
	prevs := make(map[*graph.Node] *graph.Node)

	// Initialize graph
	for _, node := range g.Nodes() {
		dists[node] = math.MaxFloat64
		prevs[node] = nil
	}

	// Set distance from source to zero
	dists[source] = 0.0

	// Repeatedly relax edges
	for _, u := range g.Nodes() {
		for _, e := range g.Neighbors(u) {
			v := e.To()
			w := e.Weight()
			if dists[u] + w < dists[v] {
				dists[v] = dists[u] + w
				prevs[v] = u
			}
		}
	}

	// Check for negative weight cycles
	paths := make([]*graph.Path, 0)
	for _, u := range g.Nodes() {
		for _, e := range g.Neighbors(u) {
			v := e.To()
			w := e.Weight()
			if dists[u] + w < dists[v] {
				x := prevs[u]
				nodesUsed := make(map[*graph.Node] bool)
				var nodes []*graph.Node
				nodes = append(nodes, u)
				for {
					if _, ok := nodesUsed[x]; ok {
						break
					}
					nodesUsed[x] = true
					nodes = append([]*graph.Node{x}, nodes...)
					x = prevs[x]
				}
				path := graph.NewPath(nodes, dists[u])
				paths = append(paths, path)
			}
		}
	}

	return paths
}
