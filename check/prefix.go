package check

import (
	"strings"
)

var _ executor[string] = (*prefix)(nil)

type prefix struct{}

func newPrefix() *prefix {
	return new(prefix)
}

func (p *prefix) check(check string, target string) bool {
	return strings.HasPrefix(target, check)
}
