package client

import (
	"fmt"
	"github.com/adshao/go-binance"
	"github.com/anthonykrivonos/graph-arb/models"
	"strconv"
	"sync"
)

type binanceClient struct {
	client
}

func (c *binanceClient) Watch(callback func(float64, models.Symbol), symbols ...models.Symbol) {
	var wg sync.WaitGroup
	for _, symbol := range symbols {
		wg.Add(1)
		go c.watch(symbol, callback)
	}
	wg.Wait()
}

func (c *binanceClient) watch(symbol models.Symbol, callback func(float64, models.Symbol)) {
	fmt.Printf("Registered watcher for %s\n", symbol.String())
	wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
		if price, err := strconv.ParseFloat(event.Price, 32); err == nil {
			callback(price, symbol)
		}
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsAggTradeServe(symbol.String(), wsAggTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func NewBinanceClient() Client {
	gc := binanceClient{}
	return &gc
}

var _ Client = &binanceClient{}
