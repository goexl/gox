package gox

import (
	"math"
	"strings"

	"github.com/goexl/gox/internal/callback/constraint"
	"github.com/goexl/gox/internal/constant"
	"github.com/goexl/gox/internal/number"
	"github.com/goexl/gox/internal/number/callback"
)

// FormatInt 增强版进制转换，最大支持76进制
func FormatInt[T constraint.Int](num T, base uint8, callback callback.FormatInt[T]) string {
	final := new(strings.Builder)
	for 0 != num {
		var current string
		remainder := int(int64(num) % int64(base))
		if constant.MaxFormat > remainder && constant.MaxDecimal < remainder {
			current = number.TenToAny[remainder]
		} else {
			current = callback(T(remainder), 10)
		}
		final.WriteString(current)
		num = num / T(base)
	}

	return final.String()
}

func FormatUint[T constraint.Uint | constraint.Uint](num T, base uint8, callback callback.FormatUint[T]) string {
	final := new(strings.Builder)
	for 0 != num {
		var current string
		remainder := int(uint64(num) % uint64(base))
		if constant.MaxFormat > remainder && constant.MaxDecimal < remainder {
			current = number.TenToAny[remainder]
		} else {
			current = callback(T(remainder), 10)
		}
		final.WriteString(current)
		num = num / T(base)
	}

	return final.String()
}

// Atoi 增强版进制转换，支持任意进制，最大76进制
func Atoi[T constraint.Int | constraint.Uint](num string, base int) T {
	final := 0.0
	length := len(strings.Split(num, constant.Empty)) - 1
	for _, value := range strings.Split(num, constant.Empty) {
		tmp := float64(number.FindKey(value))
		if -1 != tmp {
			final = final + tmp*math.Pow(float64(base), float64(length))
			length = length - 1
		} else {
			break
		}
	}

	return T(final)
}
