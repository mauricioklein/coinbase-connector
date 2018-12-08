package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/imroc/req"
	"github.com/stretchr/testify/assert"

	"github.com/mauricioklein/coinbase-connector/types"
)

type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestValidTicker(t *testing.T) {
	ticker := types.Ticker{
		Price:   "5015.00000000",
		Size:    "0.00700000",
		Bid:     "5015",
		Ask:     "6000",
		Volume:  "254.95817196",
		Time:    time.Now(),
		TradeID: 2183411,
	}

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.Path, "/products/BTC-USD/ticker")
		body, _ := json.Marshal(&ticker)

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
		}
	})

	req := req.New()
	req.SetClient(client)
	conn := NewConnector(req)

	resp, err := conn.Ticker("BTC-USD")
	assert.NoError(t, err)

	// Check ticker attributes
	assert.Equal(t, ticker.Ask, resp.Ask)
	assert.Equal(t, ticker.Bid, resp.Bid)
	assert.Equal(t, ticker.Price, resp.Price)
	assert.Equal(t, ticker.Size, resp.Size)
	assert.Equal(t, ticker.Time.Sub(resp.Time), time.Duration(0))
	assert.Equal(t, ticker.TradeID, resp.TradeID)
	assert.Equal(t, ticker.Volume, resp.Volume)
}

func TestInvalidTicker(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.Path, "/products/FOO-BAR/ticker")

		return &http.Response{
			StatusCode: 404,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"message": "NotFound"}`))),
		}
	})

	req := req.New()
	req.SetClient(client)
	conn := NewConnector(req)

	resp, err := conn.Ticker("FOO-BAR")
	assert.Nil(t, resp)
	assert.Error(t, err)
}
