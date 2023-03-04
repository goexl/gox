package rand

import "time"

type timeParams struct {
	from time.Time
	to time.Time
}

func newTimeParams() *timeParams {
	return new(timeParams)
}

func (tp *timeParams) diff() time.Duration {
	return tp.to.Sub(tp.from)
}
