package client

import (
	"github.com/anthonykrivonos/graph-arb/models"
)

type Client interface {
	Watch(func(float64, models.Symbol), ...models.Symbol)
}

type client struct {
	watchHost string
	watchPath string
}
