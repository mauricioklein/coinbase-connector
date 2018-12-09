package main

import (
	"fmt"

	"github.com/mauricioklein/coinbase-connector/types"
)

// GetTime returns the time from the Coinbase API
func (c *Connector) GetTime() (*types.TimeResponse, error) {
	url := fmt.Sprintf("%s/time", coinbaseURI)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	// unmarshal the response into a Ticker type
	var time types.TimeResponse
	resp.ToJSON(&time)

	return &time, nil
}
