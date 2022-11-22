package gox

import (
	"encoding/json"
	"strconv"
)

type (
	flattenType interface {
		map[string]any | string
	}

	flatten[T flattenType] struct {
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

func Flatten[T flattenType](from T) *flatten[T] {
	return &flatten[T]{
		from:   from,
		style:  &flattenStyle{middle: "."},
		prefix: "",
	}
}

func (mf *flatten[T]) DotStyle() *flatten[T] {
	mf.style = &flattenStyle{middle: "."}

	return mf
}

func (mf *flatten[T]) PathStyle() *flatten[T] {
	mf.style = &flattenStyle{middle: "/"}

	return mf
}

func (mf *flatten[T]) RailsStyle() *flatten[T] {
	mf.style = &flattenStyle{before: "[", after: "]"}

	return mf
}

func (mf *flatten[T]) UnderscoreStyle() *flatten[T] {
	mf.style = &flattenStyle{middle: "_"}

	return mf
}

func (mf *flatten[T]) Prefix(prefix string) *flatten[T] {
	mf.prefix = prefix

	return mf
}

func (mf *flatten[T]) Convert() (to map[string]any, err error) {
	to = make(map[string]any)
	from := any(mf.from)
	switch _from := from.(type) {
	case map[string]any:
		err = mf.flatten(true, &to, _from)
	case string:
		err = mf.string([]byte(_from), &to)
	}

	return
}

func (mf *flatten[T]) string(bytes []byte, to *map[string]any) (err error) {
	from := new(map[string]any)
	if err = json.Unmarshal(bytes, from); nil == err {
		err = mf.flatten(true, to, from)
	}

	return
}

func (mf *flatten[T]) flatten(top bool, flatMap *map[string]any, nested any) (err error) {
	switch nested.(type) {
	case map[string]any:
		for key, value := range nested.(map[string]any) {
			newKey := mf.enKey(top, key)
			err = mf.assign(flatMap, newKey, value)
		}
	case []any:
		for index, value := range nested.([]interface{}) {
			newKey := mf.enKey(top, strconv.Itoa(index))
			err = mf.assign(flatMap, newKey, value)
		}
	}

	return
}

func (mf *flatten[T]) assign(flatMap *map[string]any, newKey string, value any) (err error) {
	switch value.(type) {
	case map[string]any, []any:
		if err := mf.flatten(false, flatMap, value); err != nil {
			return err
		}
	default:
		(*flatMap)[newKey] = value
	}

	return nil
}

func (mf *flatten[T]) enKey(top bool, subKey string) string {
	key := mf.prefix

	if top {
		key += subKey
	} else {
		key += mf.style.before + mf.style.middle + subKey + mf.style.after
	}

	return key
}
