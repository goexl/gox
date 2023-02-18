package check

import (
	"strings"
)

var _ executor[string] = (*contains)(nil)

type contains struct{}

func newContains() *contains {
	return new(contains)
}

// nolint: unused
func (c *contains) check(check string, item string) bool {
	return strings.Contains(check, item)
}
