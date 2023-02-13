package gox

import (
	"strings"
)

const (
	stringCheckTypePrefix stringCheckerType = iota
	stringCheckTypeSuffix
)

type (
	stringCheckerType uint8

	stringChecker struct {
		*checkerBuilder

		check  string
		checks []string
		typ    stringCheckerType
	}
)

func newStringHasBuilder(builder *checkerBuilder, check string) *stringChecker {
	return &stringChecker{
		checkerBuilder: builder,

		check: check,
		typ:   stringCheckTypePrefix,
	}
}

func (s *stringChecker) Prefix(checks ...string) *stringChecker {
	s.typ = stringCheckTypePrefix
	s.checks = append(s.checks, checks...)

	return s
}

func (s *stringChecker) Suffix(checks ...string) *stringChecker {
	s.typ = stringCheckTypeSuffix
	s.checks = append(s.checks, checks...)

	return s
}

func (s *stringChecker) any() (match bool) {
	switch s.typ {
	case stringCheckTypePrefix:
		match = s.anyPrefix(s.check, s.checks)
	case stringCheckTypeSuffix:
		match = s.anySuffix(s.check, s.checks)
	default:
		match = s.anyPrefix(s.check, s.checks)
	}

	return
}

func (s *stringChecker) all() (match bool) {
	switch s.typ {
	case stringCheckTypePrefix:
		match = s.allPrefix(s.check, s.checks)
	case stringCheckTypeSuffix:
		match = s.allSuffix(s.check, s.checks)
	default:
		match = s.allPrefix(s.check, s.checks)
	}

	return
}

func (s *stringChecker) anyPrefix(str string, checks []string) (has bool) {
	has = false
	for _, check := range checks {
		if strings.HasPrefix(str, check) {
			has = true
		}

		if has {
			break
		}
	}

	return
}

func (s *stringChecker) allPrefix(str string, checks []string) (has bool) {
	has = true
	for _, check := range checks {
		if !strings.HasPrefix(str, check) {
			has = false
		}

		if !has {
			break
		}
	}

	return
}

func (s *stringChecker) anySuffix(str string, checks []string) (has bool) {
	has = false
	for _, check := range checks {
		if strings.HasSuffix(str, check) {
			has = true
		}

		if has {
			break
		}
	}

	return
}

func (s *stringChecker) allSuffix(str string, checks []string) (has bool) {
	has = true
	for _, check := range checks {
		if !strings.HasSuffix(str, check) {
			has = false
		}

		if !has {
			break
		}
	}

	return
}
