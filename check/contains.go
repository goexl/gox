package check

import (
	"strings"
)

var _ executor[string] = (*contains)(nil)

type contains struct{}

func newContains() *contains {
	return new(contains)
}

func (p *contains) check(check string, item string) bool {
	return strings.Contains(item, check)
}
