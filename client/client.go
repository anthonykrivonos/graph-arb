package client

import (
	"github.com/anthonykrivonos/graph-arb/models"
)

type Client interface {
	Watch(...models.Security)
}

type client struct {
	watchHost string
	watchPath string
}
