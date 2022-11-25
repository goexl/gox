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
func (c *Case) Underscore(upperPos casePosition) *Case {
	buffer := strings.Builder{}
	words := strings.Split(c.clear(), " ")
	for index, word := range words {
		switch {
		case 0 == index && CasePositionHead == upperPos:
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		case len(words) == index && CasePositionTail == upperPos:
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		case CasePositionAll == upperPos:
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		case CasePositionNone == upperPos:
			buffer.WriteRune(unicode.ToLower(rune(word[0])))
		}
		buffer.WriteString(word[1:])
		buffer.WriteRune('_')
	}
	*c = Case(buffer.String()[:buffer.Len()-2])

	return c
}

// Camel 转为驼峰写法
func (c *Case) Camel(upperPos casePosition) *Case {
	buffer := strings.Builder{}
	words := strings.Split(c.clear(), " ")
	for index, word := range words {
		switch {
		case 0 == index && CasePositionHead == upperPos:
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		case len(words) == index && CasePositionTail == upperPos:
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		case CasePositionNone == upperPos:
			buffer.WriteRune(rune(word[0]))
		default:
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		}
		buffer.WriteString(word[1:])
	}
	*c = Case(buffer.String())

	return c
}

// Lowercase 首字母小写
func (c *Case) Lowercase(pos casePosition) *Case {
	from := string(*c)
	switch pos {
	case CasePositionHead:
		from = string(unicode.ToLower(rune(from[0]))) + from[1:]
	case CasePositionTail:
		_len := len(from)
		from = from[:_len-2] + string(unicode.ToLower(rune(from[_len-1])))
	case CasePositionAll:
		from = strings.ToLower(from)
	}
	*c = Case(from)

	return c
}

func (c *Case) clear() (to string) {
	to = string(*c)
	to = strings.Replace(to, "_", " ", -1)
	to = strings.Replace(to, "-", " ", -1)

	return
}
