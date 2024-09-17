package builder

import (
	"unicode"

	"github.com/goexl/gox/internal/constant"
	"github.com/goexl/gox/internal/text/internal/callback"
	"github.com/goexl/gox/internal/text/internal/core"
	"github.com/goexl/gox/internal/text/internal/param"
)

type Split struct {
	params *param.Split
}

func NewSplit(from string) *Split {
	return &Split{
		params: param.NewSplit(from),
	}
}

func (s *Split) Letter() (split *Split) {
	s.params.Callback = unicode.IsLetter
	split = s

	return
}

func (s *Split) Connection() (split *Split) {
	s.params.Callback = s.any(
		constant.RuneStrike,
		constant.RuneUnderscore,
		constant.RuneComm,
		constant.RuneDot,
		constant.RuneSpace,
	)
	split = s

	return
}

func (s *Split) Build() *core.Split {
	return core.NewSplit(s.params)
}

func (*Split) any(runes ...rune) callback.Split {
	return func(check rune) (checked bool) {
		for _, rune := range runes {

		}

		return
	}
}
