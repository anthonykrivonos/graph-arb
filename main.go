package main

import (
	"fmt"
	"github.com/anthonykrivonos/graph-arb/algos"
	"github.com/anthonykrivonos/graph-arb/client"
	"github.com/anthonykrivonos/graph-arb/constants"
	"github.com/anthonykrivonos/graph-arb/graph"
	"github.com/anthonykrivonos/graph-arb/math"
	"github.com/anthonykrivonos/graph-arb/models"
)


func main() {
	c := client.NewBinanceClient()

	securities := SECURITIES[:30]

	g := &graph.Graph{}
	secList := make([]*models.Security, 0)
	symbols := make([]models.Symbol, 0)
	nodeMap := map[models.Security]*graph.Node{}

	// Create securities and nodes in the graph
	for _, secName := range securities {
		sec := models.NewSecurity(secName)
		secList = append(secList, sec)
		secNode := g.Add(sec)
		nodeMap[*sec] = secNode
	}

	// Make a list of symbols
	for i := 0; i < len(securities); i++ {
		for j := 0; j < len(securities) && j != i; j++ {
			symbolIJ := models.NewSymbol(*secList[i], *secList[j])
			symbolJI := models.NewSymbol(*secList[j], *secList[i])
			symbols = append(symbols, *symbolIJ, *symbolJI)
		}
	}

	onNewPrice := func(price float64, symbol models.Symbol) {
		// Update the graph
		fromNode := nodeMap[symbol.FromSecurity()]
		toNode := nodeMap[symbol.ToSecurity()]
		g.AddWeightedEdge(fromNode, toNode, math.NegLog(price))

		// Get a list of arb paths
		paths := algos.Bellman(g)
		if len(paths) > 0 {
			for _, path := range paths {
				fmt.Println(path.String())
			}
		} else {
			//fmt.Printf("%s: No arb opportunities found\n", symbol.String())
		}
	}

	c.Watch(onNewPrice, symbols...)
}

var (
	SECURITIES = []string{
		constants.BTC,
		constants.ETH,
		constants.USD,
		constants.XRP,
		constants.LTC,
		constants.BNB,
		constants.ADA,
		constants.BTT,
		constants.ERD,
		constants.XTZ,
		constants.BUS,
		constants.AEX,
		constants.AGI,
		constants.AIO,
		constants.ALG,
		constants.ANK,
		constants.ARP,
		constants.ATO,
		constants.BND,
		constants.BAT,
		constants.BCP,
		constants.BEA,
		constants.BGB,
		constants.BKR,
		constants.BLZ,
		constants.BNT,
		constants.BQX,
		constants.BRD,
		constants.BTS,
		constants.CDT,
		constants.CEL,
		constants.CHR,
		constants.CHZ,
		constants.CMT,
		constants.CND,
		constants.COC,
		constants.COM,
		constants.COS,
		constants.COT,
		constants.CTS,
		constants.CVC,
		constants.DSH,
		constants.DAT,
		constants.DCR,
		constants.DET,
		constants.DGB,
		constants.DLT,
		constants.DNT,
		constants.DCK,
		constants.DOG,
		constants.DRE,
		constants.DSK,
		constants.ELF,
		constants.ENG,
		constants.ENJ,
		constants.EOS,
		constants.ETC,
		constants.EVX,
		constants.FET,
		constants.FTM,
		constants.FTT,
		constants.FUN,
		constants.GAS,
		constants.GDP,
		constants.GNT,
		constants.GRS,
		constants.GTO,
		constants.GVT,
		constants.GXS,
		constants.HBR,
		constants.HCX,
		constants.HVE,
		constants.HOT,
		constants.ICX,
		constants.INS,
		constants.IOS,
		constants.IOT,
		constants.IOX,
		constants.IQX,
		constants.IRI,
		constants.KVA,
		constants.KEY,
		constants.KMD,
		constants.KNC,
		constants.LND,
		constants.LNK,
		constants.LRC,
		constants.LSK,
		constants.LTO,
		constants.MAN,
		constants.MAT,
		constants.MBL,
		constants.MCO,
		constants.MDA,
		constants.MDT,
		constants.MFT,
		constants.MTH,
		constants.MTL,
		constants.NAN,
		constants.NCS,
		constants.NEB,
		constants.NEO,
		constants.NKN,
		constants.NXS,
		constants.NLS,
		constants.OAX,
		constants.OGN,
		constants.OMG,
		constants.ONE,
		constants.ONG,
		constants.ONT,
		constants.OST,
		constants.PAX,
		constants.PER,
		constants.PHB,
		constants.PIV,
		constants.PNT,
		constants.POA,
		constants.POE,
		constants.POL,
		constants.POW,
		constants.PPT,
		constants.QKC,
		constants.QLC,
		constants.QSP,
		constants.QTM,
		constants.RCN,
		constants.RDN,
		constants.REN,
		constants.REP,
		constants.REQ,
		constants.RLC,
		constants.RVN,
		constants.SCX,
		constants.SKY,
		constants.SNG,
		constants.SNM,
		constants.SNT,
		constants.SNX,
		constants.SOL,
		constants.STM,
		constants.SXX,
		constants.STR,
		constants.STP,
		constants.SRA,
		constants.STX,
		constants.SXP,
		constants.SYS,
		constants.TCT,
		constants.TFU,
		constants.THE,
		constants.TNB,
		constants.TNT,
		constants.TOM,
		constants.TRY,
		constants.TRX,
		constants.USC,
		constants.VET,
		constants.VIA,
		constants.VIB,
		constants.VTE,
		constants.VTH,
		constants.WAB,
		constants.WAN,
		constants.WAV,
		constants.WIN,
		constants.WPR,
		constants.WRX,
		constants.WTX,
		constants.XEM,
		constants.XLM,
		constants.XMR,
		constants.XVG,
		constants.XZC,
		constants.ZEC,
		constants.ZEN,
		constants.ZIL,
		constants.ZRX,
	}
)
