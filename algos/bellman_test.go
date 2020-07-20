package algos

import (
	"fmt"
	"github.com/anthonykrivonos/graph-arb/graph"
	"github.com/anthonykrivonos/graph-arb/math"
	"github.com/anthonykrivonos/graph-arb/models"
	"testing"
)

func TestBellman(t *testing.T) {
	g := new(graph.Graph)

	// Create securities
	btc := models.NewSecurity("btc")
	eth := models.NewSecurity("eth")
	bch := models.NewSecurity("bch")
	usd := models.NewSecurity("usd")

	// Add securities to the graph
	btcNode := g.Add(btc)
	ethNode := g.Add(eth)
	bchNode := g.Add(bch)
	usdNode := g.Add(usd)

	// Create arbitrage opportunities
	g.AddWeightedEdge(*btcNode, *ethNode, math.NegLog(50))
	g.AddWeightedEdge(*ethNode, *bchNode, math.NegLog(40))
	g.AddWeightedEdge(*bchNode, *usdNode, math.NegLog(50))
	g.AddWeightedEdge(*bchNode, *ethNode, math.NegLog(150))
	g.AddWeightedEdge(*usdNode, *btcNode, math.NegLog(100))

	// Get arbitrage paths
	arbPaths := Bellman(g, btcNode)
	for _, path := range arbPaths {
		fmt.Println(path.String())
	}
}
