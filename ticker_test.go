package coinbase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/mauricioklein/coinbase-connector/types"

	"gopkg.in/h2non/gock.v1"
)

func TestGetTicker_Success(t *testing.T) {
	credentials := Credentials{
		URI:        "https://api-public.sandbox.pro.coinbase.com",
		Key:        "A_Key",
		Secret:     "A_Secret",
		Passphrase: "A_Passphrase",
	}

	tickerResponse := types.TickerResponse{
		Price:   "5015.00000000",
		Size:    "0.00700000",
		Bid:     "5015",
		Ask:     "6000",
		Volume:  "254.95817196",
		Time:    time.Now(),
		TradeID: 2183411,
	}

	defer gock.Off()

	gock.New(credentials.URI).
		Get("products/BTC-USD/ticker").
		Reply(200).
		JSON(tickerResponse)

	conn := New(credentials)

	gock.InterceptClient(conn.Client())

	resp, err := conn.GetTicker("BTC-USD")
	assert.NoError(t, err)

	// Check ticker attributes
	assert.Equal(t, tickerResponse.Ask, resp.Ask)
	assert.Equal(t, tickerResponse.Bid, resp.Bid)
	assert.Equal(t, tickerResponse.Price, resp.Price)
	assert.Equal(t, tickerResponse.Size, resp.Size)
	assert.Equal(t, tickerResponse.Time.Sub(resp.Time), time.Duration(0))
	assert.Equal(t, tickerResponse.TradeID, resp.TradeID)
	assert.Equal(t, tickerResponse.Volume, resp.Volume)
}

func TestGetTicker_Error(t *testing.T) {
	credentials := Credentials{
		URI:        "https://api-public.sandbox.pro.coinbase.com",
		Key:        "abcd1234efgh5678",
		Secret:     "ABCDEFGH12345678",
		Passphrase: "arandompassphrase",
	}

	defer gock.Off()

	gock.New(credentials.URI).
		Get("products/FOO-BAR/ticker").
		Reply(404).
		JSON(`{"message": "NotFound"}`)

	conn := New(credentials)

	gock.InterceptClient(conn.Client())

	resp, err := conn.GetTicker("FOO-BAR")
	assert.Nil(t, resp)
	assert.Error(t, err)
}
