package gox

import (
	`encoding/json`
	`fmt`
	`strconv`
	`strings`
)

// Int64s 长整形列表
type Int64s []int64

func (s Int64s) MarshalJSON() (bytes []byte, err error) {
	values := make([]string, len(s))
	for index, value := range []int64(s) {
		values[index] = fmt.Sprintf(`"%d"`, value)
	}
	bytes = []byte(fmt.Sprintf("[%s]", strings.Join(values, ",")))

	return
}

func (s *Int64s) UnmarshalJSON(bytes []byte) (err error) {
	var int64Values []int64
	if err = json.Unmarshal(bytes, &int64Values); nil == err {
		*s = int64Values
		return
	} else {
		var stringValues []string
		if err = json.Unmarshal(bytes, &stringValues); nil != err {
			return
		}
		err = s.toInt64Slice(stringValues)
	}

	return
}

func (s *Int64s) UnmarshalParam(src string) (err error) {
	if "" == src {
		return
	}
	values := strings.Split(src, ",")
	err = s.toInt64Slice(values)

	return
}

func (s *Int64s) toInt64Slice(values []string) (err error) {
	realValue := int64(0)
	*s = make([]int64, len(values))
	for index, value := range values {
		if realValue, err = strconv.ParseInt(value, 10, 64); nil != err {
			return
		}
		(*s)[index] = realValue
	}

	return
}
