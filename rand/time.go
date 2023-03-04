package rand

import (
	"crypto/rand"
	"math/big"
	"time"
)

type _time struct {
	params *timeParams
}

func newTime(params *timeParams) *_time {
	return &_time{
		params: params,
	}
}

func (t *_time) Generate() (_time time.Time) {
	diff := t.params.diff()
	if random, re := rand.Int(rand.Reader, big.NewInt(int64(diff))); nil != re {
		_time = t.params.from
	} else {
		_time = t.params.from.Add(time.Duration(random.Int64()))
	}

	return
}
