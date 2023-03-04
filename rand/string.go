package rand

import (
	"math/rand"
	"time"
)

type _string struct {
	source rand.Source
	params *stringParams
}

func newString(params *stringParams) *_string {
	return &_string{
		source: rand.NewSource(time.Now().UnixNano()),
		params: params,
	}
}

func (s *_string) Generate() string {
	return s.params.rand(s.source)
}

