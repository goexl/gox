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
	words = make([]string, 16)
	builder := new(strings.Builder)
	latest := ' '
	from := []rune(s.params.From)
	callback := s.params.Callback
	for index := 0; index < len(from); index++ {
		latest = from[index]
		next := from[index+1]
		if !callback(latest) {
			builder.WriteRune(latest)
		} else if callback(latest) && !callback(next) {
			words = append(words, builder.String())
			builder.Reset()
		}
	}

	return
}
