package tag

import (
	"strings"
)

var _ = New

type named struct {
	base

	name string
}

// New 创建标签
func New(tag string) (tagger Tagger) {
	options := strings.Split(tag, ",")
	tagger = &named{
		base: base{options: options[1:]},
		name: options[0],
	}

	return
}

// Name 获得标签名
func (n *named) Name() string {
	return n.name
}
