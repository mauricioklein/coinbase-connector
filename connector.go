package coinbase

import (
	"net/http"

	"github.com/imroc/req"
)

// Connector defines the interface between the library
// and the Coinbase API
type Connector struct {
	req         *req.Req
	credentials Credentials
}

// Credentials define the Coinbase credentials
// used to authenticate in the API
type Credentials struct {
	URI        string
	Key        string
	Secret     string
	Passphrase string
}

// New instantiates a new connector
func New(cred Credentials) *Connector {
	return &Connector{
		req:         req.New(),
		credentials: cred,
	}
}

// Client returns the underlying HTTP Client
func (c *Connector) Client() *http.Client {
	return c.req.Client()
}
