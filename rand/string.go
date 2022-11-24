package rand

import (
	"math/rand"
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
	length int
	bytes  string
}

func newString(seed int64) *_string {
	return &_string{
		source: rand.NewSource(seed),
		length: 8,
		bytes:  letters,
	}
}

func (s *_string) Length(length int) *_string {
	s.length = length

	return s
}

func (s *_string) Code() *_string {
	s.length = 6
	s.bytes = numbers

	return s
}

func (s *_string) Digital() *_string {
	s.bytes = numbers

	return s
}

func (s *_string) Generate() (value string) {
	return s.rand(s.length, s.bytes)
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
