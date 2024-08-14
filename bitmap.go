package gox

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
