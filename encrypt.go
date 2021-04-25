package gox

import (
	`crypto/hmac`
	`crypto/md5`
	`crypto/rand`
	`crypto/sha256`
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

// Sha256Hmac 计算SHA256加密
func Sha256Hmac(data string, secret string) (sha string, err error) {
	h := hmac.New(sha256.New, []byte(secret))
	if _, err = h.Write([]byte(data)); nil != err {
		return
	}

	sha = hex.EncodeToString(h.Sum(nil))

	return
}
