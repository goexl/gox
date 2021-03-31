package gox

import (
	`crypto/md5`
	`crypto/rand`
	`encoding/hex`
)

// Md5 将字符串加密成Md5字符串
func Md5(from string) (to string, err error) {
	alg := md5.New()
	if _, err = alg.Write([]byte(from)); nil != err {
		return
	}

	to = hex.EncodeToString(alg.Sum(nil))

	return
}

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
