package main

import (
	"fmt"
	"github.com/anthonykrivonos/graph-arb/algos"
	"github.com/anthonykrivonos/graph-arb/client"
	"github.com/anthonykrivonos/graph-arb/graph"
	"github.com/anthonykrivonos/graph-arb/math"
	"github.com/anthonykrivonos/graph-arb/models"
)

const (
	BTC = "btc"
	ETH = "eth"
	USD = "usdt"
	XRP = "xrp"
	LTC = "ltc"
	BNB = "bnb"
	ADA = "ada"
	BTT = "btt"
	ERD = "erd"
	XTZ = "xtz"
)

func main() {
	c := client.NewBinanceClient()

	// Securities to watch
	btc := models.NewSecurity(BTC)
	eth := models.NewSecurity(ETH)
	usd := models.NewSecurity(USD)
	xrp := models.NewSecurity(XRP)
	ltc := models.NewSecurity(LTC)
	bnb := models.NewSecurity(BNB)
	ada := models.NewSecurity(ADA)
	btt := models.NewSecurity(BTT)
	erd := models.NewSecurity(ERD)
	xtz := models.NewSecurity(XTZ)

	g := &graph.Graph{}
	btcNode := g.Add(btc)
	ethNode := g.Add(eth)
	usdNode := g.Add(usd)
	xrpNode := g.Add(xrp)
	ltcNode := g.Add(ltc)
	bnbNode := g.Add(bnb)
	adaNode := g.Add(ada)
	bttNode := g.Add(btt)
	erdNode := g.Add(erd)
	xtzNode := g.Add(xtz)

	nodeMap := map[models.Security]*graph.Node{
		*btc: btcNode,
		*eth: ethNode,
		*usd: usdNode,
		*xrp: xrpNode,
		*ltc: ltcNode,
		*bnb: bnbNode,
		*ada: adaNode,
		*btt: bttNode,
		*erd: erdNode,
		*xtz: xtzNode,
	}

	onNewPrice := func(price float64, symbol models.Symbol) {
		// Update the graph
		fromNode := nodeMap[symbol.FromSecurity()]
		toNode := nodeMap[symbol.ToSecurity()]
		g.AddWeightedEdge(fromNode, toNode, math.NegLog(price))

		// Get a list of arb paths
		paths := algos.Bellman(g, btcNode)
		if len(paths) > 0 {
			for _, path := range paths {
				fmt.Println(path.String())
			}
		} else {
			fmt.Println("No arb opportunities found")
			//fmt.Print(g.String())
		}
	}

	securities := []*models.Security{btc, eth, usd, xrp, ltc, bnb, ada, btt, erd, xtz}
	symbols := make([]models.Symbol, 0)
	for i := 0; i < len(securities); i++ {
		for j := 0; j < len(securities) && j != i; j++ {
			symbolIJ := models.NewSymbol(*securities[i], *securities[j])
			symbolJI := models.NewSymbol(*securities[j], *securities[i])
			symbols = append(symbols, *symbolIJ, *symbolJI)
		}
	}

	fmt.Print(symbols)

	c.Watch(onNewPrice, symbols...)
}
