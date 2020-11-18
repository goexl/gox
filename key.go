package gox

import (
	`crypto/rand`
)

// SymmetricKey 生成对称加密算法密钥
func SymmetricKey(bits uint) (key []byte, err error) {
	if 0 != bits%8 {
		err = ErrKeySize

		return
	}

	size := bits / 8
	key = make([]byte, size)
	_, err = rand.Read(key)

	return
}
