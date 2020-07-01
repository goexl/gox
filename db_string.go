package gox

type DBString string

func (s DBString) ToDB() ([]byte, error) {
	if "" == s {
		return nil, nil
	}

	return []byte(s), nil
}

func (s *DBString) FromDB(bytes []byte) error {
	if 0 == len(bytes) {
		*s = ""
	}
	*s = DBString(bytes)

	return nil
}
