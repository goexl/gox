package args

import (
	"github.com/goexl/gox"
)

// Builder 参数列表构建器
type Builder struct {
	args   []any
	params *params
	_      gox.CannotCopy
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

func (b *Builder) Subcommand(command string, others ...string) *Builder {
	b.args = append(b.args, command)
	for _, other := range others {
		b.args = append(b.args, other)
	}

	return b
}

func (b *Builder) Arg(key string, value any) *Builder {
	arg := gox.StringBuilder(b.key(key), key, b.params.equal, value).String()
	b.args = append(b.args, arg)

	return b
}

func (b *Builder) Args(key string, value any, others ...any) *Builder {
	b.args = append(b.args, b.key(key), gox.ToString(value))
	for _, other := range others {
		b.args = append(b.args, gox.ToString(other))
	}

	return b
}

func (b *Builder) Flag(flags ...string) *Builder {
	for _, flag := range flags {
		b.args = append(b.args, b.key(flag))
	}

	return b
}

func (b *Builder) Build() *Args {
	return newArgs(b.args)
}

func (b *Builder) key(original string) (final string) {
	placeholder := b.params.long
	if 1 == len(original) {
		placeholder = b.params.short
	}
	final = gox.StringBuilder(placeholder, original).String()

	return
}
