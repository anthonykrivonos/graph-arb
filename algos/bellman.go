package algos

import (
	"github.com/anthonykrivonos/graph-arb/graph"
	"math"
)

func Bellman(g *graph.Graph) []*graph.Path {
	bellmanInner := func(start *graph.Node, pathList []*graph.Path) []*graph.Path {
		// List of distances from the start node
		dists := make(map[*graph.Node] float64)

		// List of previous nodes in the path to the current node
		prevs := make(map[*graph.Node] *graph.Node)

		// Initialize graph
		for _, node := range g.Nodes() {
			dists[node] = math.MaxFloat64
			prevs[node] = nil
		}

		// Set distance from start to zero
		dists[start] = 0.0

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
		for _, u := range g.Nodes() {
			for _, e := range g.Neighbors(u) {
				v := e.To()
				w := e.Weight()
				if x := prevs[u]; dists[u]+w < dists[v] && x != nil {
					nodes := make([]*graph.Node, 0)
					nodes = append(nodes, u, v)
					nodesUsed := make(map[*graph.Node] bool)
					nodesUsed[u] = true
					for {
						if _, ok := nodesUsed[x]; ok || x == nil {
							break
						}
						nodesUsed[x] = true
						nodes = append([]*graph.Node{x}, nodes...)
						x = prevs[x]
					}
					if x != nil {
						path := graph.NewPath(nodes, dists[u])
						pathList = append(pathList, path)
					}
				}
			}
		}
		return pathList
	}

	paths := make([]*graph.Path, 0)
	for _, node := range g.Nodes() {
		paths = bellmanInner(node, paths)
	}

	return paths
}
