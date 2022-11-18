package gox

type stringField struct {
	key   string
	value string
}

func newStringField(key string, value string) *stringField {
	return &stringField{
		key:   key,
		value: value,
	}
}

func (s *stringField) Key() string {
	return s.key
}

func (s *stringField) Value() any {
	return s.value
}
