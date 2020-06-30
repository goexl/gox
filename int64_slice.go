package gox

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Int64Slice []int64

func (s Int64Slice) MarshalJSON() ([]byte, error) {
	return s.marshalJSON(false)
}

func (s Int64Slice) ToDB() ([]byte, error) {
	return s.marshalJSON(true)
}

func (s *Int64Slice) UnmarshalJSON(bytes []byte) error {
	return s.unmarshalJSON(bytes)
}

func (s *Int64Slice) FromDB(bytes []byte) error {
	return s.unmarshalJSON(bytes)
}

func (s Int64Slice) marshalJSON(toDb bool) ([]byte, error) {
	values := make([]string, len(s))
	for i, value := range []int64(s) {
		if toDb {
			values[i] = fmt.Sprintf(`%v`, value)
		} else {
			values[i] = fmt.Sprintf(`"%v"`, value)
		}
	}
	bytes := []byte(fmt.Sprintf("[%v]", strings.Join(values, ",")))

	return bytes, nil
}

func (s *Int64Slice) unmarshalJSON(bytes []byte) (err error) {
	var values []int64
	if err = json.Unmarshal(bytes, &values); err == nil {
		*s = values
		return
	}

	var sValues []string
	if err = json.Unmarshal(bytes, &sValues); err == nil {
		*s = make([]int64, len(values))
		for i, value := range sValues {
			value, err := strconv.ParseInt(value, 10, 64)
			if nil != err {
				return err
			}
			(*s)[i] = value
		}
	}

	return
}
