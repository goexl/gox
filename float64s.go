package gox

import (
	`encoding/json`
	`fmt`
	`strconv`
	`strings`
)

// Float64s 浮点型列表
type Float64s []float64

func (f Float64s) MarshalJSON() (bytes []byte, err error) {
	values := make([]string, len(f))
	for index, value := range []float64(f) {
		values[index] = fmt.Sprintf(`"%f"`, value)
	}
	bytes = []byte(fmt.Sprintf("[%s]", strings.Join(values, ",")))

	return
}

func (f *Float64s) UnmarshalJSON(bytes []byte) (err error) {
	var int64Values []float64
	if err = json.Unmarshal(bytes, &int64Values); nil == err {
		*f = int64Values
		return
	} else {
		var stringValues []string
		if err = json.Unmarshal(bytes, &stringValues); nil != err {
			return
		}
		err = f.toFloat64Slice(stringValues)
	}

	return
}

func (f *Float64s) UnmarshalParam(src string) (err error) {
	if "" == src {
		return
	}
	values := strings.Split(src, ",")
	err = f.toFloat64Slice(values)

	return
}

func (f *Float64s) toFloat64Slice(values []string) (err error) {
	realValue := float64(0)
	*f = make([]float64, len(values))
	for index, value := range values {
		if realValue, err = strconv.ParseFloat(value, 64); nil != err {
			return
		}
		(*f)[index] = realValue
	}

	return
}
