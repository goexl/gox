package rand

import (
	"math/rand"
)

type _string struct {
	source rand.Source
	params *stringParams
}

func newString(source rand.Source, params *stringParams) *_string {
	return &_string{
		source: source,
		params: params,
	}
}

func (s *stringBuilder) Generate() string {
	return s.params.rand(s.source)
}

