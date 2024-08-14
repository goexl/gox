package gox

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Bitmap 位图，通过使用位运算来设置位的开关
type Bitmap struct {
	bits      []uint64
	threshold int
	_         CannotCopy
}

// NewBitmap 创建位图
func NewBitmap() *Bitmap {
	return &Bitmap{
		bits:      make([]uint64, 0, 1),
		threshold: 256,
	}
}

func (b *Bitmap) Set(pos uint) *Bitmap {
	blkAt := int(pos >> 6)
	bitAt := int(pos % 64)
	if size := len(b.bits); blkAt >= size {
		b.grow(blkAt)
	}
	b.bits[blkAt] |= 1 << bitAt

	return b
}

func (b *Bitmap) Unset(pos uint) *Bitmap {
	if blkAt := int(pos >> 6); blkAt < len(b.bits) {
		bitAt := int(pos % 64)
		b.bits[blkAt] &^= 1 << bitAt
	}

	return b
}

func (b *Bitmap) Contains(pos uint) (contains bool) {
	blkAt := int(pos >> 6)
	if size := len(b.bits); blkAt >= size {
		contains = false
	} else {
		bitAt := int(pos % 64)
		contains = (b.bits[blkAt] & (1 << bitAt)) > 0
	}

	return
}

func (b *Bitmap) MarshalJSON() ([]byte, error) {
	builder := new(strings.Builder)
	for index := len(b.bits) - 1; index >= 0; index-- {
		b.writeHexDecimal(builder, b.bits[index], true)
	}

	return json.Marshal(builder.String())
}

func (b *Bitmap) UnmarshalJSON(data []byte) (err error) {
	value := ""
	if nil == data {
		b.bits = make([]uint64, 0)
	} else if me := json.Unmarshal(data, &value); nil != me {
		err = me
	} else if fhe := b.FromHex(value); nil != fhe {
		err = fhe
	}

	return
}

func (b *Bitmap) FromHex(value string) (err error) {
	if bytes, dse := hex.DecodeString(value); nil != dse {
		err = dse
	} else if 0 != len(bytes) {
		for left, right := 0, len(bytes)-1; left < right; left, right = left+1, right-1 {
			bytes[left], bytes[right] = bytes[right], bytes[left]
		}

		for len(bytes)%8 != 0 {
			bytes = append(bytes, 0)
		}
		err = b.FromBytes(bytes)
	}

	return
}

func (b *Bitmap) FromBytes(bytes []byte) (err error) {
	if 0 != len(bytes)%8 {
		err = errors.New("数组长度必须是8的倍数")
	} else {
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&b.bits))
		hdr.Len = len(bytes) >> 3
		hdr.Cap = hdr.Len
		hdr.Data = uintptr(unsafe.Pointer(&(bytes)[0]))
	}

	return
}

func (b *Bitmap) writeHexDecimal(builder *strings.Builder, value uint64, pad bool) {
	maxLen := 16
	hexadecimal := strings.ToUpper(strconv.FormatUint(value, 16))
	hexLen := len(hexadecimal)
	if !pad || hexLen == maxLen {
		builder.WriteString(hexadecimal)
	} else {
		for index := hexLen; index < maxLen; index++ {
			builder.WriteString("0")
		}
		builder.WriteString(hexadecimal)
	}
}

func (b *Bitmap) grow(at int) {
	if len(b.bits) > at {
		return
	}

	if cap(b.bits) > at { // 如果存储空间足够，只需要取子数组
		b.bits = b.bits[:at+1]
	} else {
		old := b.bits
		b.bits = make([]uint64, at+1, b.resize(cap(old), at+1))
		copy(b.bits, old)
	}
}

func (b *Bitmap) size(original int, now int) (size int) {
	if now < b.threshold {
		now |= now >> 1
		now |= now >> 2
		now |= now >> 4
		now |= now >> 8
		now |= now >> 16
		now++
		size = now
	} else {
		size = b.resize(original, now)
	}

	return
}

func (b *Bitmap) resize(original int, now int) (resize int) {
	if original < b.threshold {
		original = b.threshold
	}

	for 0 < original && original < (now+1) {
		original += (original + 3*b.threshold) / 4
	}
	resize = original

	return
}
