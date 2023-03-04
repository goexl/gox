package rand

import (
	"unsafe"
	"math/rand"
)

type stringParams struct {
	length  int
	letters string
	numbers string

	idBits int
	idMask  int64
	idMax   int
}

func newStringParams() *stringParams {
	idBits := 6
	var idMask int64 = 1<<idBits - 1
	idMax := 63 / idBits

	return &stringParams{
		length:  8,
		letters: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789",
		numbers: "123456789",

		idBits: idBits,
		idMask:  idMask,
		idMax:   idMax,
	}
}

func (sp *stringParams) rand(source rand.Source) string {
	bytes := make([]byte, sp.length)
	for index, cache, remain := sp.length-1, source.Int63(), sp.idMax; index >= 0; {
		if 0 == remain {
			cache, remain = source.Int63(), sp.idMax
		}
		if idx := int(cache & sp.idMask); idx < len(sp.letters) {
			bytes[index] = sp.letters[idx]
			index--
		}
		cache >>= sp.idBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&bytes))
}
