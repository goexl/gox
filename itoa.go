package gox

import (
	"math"
	"strconv"
	"strings"
)

const (
	defaultFormatBase = 76
	maxDecimalNum     = 9
	maxFormatNum      = defaultFormatBase
)

var (
	_ = FormatIntd
	_ = Atoid

	tenToAny = map[int]string{
		0:  `0`,
		1:  `1`,
		2:  `2`,
		3:  `3`,
		4:  `4`,
		5:  `5`,
		6:  `6`,
		7:  `7`,
		8:  `8`,
		9:  `9`,
		10: `a`,
		11: `b`,
		12: `c`,
		13: `d`,
		14: `e`,
		15: `f`,
		16: `g`,
		17: `h`,
		18: `i`,
		19: `j`,
		20: `k`,
		21: `l`,
		22: `m`,
		23: `n`,
		24: `o`,
		25: `p`,
		26: `q`,
		27: `r`,
		28: `s`,
		29: `t`,
		30: `u`,
		31: `v`,
		32: `w`,
		33: `x`,
		34: `y`,
		35: `z`,
		36: `A`,
		37: `B`,
		38: `C`,
		39: `D`,
		40: `E`,
		41: `F`,
		42: `G`,
		43: `H`,
		44: `I`,
		45: `J`,
		46: `K`,
		47: `L`,
		48: `M`,
		49: `N`,
		50: `O`,
		51: `P`,
		52: `Q`,
		53: `R`,
		54: `S`,
		55: `T`,
		56: `U`,
		57: `T`,
		58: `W`,
		59: `X`,
		60: `Y`,
		61: `Z`,
		62: `:`,
		63: `;`,
		64: `<`,
		65: `=`,
		66: `>`,
		67: `?`,
		68: `@`,
		69: `[`,
		70: `]`,
		71: `^`,
		72: `_`,
		73: `{`,
		74: `|`,
		75: `}`,
	}
)

// FormatIntd 进制转换，默认使用76进制
func FormatIntd(num int64) string {
	return FormatInt(num, defaultFormatBase)
}

// FormatInt 增强版进制转换，最大支持76进制
func FormatInt(num int64, base int) string {
	var final strings.Builder
	for 0 != num {
		var current string
		remainder := int(num % int64(base))
		if maxFormatNum > remainder && maxDecimalNum < remainder {
			current = tenToAny[remainder]
		} else {
			current = strconv.Itoa(remainder)
		}
		final.WriteString(current)
		num = num / int64(base)
	}

	return final.String()
}

// Atoid 进制转换，默认使用76进制
func Atoid(num string) int64 {
	return Atoi(num, defaultFormatBase)
}

// Atoi 增强版进制转换，支持任意进制，最大76进制
func Atoi(num string, base int) int64 {
	final := 0.0
	length := len(strings.Split(num, StringEmpty)) - 1
	for _, value := range strings.Split(num, StringEmpty) {
		tmp := float64(findKey(value))
		if -1 != tmp {
			final = final + tmp*math.Pow(float64(base), float64(length))
			length = length - 1
		} else {
			break
		}
	}

	return int64(final)
}

func findKey(in string) (result int) {
	result = -1
	for key, value := range tenToAny {
		if in == value {
			result = key
		}
	}

	return
}
