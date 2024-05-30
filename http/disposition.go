package http

import (
	"github.com/goexl/gox/http/internal/builder"
)

func Disposition(filename string) *builder.Disposition {
	return builder.NewDisposition(filename)
}
