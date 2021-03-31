package gox

type String string

func (s String) ToDB() ([]byte, error) {
	if "" == s {
		return nil, nil
	}

	return []byte(s), nil
}

func (s *String) FromDB(bytes []byte) error {
	if 0 == len(bytes) {
		*s = ""
	}
	*s = String(bytes)

	return nil
}
