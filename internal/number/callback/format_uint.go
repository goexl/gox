package callback

import (
	"github.com/goexl/gox/internal/callback/constraint"
)

type FormatUint[T constraint.Int | constraint.Uint] func(T, int) string
