package rand

import "time"

type durationParams struct {
	from time.Duration
	to time.Duration
}

func newDurationParams() *durationParams {
	return new(durationParams)
}
