package types

import "time"

// TimeResponse defines the time response
// returned by Coinbase API
type TimeResponse struct {
	ISO   time.Time
	Epoch float64
}
