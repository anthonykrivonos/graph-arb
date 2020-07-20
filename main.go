package main

import (
	"fmt"
	"github.com/anthonykrivonos/graph-arb/algos"
	"github.com/anthonykrivonos/graph-arb/client"
	"github.com/anthonykrivonos/graph-arb/graph"
	"github.com/anthonykrivonos/graph-arb/math"
	"github.com/anthonykrivonos/graph-arb/models"
)

func main() {
	c := client.NewGeminiClient()

	// Securities to watch
	btc := *models.NewSecurity("btc")
	eth := *models.NewSecurity("eth")
	usd := *models.NewSecurity("usd")

	g := graph.Graph{}
	btcNode := g.Add(btc)
	g.Add(eth)
	g.Add(usd)

	onNewPrice := func(price float64, symbol models.Symbol) {
		// Update the graph
		fromNode := graph.NewNode(symbol.FromSecurity())
		toNode := graph.NewNode(symbol.ToSecurity())
		g.AddWeightedEdge(*fromNode, *toNode, math.NegLog(price))

		// Get a list of arb paths
		paths := algos.Bellman(&g, btcNode)
		for _, path := range paths {
			fmt.Println(path.String())
		}
	}

	c.Watch(onNewPrice, *models.NewSymbol(btc, usd), *models.NewSymbol(eth, usd))
}
