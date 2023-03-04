package rand

import (
	"crypto/rand"
	"math/big"
	"time"
)

type duration struct {
	params *durationParams
}

func (d *duration) G() time.Duration {
		from := time.Duration(int64(float64(b.Backoff) * 0.3))
	if duration, re := rand.Int(rand.Reader, big.NewInt(int64(b.Backoff-from))); nil != re {
		backoff = b.Backoff
	} else {
		backoff = from + time.Duration(duration.Int64()).Truncate(time.Second)
	}

	return backoff
}
