package gox

// Args 参数列表
type Args []any

func (a Args) Add(args ...any) Args {
	return append(a, args...)
}
