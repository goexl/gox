package rand

import (
	"crypto/rand"
	"math/big"
	"time"
)

type duration struct {
	params *durationParams
}

func (d *duration) Generate() (duration time.Duration) {
	if random, re := rand.Int(rand.Reader, big.NewInt(int64(d.params.diff()))); nil != re {
		duration = time.Duration(random.Int64())
	} else {
		duration = d.params.from + time.Duration(random.Int64())
	}

	return
}
