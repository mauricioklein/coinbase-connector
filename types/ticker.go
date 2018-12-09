package types

import "time"

// TickerResponse defines the Coinbase
// response for a Ticker request
type TickerResponse struct {
	Price   string    `json:"price"`
	Size    string    `json:"size"`
	Bid     string    `json:"bid"`
	Ask     string    `json:"ask"`
	Volume  string    `json:"volume"`
	Time    time.Time `json:"time"`
	TradeID int64     `json:"trade_id"`
}
