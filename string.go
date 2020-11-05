package gox

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
)

// StringBuilder 内嵌bytes.StringBuilder，支持连写
type StringBuilder struct {
	*bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{Buffer: new(bytes.Buffer)}
}

func (b *StringBuilder) Append(i interface{}) *StringBuilder {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		_, _ = b.Write(val)
	case rune:
		_, _ = b.WriteRune(val)
	}

	return b
}

func (b *StringBuilder) append(s string) *StringBuilder {
	_, _ = b.WriteString(s)

	return b
}

// 驼峰式写法转为下划线写法
func UnderscoreName(name string, upperInitial bool) string {
	buffer := NewStringBuilder()
	for index, char := range name {
		if unicode.IsUpper(char) {
			if 0 != index {
				buffer.Append('_')
			}
			if upperInitial {
				buffer.Append(unicode.ToUpper(char))
			} else {
				buffer.Append(unicode.ToLower(char))
			}
		} else {
			if 0 == index && upperInitial {
				buffer.Append(unicode.ToUpper(char))
			} else {
				buffer.Append(char)
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
func Contract(separator string, strings ...string) (result string) {
	buffer := NewStringBuilder()

	for index, string := range strings {
		buffer.Append(string)
		if len(strings)-1 != index {
			buffer.Append(separator)
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
