package gox

import (
	"encoding/json"
)

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
