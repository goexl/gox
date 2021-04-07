package gox

import (
	"encoding/json"
	`fmt`
)

// JSONSyntaxError JSON语法错误
type JSONSyntaxError struct {
	original []byte
}

func (j JSONSyntaxError) Error() string {
	return fmt.Sprintf("错误的JSON语法%q", string(j.original))
}

// NewJSONSyntaxError 创建一个JSON语法错误
func NewJSONSyntaxError(original []byte) error {
	return &JSONSyntaxError{
		original: original,
	}
}

// JSONString 转换成JSON字符串
func JSONString(obj interface{}) string {
	str := ""
	if data, err := json.Marshal(obj); nil != err {
		str = ""
	} else {
		str = string(data)
	}

	return str
}
