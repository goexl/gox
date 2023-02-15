package check

import (
	"strings"
)

var _ executor[string] = (*suffix)(nil)

type suffix struct{}

func newSuffix() *suffix {
	return new(suffix)
}

func (s *suffix) check(check string, item string) bool {
	return strings.HasSuffix(item, check)
}
