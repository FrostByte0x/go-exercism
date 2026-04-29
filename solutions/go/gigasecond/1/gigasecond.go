// Package gigasecond has functions for gigaseconds
package gigasecond

import "time"

// AddGigasecond returns the time in the input
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
