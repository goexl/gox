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

func (ab *argsBuilder) Short(key any, value any) *argsBuilder {
	return ab.arg(key, value, ab.params.short)
}

func (ab *argsBuilder) Long(key any, value any) *argsBuilder {
	return ab.arg(key, value, ab.params.long)
}

func (ab *argsBuilder) Build() Args {
	return ab.args
}

func (ab *argsBuilder) arg(key any, value any, placeholder string) *argsBuilder {
	arg := gox.StringBuilder(placeholder, key, ab.params.equal, value).String()
	ab.args = append(ab.args, arg)

	return ab
}
