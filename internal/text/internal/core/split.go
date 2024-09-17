package core

import (
	"strings"

	"github.com/goexl/gox/internal/text/internal/param"
)

type Split struct {
	params *param.Split
}

func NewSplit(params *param.Split) *Split {
	return &Split{
		params: params,
	}
}

func (s *Split) Apply() (words []string) {
	if "" == s.params.Separator {
		words = s.withRune()
	}

	return
}

func (s *Split) withRune() (words []string) {
	words = make([]string, 0, 16)
	builder := new(strings.Builder)
	latest := ' '
	from := []rune(s.params.From)
	callback := s.params.Callback
	length := len(from)
	for index := 0; index < length; index++ {
		latest = from[index]
		if !callback(latest) {
			builder.WriteRune(latest)
		} else if length-1 > index && callback(latest) && !callback(from[index+1]) {
			words = append(words, builder.String())
			builder.Reset()
		}

		if length-1 == index {
			words = append(words, builder.String())
		}
	}

	return
}
