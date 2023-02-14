package check

type _string struct {
	typ   checkerType
	check string
	items []string
}

func newString(typ checkerType, check string) *_string {
	return &_string{
		typ:   typ,
		check: check,
	}
}

func (s *_string) Items(items ...string) *_string {
	s.items = append(s.items, items...)

	return s
}

func (s *_string) Prefix() *checker[string] {
	return newChecker[string](s.typ, s.check, s.items, newPrefix())
}

func (s *_string) Suffix() *checker[string] {
	return newChecker[string](s.typ, s.check, s.items, newSuffix())
}

func (s *_string) Contains() *checker[string] {
	return newChecker[string](s.typ, s.check, s.items, newContains())
}
