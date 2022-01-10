package exec

import (
	`io`
)

type (
	option interface {
		apply(options *options)
	}

	options struct {
		args   []string
		stdout io.Writer
		stderr io.Writer
	}
)

func defaultOptions() *options {
	return &options{}
}
