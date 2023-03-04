package rand

import (
	"math/rand"
)

type stringBuilder struct {
	source rand.Source
	params *stringParams
}

func newStringBuilder(source rand.Source) *stringBuilder {
	return &stringBuilder{
		source: source,
		params: newStringParams(),
	}
}

func (s *stringBuilder) Length(length int) *stringBuilder {
	s.params.length = length

	return s
}

func (s *stringBuilder) Code() *stringBuilder {
	s.params.length = 6
	s.params.letters = s.params.numbers

	return s
}

func (s *stringBuilder) Digital() *stringBuilder {
	s.params.letters = s.params.numbers

	return s
}
