package kernel

const (
	// PositionNone 无
	PositionNone Position = iota + 1
	// PositionHead 首
	PositionHead
	// PositionTail 尾
	PositionTail
	// PositionAll 全部
	PositionAll
)

type Position uint8
