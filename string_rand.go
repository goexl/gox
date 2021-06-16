package gox

import (
	`math/rand`
	`time`
	`unsafe`
)

const (
	// 随机字符串
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
	// 随机数字
	digitBytes = "123456789"
)

var src = rand.NewSource(time.Now().UnixNano())

// RandString 生成随时字符串
func RandString(length int) string {
	return randString(length, letterBytes)
}

// RandDigit 生成随机数字字符串
func RandDigit(length int) string {
	return randString(length, digitBytes)
}

// RandCode 生成随机验证码
func RandCode() string {
	return RandDigit(6)
}

// 生成随时字符串
func randString(length int, letterBytes string) string {
	b := make([]byte, length)

	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if 0 == remain {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
