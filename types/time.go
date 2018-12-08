package types

import "time"

// Time defines the time response returned by
// Coinbase API
type Time struct {
	ISO   time.Time
	Epoch float64
}
