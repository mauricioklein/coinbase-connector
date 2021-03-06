package coinbase

import (
	"encoding/json"
	"fmt"

	"github.com/imroc/req"

	"github.com/mauricioklein/coinbase-connector/types"
)

// CreateOrder creates a new order on Coinbase
func (c *Connector) CreateOrder(orderReq *types.OrderRequest) (*types.OrderResponse, error) {
	url := fmt.Sprintf("%s/orders", c.credentials.URI)

	time, err := c.GetTime()
	if err != nil {
		return nil, err
	}

	timestamp := fmt.Sprintf("%f", time.Epoch)

	m, err := json.Marshal(orderReq)
	if err != nil {
		return nil, err
	}

	body := string(m)

	sign, err := generateSignature(body, timestamp, c.credentials.Secret)
	if err != nil {
		return nil, err
	}

	headers := req.Header{
		"Accept":               "application/json",
		"Content-Type":         "application/json",
		"CB-ACCESS-KEY":        c.credentials.Key,
		"CB-ACCESS-SIGN":       sign,
		"CB-ACCESS-TIMESTAMP":  timestamp,
		"CB-ACCESS-PASSPHRASE": c.credentials.Passphrase,
	}

	resp, err := c.req.Post(url, headers, m)
	if err != nil {
		return nil, err
	}

	// unmarshal the response into a Ticker type
	var orderResp types.OrderResponse
	resp.ToJSON(&orderResp)

	return &orderResp, nil
}
