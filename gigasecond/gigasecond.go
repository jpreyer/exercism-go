package gigasecond

import "time"

// AddGigasecond : Given a time, returns the Time 10**9 seconds after that
func AddGigasecond(t time.Time) time.Time {
    t = t.Add(time.Second * (1000000000))
	return t
}
