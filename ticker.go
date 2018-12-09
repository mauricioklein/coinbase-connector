package coinbase

import (
	"errors"
	"fmt"

	"github.com/mauricioklein/coinbase-connector/types"
)

// GetTicker returns the ticker for the given productID
func (c *Connector) GetTicker(productID string) (*types.TickerResponse, error) {
	url := fmt.Sprintf("%s/products/%s/ticker", c.credentials.URI, productID)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.Response().StatusCode != 200 {
		return nil, errors.New("Not Found")
	}

	// unmarshal the response into a Ticker type
	var ticker types.TickerResponse
	resp.ToJSON(&ticker)

	return &ticker, nil
}
