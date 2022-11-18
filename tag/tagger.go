package tag

// Tagger 支持Tag的接口
type Tagger interface {
	Contains(option string) bool
}
