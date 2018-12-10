/*
Package coinbase is the entry point for the Coinbase Connector.
Connector instantiation, authentication and interface with Coinbase API is handled by this package.

Basics

Start by importing the library

	import (
		cc "github.com/mauricioklein/coinbase-connector"
	)

Before interacting with the system, you need to authenticate with Coinbase. This is done providing a
Credentials object to the Connector constructor, containing your personal Coinbase credentials

	credentials := cc.Credentials{
    	URI:        "<the Coinbase URI>",
    	Key:        "<Your coinbase key>",
    	Secret:     "<Your coinbase secret>",
    	Passphrase: "<Your coinbase passphrase>",
	}

Finally, you can instantiate a connector and start interacting with Coinbase API:

	conn := cc.New(credentials)

Ticker

The ticker for a product can be queried using the "GetTicker" method, passing the productID as parameter:

	tickerResp, err := conn.GetTicker("BTC-USD")

This method returns:
- A TickerResponse struct, with the response from Coinbase (in case of success)
- An error (in case of failure)

Order

To issue a new Order on Coinbase, use the method "CreateOrder", passing a OrderRequest object:

	orderResp, err := conn.CreateOrder(&ccTypes.OrderRequest{
		Size:      1.00,
		Price:     1.00,
		Side:      "buy",
		ProductID: "BTC-USD",
	})

This method returns:
- A OrderResponse struct, with the response from Coinbase (in case of success)
- An error (in case of failure)

Time

Time is used to fetch the current time from Coinbase API:

	timeResp, err := conn.GetTime()

This method returns:
- A TimeResponse struct, with the response from Coinbase (in case of success)
- An error (in case of failure)
*/
package coinbase
