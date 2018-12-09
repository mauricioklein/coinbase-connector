package main

import (
	"testing"
	"time"

	"github.com/imroc/req"
	"github.com/stretchr/testify/assert"

	"github.com/mauricioklein/coinbase-connector/types"

	"gopkg.in/h2non/gock.v1"
)

func TestGetTime(t *testing.T) {
	timeResponse := types.TimeResponse{
		ISO:   time.Time{},
		Epoch: 123456.00,
	}

	defer gock.Off()

	// Mock time request to Coinbase
	gock.New("https://api-public.sandbox.pro.coinbase.com/").
		Get("/time").
		Reply(404).
		JSON(timeResponse)

	req := req.New()
	conn := NewConnector(req)

	gock.InterceptClient(req.Client())

	resp, err := conn.GetTime()

	assert.NoError(t, err)

	// Check order attributes
	assert.Equal(t, resp.ISO, timeResponse.ISO)
	assert.Equal(t, resp.Epoch, timeResponse.Epoch)
}
