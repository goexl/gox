package gox

import (
	`sync`
)

// CannotCopy 不能复制
type CannotCopy [0]sync.Mutex
