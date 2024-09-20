package core

import (
	"strings"
	"unicode"

	"github.com/goexl/gox/internal/constant"
	"github.com/goexl/gox/internal/text/internal/internal/token"
	"github.com/goexl/gox/internal/text/internal/kernel"
	"github.com/goexl/gox/internal/text/internal/param"
)

type Switch struct {
	params *param.Switch
}

func NewSwitch(params *param.Switch) *Switch {
	return &Switch{
		params: params,
	}
}

func (s *Switch) Case() (to string) {
	switch s.params.Type {
	case kernel.TypeCamel:
		to = s.camel()
	case kernel.TypeUnderscore:
		to = s.underscore()
	case kernel.TypeLowercase:
		to = s.lowercase()
	case kernel.TypeUppercase:
		to = s.uppercase()
	case kernel.TypeStrike:
		to = s.strike()
	default:
		to = s.underscore()
	}

	return
}

// 下划线
func (s *Switch) underscore() string {
	return s.convert(constant.RuneUnderscore)
}

// 中划线
func (s *Switch) strike() string {
	return s.convert(constant.RuneStrike)
}

// 驼峰
func (s *Switch) camel() string {
	buffer := new(strings.Builder)
	words := token.Naming(s.params.From)
	for index, word := range words {
		switch {
		case 0 == index && kernel.PositionHead == s.params.Position:
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		case 0 == index && kernel.PositionHead != s.params.Position:
			buffer.WriteRune(unicode.ToLower(rune(word[0])))
		default:
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		}
		buffer.WriteString(word[1:])
	}

	return buffer.String()
}

// 字母小写
func (s *Switch) lowercase() (to string) {
	from := s.params.From
	switch s.params.Position {
	case kernel.PositionHead:
		to = string(unicode.ToLower(rune(from[0]))) + from[1:]
	case kernel.PositionTail:
		length := len(from)
		to = from[:length-2] + string(unicode.ToLower(rune(from[length-1])))
	case kernel.PositionAll:
		to = strings.ToLower(from)
	default:
		to = strings.ToLower(from)
	}

	return
}

// 字母大写
func (s *Switch) uppercase() (to string) {
	from := s.params.From
	switch s.params.Position {
	case kernel.PositionHead:
		to = string(unicode.ToUpper(rune(from[0]))) + from[1:]
	case kernel.PositionTail:
		length := len(from)
		to = from[:length-2] + string(unicode.ToUpper(rune(from[length-1])))
	case kernel.PositionAll:
		to = strings.ToUpper(from)
	default:
		to = strings.ToUpper(from)
	}

	return
}

func (s *Switch) convert(replace rune) string {
	buffer := new(strings.Builder)
	words := token.Naming(s.params.From)
	for index, word := range words {
		switch {
		case 0 == index && kernel.PositionHead == s.params.Position:
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		case len(words) == index && kernel.PositionTail == s.params.Position:
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		case kernel.PositionAll == s.params.Position:
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
		case kernel.PositionNone == s.params.Position:
			buffer.WriteRune(unicode.ToLower(rune(word[0])))
		}
		buffer.WriteString(word[1:])
		buffer.WriteRune(replace)
	}

	return buffer.String()[:buffer.Len()-1]
}
