package check

type checker[T any] struct {
	typ      checkerType
	check    T
	targets  []T
	executor executor[T]
}

func newChecker[T any](typ checkerType, check T, targets []T, executor executor[T]) *checker[T] {
	return &checker[T]{
		typ:      typ,
		check:    check,
		targets:  targets,
		executor: executor,
	}
}

func (c *checker[T]) Check() (check bool) {
	switch c.typ {
	case hasTypeAny:
		check = c.any()
	case hasTypeAll:
		check = c.all()
	default:
		check = c.any()
	}

	return
}

func (c *checker[T]) any() (check bool) {
	check = false
	for _, target := range c.targets {
		if c.executor.check(c.check, target) {
			check = true
		}

		if check {
			break
		}
	}

	return
}

func (c *checker[T]) all() (check bool) {
	check = true
	for _, target := range c.targets {
		if !c.executor.check(c.check, target) {
			check = false
		}

		if !check {
			break
		}
	}

	return
}
