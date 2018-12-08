package main

import (
	"errors"
	"fmt"

	"github.com/imroc/req"
	"github.com/mauricioklein/coinbase-connector/types"
)

const (
	coinbaseURI = "https://api-public.sandbox.pro.coinbase.com"
)

// Connector defines the interface between the library
// and the Coinbase API
type Connector struct {
	client *req.Req
}

// NewConnector instantiates a new connector
func NewConnector(client *req.Req) *Connector {
	return &Connector{
		client: client,
	}
}

// Ticker returns the ticker for the given productID
func (c *Connector) Ticker(productID string) (*types.Ticker, error) {
	url := fmt.Sprintf("%s/products/%s/ticker", coinbaseURI, productID)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		return nil, errors.New("Not Found")
	}

	// unmarshal the response into a Ticker type
	var ticker types.Ticker
	resp.ToJSON(&ticker)

	return &ticker, nil
}
