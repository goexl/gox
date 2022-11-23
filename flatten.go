package gox

import (
	"encoding/json"
	"strconv"
)

type (
	flattenType interface {
		map[string]any | string
	}

	flattenConverter[T flattenType] struct {
		from   T
		style  *flattenStyle
		prefix string
	}

	flattenStyle struct {
		before string
		middle string
		after  string
	}
)

func Flatten[T flattenType](from T) *flattenConverter[T] {
	return &flattenConverter[T]{
		from:   from,
		style:  &flattenStyle{middle: "."},
		prefix: "",
	}
}

func (fc *flattenConverter[T]) DotStyle() *flattenConverter[T] {
	fc.style = &flattenStyle{middle: "."}

	return fc
}

func (fc *flattenConverter[T]) PathStyle() *flattenConverter[T] {
	fc.style = &flattenStyle{middle: "/"}

	return fc
}

func (fc *flattenConverter[T]) RailsStyle() *flattenConverter[T] {
	fc.style = &flattenStyle{before: "[", after: "]"}

	return fc
}

func (fc *flattenConverter[T]) UnderscoreStyle() *flattenConverter[T] {
	fc.style = &flattenStyle{middle: "_"}

	return fc
}

func (fc *flattenConverter[T]) Prefix(prefix string) *flattenConverter[T] {
	fc.prefix = prefix

	return fc
}

func (fc *flattenConverter[T]) Convert() (to map[string]any, err error) {
	to = make(map[string]any)
	switch from := any(fc.from).(type) {
	case map[string]any:
		err = fc.flatten(true, to, from, fc.prefix)
	case string:
		err = fc.string([]byte(from), to)
	}

	return
}

func (fc *flattenConverter[T]) string(bytes []byte, to map[string]any) (err error) {
	from := new(map[string]any)
	if err = json.Unmarshal(bytes, from); nil == err {
		err = fc.flatten(true, to, from, fc.prefix)
	}

	return
}

func (fc *flattenConverter[T]) flatten(top bool, flat map[string]any, nested any, prefix string) (err error) {
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

func (fc *flattenConverter[T]) assign(flat map[string]any, newKey string, value any) (err error) {
	switch value.(type) {
	case map[string]any, []any:
		err = fc.flatten(false, flat, value, newKey)
	default:
		flat[newKey] = value
	}

	return
}

func (fc *flattenConverter[T]) enKey(top bool, prefix string, subKey string) string {
	key := prefix

	if top {
		key += subKey
	} else {
		key += fc.style.before + fc.style.middle + subKey + fc.style.after
	}

	return key
}
