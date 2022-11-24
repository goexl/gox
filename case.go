package gox

import (
	"strings"
	"unicode"
)

// Case 字符串转换
type Case string

// Underscore 转为下划线写法
func (c Case) Underscore(upperInitial bool) string {
	buffer := strings.Builder{}
	from := string(c)
	for index, char := range from {
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

// Camel 转为驼峰写法
func (c Case) Camel() string {
	from := string(c)
	from = strings.Replace(from, "_", " ", -1)
	from = strings.Title(from)

	return strings.Replace(from, " ", "", -1)
}

// InitialLowercase 首字母小写
func (c Case) InitialLowercase() Case {
	from := string(c)
	c = Case(string(unicode.ToLower(rune(from[0]))) + from[1:])

	return c
}
