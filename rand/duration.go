package rand

import (
	"crypto/rand"
	"math/big"
	"time"
)

type duration struct {
	params *durationParams
}

func newDuration(params *durationParams) *duration {
	return &duration{
		params: params,
	}
}

func (d *duration) Generate() (duration time.Duration) {
	diff := d.params.diff()
	if random, re := rand.Int(rand.Reader, big.NewInt(int64(diff))); nil != re {
		duration = diff
	} else {
		duration = d.params.from + time.Duration(random.Int64())
	}

	return
}
