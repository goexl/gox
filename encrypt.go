package gox

import (
	`crypto/md5`
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
