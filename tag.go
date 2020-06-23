package gox

import "strings"

type (
	// BaseTag Tag基础
	BaseTag struct {
		options []string
	}

	// Tagger 支持Tag的接口
	Tagger interface {
		Contains(option string) (contains bool)
	}
)

// Contains 判断Tag是否包含某个属性
func (tag *BaseTag) Contains(option string) (contains bool) {
	contains = false

	for _, opt := range tag.options {
		if strings.ToLower(opt) == option {
			contains = true
			break
		}
	}

	return
}

// NamedTag 描述Golang的Tag标记
type NamedTag struct {
	BaseTag

	name string
}

// NewTag 创建新Tag
func NewNamedTag(tagString string) *NamedTag {
	tagArray := strings.Split(tagString, ",")
	tag := &NamedTag{
		BaseTag: BaseTag{options: tagArray[1:]},
		name:    tagArray[0],
	}

	return tag
}

// GetName 获得Tag的名称
func (tag *NamedTag) GetName() string {
	return tag.name
}
