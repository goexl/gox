package rand

import (
	"time"
)

var _ = New

type generator struct {
	seed int64
}

// New 创建随机字符串
func New() *generator {
	return &generator{
		seed: time.Now().UnixNano(),
	}
}

func (g *generator) Seed(seed int64) *generator {
	g.seed = seed

	return g
}

func (g *generator) String() *_string {
	return newString(g.seed)
}
