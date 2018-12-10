/*
Package types defines the schema for requests and responses for Coinbase API.
All structs are assigned with JSON directives, which matches exactly the payload sent and received from Coinbase.

Types

"TickerResponse" defined the payload from a response of a ticker request to Coinbase:

	type TickerResponse struct {
		Price   string    `json:"price"`
		Size    string    `json:"size"`
		Bid     string    `json:"bid"`
		Ask     string    `json:"ask"`
		Volume  string    `json:"volume"`
		Time    time.Time `json:"time"`
		TradeID int64     `json:"trade_id"`
	}

"OrderRequest" defines the payload to an order creation request to Coinbase:

	type OrderRequest struct {
		Size      float64 `json:"size"`
		Price     float64 `json:"price"`
		Side      string  `json:"side"`
		ProductID string  `json:"product_id"`
	}

"OrderResponse" defines the payload from a response of a order creation request to Coinbase:

	type OrderResponse struct {
		ID            string    `json:"id"`
		Price         string    `json:"price"`
		Size          string    `json:"size"`
		ProductID     string    `json:"product_id"`
		Side          string    `json:"side"`
		Stp           string    `json:"stp"`
		Type          string    `json:"type"`
		TimeInForce   string    `json:"time_in_force"`
		PostOnly      bool      `json:"post_only"`
		CreatedAt     time.Time `json:"created_at"`
		FillFees      string    `json:"fill_fees"`
		FilledSize    string    `json:"filled_size"`
		ExecutedValue string    `json:"executed_value"`
		Status        string    `json:"status"`
		Settled       bool      `json:"settled"`
	}

"TimeResponse" defines the payload from a response of a time request to Coinbase:

	type TimeResponse struct {
		ISO   time.Time
		Epoch float64
	}
*/
package types
