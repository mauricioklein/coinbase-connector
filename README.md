[![Build Status](https://travis-ci.org/mauricioklein/coinbase-connector.svg?branch=master)](https://travis-ci.org/mauricioklein/coinbase-connector)
[![Coverage Status](https://coveralls.io/repos/github/mauricioklein/coinbase-connector/badge.svg?branch=master)](https://coveralls.io/github/mauricioklein/coinbase-connector?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mauricioklein/coinbase-connector)](https://goreportcard.com/report/github.com/mauricioklein/coinbase-connector)
[![GoDoc](https://godoc.org/github.com/mauricioklein/coinbase-connector?status.svg)](https://godoc.org/github.com/mauricioklein/coinbase-connector)

# coinbase-connector

A simple connector to Coinbase API

## Dependencies

- Go 1.11
- Go modules support

## Instalation

```bash
$ go get github.com/mauricioklein/coinbase-connector
```

## Setup

```bash
# Make sure you have Go modules active
$ export GO111MODULE=on

# Download the dependencies
$ go get -v ./...
```

## Usage

```go
// Import library
import (
	"fmt"

	cc "github.com/mauricioklein/coinbase-connector"
	ccTypes "github.com/mauricioklein/coinbase-connector/types"
)

credentials := cc.Credentials{
    URI:        "<the Coinbase URI>",
    Key:        "<Your coinbase key>",
    Secret:     "<Your coinbase secret>",
    Passphrase: "<Your coinbase passphrase>",
}

// creates a connector instance
conn := cc.New(credentials)

// Get a ticker
tickerResp, err := conn.GetTicker("BTC-USD")

/*
    *types.TickerResponse{
        Price: "4999.00000000",
        Size: "0.00100000",
        Bid: "4649.07",
        Ask: "5000",
        Volume: "344.18940008",
        Time: 2018-12-10 14:33:40.699 +0000 UTC,
        TradeID: 2183687
    }, nil
*/

// Place an order
orderResp, err := conn.CreateOrder(&ccTypes.OrderRequest{
    Size:      1.00,
    Price:     1.00,
    Side:      "buy",
    ProductID: "BTC-USD",
})

/*
    *types.OrderResponse{
        ID: "a8f6c2ab-0ffb-4424-a002-026c468532b2",
        Price: "1.00000000",
        Size: "1.00000000",
        ProductID: "BTC-USD",
        Side: "buy",
        Stp: "dc",
        Type: "limit",
        TimeInForce: "GTC",
        PostOnly: false,
        CreatedAt: 2018-12-10 14:48:52.843354 +0000 UTC,
        FillFees: "0.0000000000000000",
        FilledSize: "0.00000000",
        ExecutedValue: "0.0000000000000000",
        Status: "pending",
        Settled: false
    }, nil
*/

// Get the Coinbase server time
timeResp, err := conn.GetTime()

/*
    *types.TimeResponse{
        ISO: 2018-12-10 15:20:25.391 +0000 UTC,
        Epoch: 1.544455225391e+09
    }
*/
```

## Run tests

```bash
$ go test -v -race ./...
```
