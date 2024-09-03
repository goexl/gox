package gox

import (
	"encoding/json"
	"encoding/xml"

	"github.com/goexl/gox/internal/constant"
)

// Slice 切片，既可以兼容单个值也可以兼容数组
type Slice[T any] []T

// NewSlice 快速创建数组
func NewSlice[T any](items ...T) Slice[T] {
	return items
}

func (s *Slice[T]) Length() int {
	return len(*s)
}

func (s *Slice[T]) Capacity() int {
	return cap(*s)
}

func (s *Slice[T]) Clone() (t Slice[T]) {
	t = make(Slice[T], len(*s))
	copy(t, *s)

	return
}

func (s *Slice[T]) UnmarshalJSON(bytes []byte) (err error) {
	t := new(T)
	ts := make([]T, 0)
	start := bytes[0]
	if constant.JsonArrayStart == start {
		err = json.Unmarshal(bytes, &ts)
	} else if constant.JsonObjectStart == start {
		err = json.Unmarshal(bytes, t)
	}
	if constant.JsonObjectStart == start && nil == err {
		*s = []T{*t}
	} else if constant.JsonArrayStart == start && nil == err {
		*s = ts
	}

	return
}

func (s *Slice[T]) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) (err error) {
	t := new(T)
	if ue := decoder.DecodeElement(t, &start); nil != ue {
		err = decoder.DecodeElement(s, &start)
	} else {
		*s = []T{*t}
	}

	return
}
