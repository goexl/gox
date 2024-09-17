package param

import (
	"unicode"

	"github.com/goexl/gox/internal/text/internal/callback"
)

type Split struct {
	From      string
	Callback  callback.Split
	Separator string
}

func NewSplit(from string) *Split {
	return &Split{
		From:     from,
		Callback: unicode.IsSymbol,
	}
}
