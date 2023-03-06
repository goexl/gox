package arg

import (
	"github.com/goexl/gox"
)

type arguments struct {
	args   Args
	params *params
}

func newArguments(params *params) *arguments {
	return &arguments{
		args: make(Args, 0, params.size),
	}
}

func (a *arguments) Add(args ...Arg) *arguments {
	a.args = append(a.args, args...)

	return a
}

func (a *arguments) Short(key any, value any) *arguments {
	return a.arg(key, value, a.params.short)
}

func (a *arguments) Long(key any, value any) *arguments {
	return a.arg(key, value, a.params.long)
}

func (a *arguments) Build() Args {
	return a.args
}

func (a *arguments) arg(key any, value any, placeholder string) *arguments {
	a.args = append(a.args, gox.StringBuilder(placeholder, key, equal, value).String())

	return a
}
