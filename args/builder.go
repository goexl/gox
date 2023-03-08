package args

import (
	"github.com/goexl/gox"
)

// Builder 参数列表构建器
type Builder struct {
	args   []any
	params *params
}

func newBuilder(params *params) *Builder {
	return &Builder{
		args:   make([]any, 0, params.capacity),
		params: params,
	}
}

func (b *Builder) Insert(args ...any) *Builder {
	newArgs := make([]any, 0, len(b.args)+len(args))
	newArgs = append(newArgs, args...)
	newArgs = append(newArgs, b.args...)
	b.args = newArgs

	return b
}

func (b *Builder) Add(args ...any) *Builder {
	b.args = append(b.args, args...)

	return b
}

func (b *Builder) Arg(key string, value any) *Builder {
	placeholder := b.params.long
	if 1 == len(key) {
		placeholder = b.params.short
	}
	arg := gox.StringBuilder(placeholder, key, b.params.equal, value).String()
	b.args = append(b.args, arg)

	return b
}

func (b *Builder) Flag(key string) *Builder {
	placeholder := b.params.long
	if 1 == len(key) {
		placeholder = b.params.short
	}
	b.args = append(b.args, gox.StringBuilder(placeholder, key).String())

	return b
}

func (b *Builder) Build() *Args {
	return newArgs(b.args)
}
