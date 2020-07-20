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
	"strings"
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

func (c *geminiClient) Watch(securities ...models.Security) {
	// Create a list of symbols to watch
	var symbols []string
	for i, secA := range securities {
		for j, secB := range securities {
			if i != j {
				symbol := strings.ToLower(fmt.Sprintf("%s%s", secA.Name(), secB.Name()))
				c.watch(fmt.Sprintf("/%s?bids=true&offers=true", symbol))
				symbols = append(symbols, symbol)
			}
		}
	}
}

func (c *geminiClient) watch(endpoint string) {
	// Create interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Create address to connect to
	u := url.URL{Scheme: "wss", Host: c.watchHost, Path: fmt.Sprint(c.watchPath, endpoint)}
	fmt.Printf("Connecting to %s\n", u.String())

	// Dial into the socket
	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial: ", err, resp.StatusCode, "\n")
	}
	defer conn.Close()

	// Receive from the socket
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			var res geminiClientResponse
			json.Unmarshal(message, &res)

			if len(res.Events) > 0 {
				log.Printf("price: %s\n", res.Events[0].Price)
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
