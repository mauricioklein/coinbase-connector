package coinbase

import (
	"github.com/imroc/req"
)

// Connector defines the interface between the library
// and the Coinbase API
type Connector struct {
	client      *req.Req
	credentials *Credentials
}

// Credentials define the Coinbase credentials
// used to authenticate in the API
type Credentials struct {
	URI        string
	Key        string
	Secret     string
	Passphrase string
}

// NewConnector instantiates a new connector
func NewConnector(client *req.Req, cred *Credentials) *Connector {
	return &Connector{
		client:      client,
		credentials: cred,
	}
}
