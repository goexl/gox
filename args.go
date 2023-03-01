package gox

// Args 参数列表
type Args []any

func (a Args) Add(args ...any) (new Args) {
	new = make(Args, 0, len(a)+len(args))
	new = append(new, a...)
	new = append(new, args...)
	args = nil

	return
}
