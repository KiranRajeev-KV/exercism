

// Package gigasecond measure time in gigaseconds.
package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
    t = t.Add(1000000000 * time.Second)
	return t
}
