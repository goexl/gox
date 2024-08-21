package callback

import (
	"github.com/goexl/gox/internal/callback/constraint"
)

type FormatInt[T constraint.Int] func(T, int) string
