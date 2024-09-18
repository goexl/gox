package param

import (
	"github.com/goexl/gox/internal/text/internal/callback"
	"github.com/goexl/gox/internal/text/internal/internal/token"
)

type Split struct {
	From     string
	Callback callback.Split
}

func NewSplit(from string) *Split {
	return &Split{
		From:     from,
		Callback: token.Named,
	}
}
