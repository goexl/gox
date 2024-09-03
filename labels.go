package gox

import (
	"slices"
	"strconv"
	"strings"

	"github.com/goexl/gox/internal/constant"
)

// Labels 标签
type Labels map[string]string

func (l Labels) String() string {
	size := 2
	builder := new(strings.Builder)
	keys := make([]string, 0, len(l))
	for _key, value := range l {
		keys = append(keys, _key)
		size += len(_key) + 2 + len(value) + 3
	}

	builder.Grow(size)
	builder.WriteRune(constant.JsonObjectStart)
	slices.Sort(keys)
	for index, _key := range keys {
		if index > 0 {
			builder.WriteRune(constant.Comm)
		}
		builder.WriteString(_key)
		builder.WriteRune(constant.Equal)
		builder.WriteString(strconv.Quote(l[_key]))
	}
	builder.WriteRune(constant.JsonObjectEnd)

	return builder.String()
}
