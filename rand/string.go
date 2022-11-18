package rand

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	// 随机字符串
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	idxBits = 6
	idxMask = 1<<idxBits - 1
	idxMax  = 63 / idxBits
	// 随机数字
	numbers = "123456789"
)

type _string struct {
	source rand.Source
}

func newString() *_string {
	return &_string{
		source: rand.NewSource(time.Now().UnixNano()),
	}
}

func (s *_string) rand(length int, letters string) string {
	bytes := make([]byte, length)

	for index, cache, remain := length-1, s.source.Int63(), idxMax; index >= 0; {
		if 0 == remain {
			cache, remain = s.source.Int63(), idxMax
		}
		if idx := int(cache & idxMask); idx < len(letters) {
			bytes[index] = letters[idx]
			index--
		}
		cache >>= idxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&bytes))
}
