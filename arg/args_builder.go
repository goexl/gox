package arg

import (
	"github.com/goexl/gox"
)

type argsBuilder struct {
	args   Args
	params *params
}

func newArgsBuilder(params *params) *argsBuilder {
	return &argsBuilder{
		args:   make(Args, 0, params.size),
		params: params,
	}
}

func (ab *argsBuilder) Insert(args ...Arg) *argsBuilder {
	new := make(Args, 0, len(ab.args) + len(args))
	new = append(new, gox.ToStrings(args...)...)
	new = append(new, ab.args...)
	ab.args = new

	return ab
}

func (ab *argsBuilder) Add(args ...Arg) *argsBuilder {
	ab.args = append(ab.args, gox.ToStrings(args...)...)

	return ab
}

func (a *argsBuilder) Short(key any, value any) *argsBuilder {
	return a.arg(key, value, a.params.short)
}

func (a *argsBuilder) Long(key any, value any) *argsBuilder {
	return a.arg(key, value, a.params.long)
}

func (a *argsBuilder) Build() Args {
	return a.args
}

func (a *argsBuilder) arg(key any, value any, placeholder string) *argsBuilder {
	a.args = append(a.args, gox.StringBuilder(placeholder, key, equal, value).String())

	return a
}
