package check

import (
	"strings"
)

var _ executor[string] = (*suffix)(nil)

type suffix struct{}

func newSuffix() *suffix {
	return new(suffix)
}

// nolint: unused
func (s *suffix) check(check string, item string) bool {
	return strings.HasSuffix(check, item)
}
