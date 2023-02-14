package check

import (
	"strings"
)

var _ executor[string] = (*suffix)(nil)

type suffix struct{}

func newSuffix() *suffix {
	return new(suffix)
}

func (p *suffix) check(check string, target string) bool {
	return strings.HasSuffix(target, check)
}
