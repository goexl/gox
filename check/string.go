package check

type _string struct {
	typ     checkerType
	check   string
	targets []string
}

func newString(typ checkerType, check string) *_string {
	return &_string{
		typ:   typ,
		check: check,
	}
}

func (s *_string) Targets(targets ...string) *_string {
	s.targets = append(s.targets, targets...)

	return s
}

func (s *_string) Prefix() *checker[string] {
	return newChecker[string](s.typ, s.check, s.targets, newPrefix())
}

func (s *_string) Suffix() *checker[string] {
	return newChecker[string](s.typ, s.check, s.targets, newSuffix())
}

func (s *_string) Contains() *checker[string] {
	return newChecker[string](s.typ, s.check, s.targets, newContains())
}
