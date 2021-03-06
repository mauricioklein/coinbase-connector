package coinbase

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/mauricioklein/coinbase-connector/types"

	"gopkg.in/h2non/gock.v1"
)

func TestCreateOrder_Success(t *testing.T) {
	credentials := Credentials{
		URI:        "https://api-public.sandbox.pro.coinbase.com",
		Key:        "abcd1234efgh5678",
		Secret:     "ABCDEFGH12345678",
		Passphrase: "arandompassphrase",
	}

	orderRequest := types.OrderRequest{
		Price:     0.01,
		Size:      0.100,
		Side:      "buy",
		ProductID: "BTC-USD",
	}

	orderResponse := types.OrderResponse{
		ID:        "ABC123",
		Price:     "1.00",
		Size:      "1.00",
		ProductID: "BTC-USD",
	}

	timeResponse := types.TimeResponse{
		ISO:   time.Time{},
		Epoch: 123456.00,
	}

	defer gock.Off()

	orderReqMarshal, _ := json.Marshal(&orderRequest)
	orderReqJSON := string(orderReqMarshal)

	// Mock time request to Coinbase
	gock.New(credentials.URI).
		Get("/time").
		Reply(200).
		JSON(timeResponse)

	// Mock POST /orders request to Coinbase
	gock.New(credentials.URI).
		Post("/orders").
		MatchHeaders(map[string]string{
			"Accept":               "application/json",
			"Content-Type":         "application/json",
			"CB-ACCESS-KEY":        credentials.Key,
			"CB-ACCESS-SIGN":       ".*",
			"CB-ACCESS-TIMESTAMP":  fmt.Sprintf("%f", timeResponse.Epoch),
			"CB-ACCESS-PASSPHRASE": credentials.Passphrase,
		}).
		MatchType("json").
		JSON(orderReqJSON).
		Reply(200).
		JSON(orderResponse)

	conn := New(credentials)

	gock.InterceptClient(conn.Client())

	resp, err := conn.CreateOrder(&orderRequest)

	assert.NoError(t, err)

	// Check order attributes
	assert.Equal(t, resp.ID, orderResponse.ID)
	assert.Equal(t, resp.Price, orderResponse.Price)
	assert.Equal(t, resp.Size, orderResponse.Size)
	assert.Equal(t, resp.ProductID, orderResponse.ProductID)
}
