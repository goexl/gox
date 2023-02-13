package gox

const (
	hasTypeAny checkerType = iota
	hasTypeAll
)

var _ = Checker

type (
	checkerType uint8

	checkerBuilder struct {
		typ     checkerType
		checker checker
	}

	checker interface {
		any() bool
		all() bool
	}
)

// Checker 判断是否存在
func Checker() *checkerBuilder {
	return &checkerBuilder{
		typ: hasTypeAny,
	}
}

func (h *checkerBuilder) Any() *checkerBuilder {
	h.typ = hasTypeAny

	return h
}

func (h *checkerBuilder) All() *checkerBuilder {
	h.typ = hasTypeAll

	return h
}

func (h *checkerBuilder) String(check string) (checker *stringChecker) {
	checker = newStringHasBuilder(h, check)
	h.checker = checker

	return
}

func (h *checkerBuilder) Check() (checked bool) {
	switch h.typ {
	case hasTypeAny:
		checked = h.checker.any()
	case hasTypeAll:
		checked = h.checker.all()
	default:
		checked = h.checker.any()
	}

	return
}
