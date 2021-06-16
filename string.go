package gox

import (
	`strings`
	`unicode`
)

// UnderscoreName 驼峰式写法转为下划线写法
func UnderscoreName(name string, upperInitial bool) string {
	buffer := strings.Builder{}
	for index, char := range name {
		if unicode.IsUpper(char) {
			if 0 != index {
				buffer.WriteRune('_')
			}
			if upperInitial {
				buffer.WriteRune(unicode.ToUpper(char))
			} else {
				buffer.WriteRune(unicode.ToLower(char))
			}
		} else {
			if 0 == index && upperInitial {
				buffer.WriteRune(unicode.ToUpper(char))
			} else {
				buffer.WriteRune(char)
			}
		}
	}

	return buffer.String()
}

// CamelName 下划线写法转为驼峰写法
func CamelName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)

	return strings.Replace(name, " ", "", -1)
}

// SearchString 字符串查询
func SearchString(slice []string, s string) int {
	index := -1

	for i, v := range slice {
		if s == v {
			index = i
		}
	}

	return index
}

// Contract 字符串连接
func Contract(separator string, contracts ...string) (result string) {
	buffer := strings.Builder{}

	for index, contract := range contracts {
		buffer.WriteString(contract)
		if len(contracts)-1 != index {
			buffer.WriteString(separator)
		}
	}
	result = buffer.String()

	return
}

// InitialLowercase 首字母小写
func InitialLowercase(from string) (to string) {
	for i, v := range from {
		to = string(unicode.ToLower(v)) + from[i+1:]
		break
	}

	return
}
