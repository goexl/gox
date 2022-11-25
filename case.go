package gox

import (
	"strings"
	"unicode"
)

// Case 字符串转换
type Case string

// String 最终结果
func (c *Case) String() string {
	return string(*c)
}

// Underscore 转为下划线写法
func (c *Case) Underscore(upper bool) *Case {
	buffer := strings.Builder{}
	from := string(*c)
	for index, char := range from {
		if unicode.IsUpper(char) {
			if 0 != index {
				buffer.WriteRune('_')
			}
			if upper {
				buffer.WriteRune(unicode.ToUpper(char))
			} else {
				buffer.WriteRune(unicode.ToLower(char))
			}
		} else {
			if 0 == index && upper {
				buffer.WriteRune(unicode.ToUpper(char))
			} else {
				buffer.WriteRune(char)
			}
		}
	}
	*c = Case(buffer.String())

	return c
}

// Camel 转为驼峰写法
func (c *Case) Camel(upper bool) *Case {
	from := string(*c)
	from = strings.Replace(from, "_", " ", -1)
	from = strings.Replace(from, "-", " ", -1)

	buffer := strings.Builder{}
	for index, word := range strings.Split(from, " ") {
		if 0 == index && !upper {
			buffer.WriteRune(unicode.ToLower(rune(word[0])))
		} else {
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		}
		buffer.WriteString(word[1:])
	}
	*c = Case(buffer.String())

	return c
}

// LowercaseInitial 首字母小写
func (c *Case) LowercaseInitial() *Case {
	from := string(*c)
	*c = Case(string(unicode.ToLower(rune(from[0]))) + from[1:])

	return c
}
