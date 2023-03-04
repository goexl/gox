package rand

import (
	"math/rand"
	"time"
)

type stringBuilder struct {
	source rand.Source
	params *stringParams
}

func newStringBuilder() *stringBuilder {
	return &stringBuilder{
		source: rand.NewSource(time.Now().UnixNano()),
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

func (sb *stringBuilder) Build() *_string {
	return newString(sb.source, sb.params)
}
