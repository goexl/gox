package json

import (
	"encoding/json"
)

var _ = String

// String 转换成JSON字符串
func String(obj any) string {
	str := ""
	if data, err := json.Marshal(obj); nil != err {
		str = ""
	} else {
		str = string(data)
	}

	return str
}
