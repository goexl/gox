package gox

import (
	"sync"
)

// Pointerized 限制结构体只能使用指针形式
type Pointerized [0]sync.Mutex
