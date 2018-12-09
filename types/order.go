package types

import "time"

// OrderRequest defines a order creation
// payload to Coinbase
type OrderRequest struct {
	Size      float64 `json:"size"`
	Price     float64 `json:"price"`
	Side      string  `json:"side"`
	ProductID string  `json:"product_id"`
}

// OrderResponse defines the response from an
// order creation in Coinbase
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
