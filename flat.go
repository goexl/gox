package gox

import (
	"encoding/json"
	"strconv"
)

type (
	flatType interface {
		map[string]any | string
	}

	flatConverter[T flatType] struct {
		from   T
		style  *flatStyle
		prefix string
	}

	flatStyle struct {
		before string
		middle string
		after  string
	}
)

func Flat[T flatType](from T) *flatConverter[T] {
	return &flatConverter[T]{
		from:   from,
		style:  &flatStyle{middle: "."},
		prefix: "",
	}
}

func (fc *flatConverter[T]) DotStyle() *flatConverter[T] {
	fc.style = &flatStyle{middle: "."}

	return fc
}

func (fc *flatConverter[T]) PathStyle() *flatConverter[T] {
	fc.style = &flatStyle{middle: "/"}

	return fc
}

func (fc *flatConverter[T]) RailsStyle() *flatConverter[T] {
	fc.style = &flatStyle{before: "[", after: "]"}

	return fc
}

func (fc *flatConverter[T]) UnderscoreStyle() *flatConverter[T] {
	fc.style = &flatStyle{middle: "_"}

	return fc
}

func (fc *flatConverter[T]) Prefix(prefix string) *flatConverter[T] {
	fc.prefix = prefix

	return fc
}

func (fc *flatConverter[T]) Convert() (to map[string]any, err error) {
	to = make(map[string]any)
	switch from := any(fc.from).(type) {
	case map[string]any:
		err = fc.flatten(true, to, from, fc.prefix)
	case string:
		err = fc.string([]byte(from), to)
	}

	return
}

func (fc *flatConverter[T]) string(bytes []byte, to map[string]any) (err error) {
	from := new(map[string]any)
	if err = json.Unmarshal(bytes, from); nil == err {
		err = fc.flatten(true, to, from, fc.prefix)
	}

	return
}

func (fc *flatConverter[T]) flatten(top bool, flat map[string]any, nested any, prefix string) (err error) {
	switch _nested := nested.(type) {
	case map[string]any:
		for key, value := range _nested {
			newKey := fc.enKey(top, prefix, key)
			err = fc.assign(flat, newKey, value)
		}
	case []any:
		for index, value := range _nested {
			newKey := fc.enKey(top, prefix, strconv.Itoa(index))
			err = fc.assign(flat, newKey, value)
		}
	}

	return
}

func (fc *flatConverter[T]) assign(flat map[string]any, newKey string, value any) (err error) {
	switch value.(type) {
	case map[string]any, []any:
		err = fc.flatten(false, flat, value, newKey)
	default:
		flat[newKey] = value
	}

	return
}

func (fc *flatConverter[T]) enKey(top bool, prefix string, subKey string) string {
	key := prefix

	if top {
		key += subKey
	} else {
		key += fc.style.before + fc.style.middle + subKey + fc.style.after
	}

	return key
}
