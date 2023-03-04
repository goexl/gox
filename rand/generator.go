package rand

import (
	"math/rand"
	"time"
)

var _ = New

type generator struct {}

// New 创建随机字符串
func New() *generator {
	return new(generator)
}

func (g *generator) String() *stringBuilder {
	return newStringBuilder()
}

func (g *generator) Duration() *stringBuilder {
	return newDurationBuilder(g.source)
}
