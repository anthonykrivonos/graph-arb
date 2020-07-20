package client

import (
	"encoding/json"
	"fmt"
	"github.com/anthonykrivonos/graph-arb/config"
	"github.com/anthonykrivonos/graph-arb/models"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

type geminiClient struct {
	client
}

type geminiClientEvent struct {
	Type string
	Side string
	Price string
	Remaining string
	Delta float64
	Reason string
}

type geminiClientResponse struct {
	Type string
	EventId string
	Timestamp int64
	Events []geminiClientEvent
}

func (c *geminiClient) Watch(callback func(float64, models.Symbol), symbols ...models.Symbol) {
	var wg sync.WaitGroup
	for _, symbol := range symbols {
		wg.Add(1)
		go c.watch(symbol, callback)
	}
	wg.Wait()
}

func (c *geminiClient) watch(symbol models.Symbol, callback func(float64, models.Symbol)) {
	// Create interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Create address to connect to
	u := url.URL{Scheme: "wss", Host: c.watchHost, Path: fmt.Sprintf("%s/%s", c.watchPath, symbol.String())}
	fmt.Printf("Connecting to %s\n", u.String())

	// Dial into the socket
	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial: ", err, resp.StatusCode, symbol.String(), "\n")
	}
	defer conn.Close()

	// Receive from the socket
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			// Receive the message
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			var res geminiClientResponse
			json.Unmarshal(message, &res)

			// Callback with the ask price
			if len(res.Events) > 0 {
				if res.Events[0].Side == "ask" {
					if price, err := strconv.ParseFloat(res.Events[0].Price, 32); err == nil && callback != nil {
						callback(price, symbol)
					}
				}
			}
		}
	}()

	// Tick every second
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// Read messages
	for {
		select {
		case <-done:
			return
		}
	}
}

func NewGeminiClient() Client {
	conf := config.NewConfig()
	c := client{
		watchHost: conf.GeminiUrl(),
		watchPath: "/v1/marketdata",
	}
	gc := geminiClient{c}
	return &gc
}

var _ Client = &geminiClient{}
