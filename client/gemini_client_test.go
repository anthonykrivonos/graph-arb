package client

import (
	"github.com/anthonykrivonos/graph-arb/models"
	"testing"
)

func TestGeminiClient(t *testing.T) {
	c := NewGeminiClient()

	// Securities to watch
	btc := *models.NewSecurity("btc")
	usd := *models.NewSecurity("usd")

	c.Watch(btc, usd)
}
